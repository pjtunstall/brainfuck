package bf

import (
	"time"
)

func WrappedInterpret(i string, t time.Duration) (string, error) {
	var output string
	var err error

	stopInterpretingChan := make(chan(struct{}))
	stopWaitingChan := make(chan(struct{}))
	
	go func() {
		output, err = Interpret(i, stopInterpretingChan, t)
		close(stopWaitingChan)
	}()

	select {
		case <-time.After(t): // Timout expired.
		case <-stopWaitingChan: // Interpret finished naturally with an output.
	}

	close(stopInterpretingChan)
	<- stopWaitingChan // Wait for the output if the timeout signal was sent.

	return output, err
}
