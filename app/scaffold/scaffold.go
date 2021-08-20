package scaffold

import (
	"context"
	"os"
	"os/signal"
	"simcart/config"
	internal_types "simcart/pkg/types"
	"syscall"
)

func NewAppScaffold() Scaffold {

	s := &skeleton{params: config.GetAppConfig()}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt,
		syscall.SIGKILL,
		syscall.SIGHUP,
		syscall.SIGTERM,
		syscall.SIGSTOP)
	go func() {
		for {
			select {
			case <-c:
				s.Close()
			}
		}
	}()

	return s
}

func (s *skeleton) Start(ctx context.Context, runServer bool, options ...RunningOption) internal_types.SelectedInterfaces {

	return func(infs ...internal_types.RegisterRPCInterface) error {

		s.context, s.cancel = context.WithCancel(ctx)

		for _, opt := range options {

			if err := opt(ctx, s); err != nil {
				return err
			}

		}
		if runServer {
			return s.server(ctx, infs...)
		}
		return nil
	}

}
func (s *skeleton) SetConfigPath(path string) {
	config.GetAppConfig().SetPath(path)
}
