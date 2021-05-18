package service

import (
	"bufio"
	"context"
	"fmt"
	"gitlab.com/distributed_lab/logan/v3"
	"log"
	"net"
	"strings"

	"tcp_echo_server/internal/config"
)

type service struct {
	logger           *logan.Entry
	cfg 			 config.Config
}

func NewService(cfg config.Config) *service {
	return &service{
		logger:    cfg.Log(),
		cfg:	   cfg,
	}
}

func (s *service) Run(ctx context.Context) error{
	listener := s.cfg.Listener()

	s.logger.Info("Listening...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			message, _ := bufio.NewReader(conn).ReadString('\n')

			s.logger.Info(fmt.Sprintf("Recived message %v", message))

			res := strings.ToUpper(message)

			c.Write([]byte(res))
		}(conn)
	}

	return nil
}

