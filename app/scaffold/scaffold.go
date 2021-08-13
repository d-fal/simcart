package scaffold

import (
	"context"
	"os"
	"os/signal"
	"simcart/config"
	"simcart/infrastructure/broker"
	"syscall"
)

func Prepare(ctx context.Context) Scaffold {

	s := &skeleton{params: config.GetAppConfig(), context: ctx}

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

func (s *skeleton) Hibernate(ctx context.Context) error {
	return s.commissioning(ctx)
}

func (s *skeleton) broker() error {
	client := broker.NewClient("nats:4222")

	if err := client.Connect(); err != nil {
		return err
	}

	return nil
}

func (s *skeleton) NormalStart(ctx context.Context) error {

	if err := s.commissioning(ctx); err != nil {
		return err
	}

	if err := s.server(ctx); err != nil {

		return err
	}

	return nil
}

func (s *skeleton) SetConfigPath(path string) {
	config.GetAppConfig().SetPath(path)
}
