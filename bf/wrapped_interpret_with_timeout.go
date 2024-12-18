package bf

import (
	"sync"
	"time"
)

func WrappedInterpretWithTimeout(i string, t time.Duration) (string, error) {
	var output string
	var err error

	stopInterpretingChan := make(chan(struct{}))
	stopWaitingChan := make(chan(struct{}))
	var once sync.Once
	
	go func() {
		output, err = Interpret(i, stopInterpretingChan, t)
		close(stopWaitingChan)
	}()

	select {
		case <-time.After(t): {
			close(stopInterpretingChan)
		}
		case <-stopWaitingChan:
	}

	once.Do(func () { close(stopInterpretingChan) })
	once.Do(func() { close(stopWaitingChan) })

	return output, err
}