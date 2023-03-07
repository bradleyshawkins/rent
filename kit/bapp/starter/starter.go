package starter

import (
	"context"
	"os"
	"os/signal"

	"go.uber.org/multierr"
)

type Starter interface {
	Start(ctx context.Context) (Stopper, error)
}

type Stopper interface {
	Stop(ctx context.Context) error
}

func Start(ctx context.Context, starters ...Starter) error {
	stoppers := make([]Stopper, len(starters))

	for i, starter := range starters {
		stop, err := starter.Start(ctx)
		if err != nil {
			return err
		}

		stoppers[i] = stop
	}

	var c chan os.Signal
	signal.Notify(c, os.Interrupt)

	<-c

	var e error
	for _, stopper := range stoppers {
		err := stopper.Stop(ctx)
		if err != nil {
			err = multierr.Append(e, err)
		}
	}

	return e
}
