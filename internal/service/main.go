package service

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"gitlab.com/distributed_lab/logan/v3"
	"log"
	"net"
	"os/exec"
	"strings"

	"tcp_echo_server/internal/config"
)

const ShellToUse = "bash"

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

			var res string

			if string(message[0]) == "/"{
				err, stdout, stderr := Shellout(message[1:])
				if err != nil {
					s.logger.Error("failed to exec cmd")
				}

				s.logger.Info(stdout)

				res = stdout + "\n" + stderr
			} else {
				res = strings.ToUpper(message)
			}

			c.Write([]byte(res))
		}(conn)
	}

	return nil
}

func Shellout(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

