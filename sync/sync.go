package sync

import (
	"fmt"
	"sync"
)

type taskGroup struct {
	wg sync.WaitGroup
}

func (g *taskGroup) Go(name string, f func()) {
	g.start(name, f, false)
}

func (g *taskGroup) GoNoPanic(name string, f func()) {
	g.start(name, f, true)
}

func (g *taskGroup) start(name string, f func(), noPanic bool) {
	g.wg.Add(1)
	go func() {
		defer func() {
			if noPanic {
				if err := recover(); err != nil {
					fmt.Printf("recover_panic, err:%v", err)
				}
			}
			g.wg.Done()
		}()
		f()
	}()
}

func (g *taskGroup) Wait() {
	g.wg.Wait()
}
