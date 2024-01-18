package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"utils"
)

type Broadcaster struct {
	dest []string
}

const (
	flipFlop = iota
	conjunction
)

type Module struct {
	moduleType int
	on         bool
	name       string
	dest       []string
	input      map[string]bool
	in         chan Pulse
	modules    *map[string]Module
}

func linesToModules(lines []string) (Broadcaster, map[string]Module) {
	b := Broadcaster{}
	modules := make(map[string]Module)
	for _, line := range lines {
		l := strings.Split(line, "->")
		l[0] = strings.Trim(l[0], " ")
		dest := strings.Split(l[1], ",")
		for i := range dest {
			dest[i] = strings.Trim(dest[i], " ")
		}
		switch l[0][0] {
		case '&', '%':
			mod := Module{
				dest:    dest,
				name:    l[0][1:],
				in:      make(chan Pulse),
				modules: &modules,
			}
			switch l[0][0] {
			case '&':
				mod.moduleType = conjunction
				mod.input = make(map[string]bool)
			case '%':
				mod.moduleType = flipFlop
			}
			modules[l[0][1:]] = mod
		default:
			b.dest = dest
		}
	}
	return b, modules
}

const high, low = true, false

type Pulse struct {
	pulse  bool
	source string
}

func (module Module) proc(wg *sync.WaitGroup) {
	pulse := <-module.in
	if module.moduleType == flipFlop {
		module.flipFlop(pulse, wg)
	} else {
		module.conj(pulse, wg)
	}
}

func (module Module) flipFlop(pulse Pulse, wg *sync.WaitGroup) {
	if pulse.pulse == low {
		mutex.Lock()
		module.on = !module.on
		(*module.modules)[module.name] = module
		if module.on {
			pulse.pulse = !pulse.pulse
		}
		pulse.source = module.name
		newWg := new(sync.WaitGroup)
		for _, d := range module.dest {
			if pulse.pulse == high {
				highCount++
			} else {
				lowCount++
			}
			(*newWg).Add(1)
			go (*module.modules)[d].proc(newWg)
			(*module.modules)[d].in <- pulse
		}
		mutex.Unlock()
		(*newWg).Wait()
		wg.Done()
	} else {
		wg.Done()
	}
}

func getConjunctionInputs(modules *map[string]Module) {
	for _, mod := range *modules {
		for _, d := range mod.dest {
			if (*modules)[d].moduleType == conjunction {
				(*modules)[d].input[mod.name] = low
			}
		}
	}
}

func (module Module) conj(pulse Pulse, wg *sync.WaitGroup) {
	mutex.Lock()
	module.input[pulse.source] = pulse.pulse
	(*module.modules)[module.name] = module
	allHigh := true
	for _, remember := range module.input {
		if allHigh && !remember {
			allHigh = false
			break
		}
	}
	if allHigh {
		pulse.pulse = low
	} else {
		pulse.pulse = high
	}
	pulse.source = module.name
	newWg := new(sync.WaitGroup)
	for _, d := range module.dest {
		if pulse.pulse == high {
			highCount++
		} else {
			lowCount++
		}
		_, ok := (*module.modules)[d]
		if !ok {
			continue
		}
		newWg.Add(1)
		go (*module.modules)[d].proc(newWg)
		(*module.modules)[d].in <- pulse
	}
	mutex.Unlock()
	newWg.Wait()
	wg.Done()
}

var lowCount, highCount = 0, 0

var mutex sync.Mutex

func main() {
	utils.CheckArgs(2, os.Args)
	broadcaster, modules := linesToModules(utils.GetLinesFromFile(os.Args[1]))
	getConjunctionInputs(&modules)
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		lowCount++
		mutex.Lock()
		for _, b := range broadcaster.dest {
			wg.Add(1)
			lowCount++
			go modules[b].proc(&wg)
			modules[b].in <- Pulse{pulse: low}
		}
		mutex.Unlock()
		wg.Wait()
	}
	fmt.Println(lowCount * highCount)
}
