package closer

import (
	"log"
	"os"
	"os/signal"
	"sync"
)

var globalCloser = New()

type Closer struct {
	mu    sync.Mutex
	once  sync.Once
	done  chan struct{}
	funcs []func() error
}

func Add(funcs ...func() error) {
	globalCloser.Add(funcs...)
}

func Wait() {
	globalCloser.Wait()
}

func CloseAll() {
	globalCloser.CloseAll()
}

func New(signals ...os.Signal) *Closer {
	closer := &Closer{done: make(chan struct{})}
	if len(signals) > 0 {
		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, signals...)
			<-ch
			signal.Stop(ch)
			closer.CloseAll()
		}()
	}

	return closer
}

func (c *Closer) Add(funcs ...func() error) {
	c.mu.Lock()
	c.funcs = append(c.funcs, funcs...)
	c.mu.Unlock()
}

func (c *Closer) Wait() {
	<-c.done
}

func (c *Closer) CloseAll() {
	c.once.Do(
		func() {
			defer close(c.done)

			c.mu.Lock()
			funcs := c.funcs
			c.funcs = nil
			c.mu.Unlock()

			errs := make(chan error, len(funcs))

			for _, f := range funcs {
				go func(f func() error) {
					errs <- f()
				}(f)
			}

			for err := range errs {
				if err != nil {
					log.Printf("erros: %v\n", err.Error())
				}
			}
		},
	)
}
