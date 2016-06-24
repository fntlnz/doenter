package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/goburrow/serial"
	"github.com/mitchellh/go-homedir"
)

func read(s serial.Port) {
	for {
		buf := make([]byte, 128)
		n, _ := s.Read(buf)
		fmt.Fprintf(os.Stdout, "%s", buf[:n])
	}
}

func write(s serial.Port) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		s.Write([]byte(text))
	}
}

func sigh(s serial.Port) {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan)

	ascii := map[os.Signal]string{
		syscall.SIGINT:  "\x03",
		syscall.SIGTSTP: "\x1A",
	}

	for sig := range sigchan {
		if val, ok := ascii[sig]; ok {
			s.Write([]byte(val))
		}
	}
}

func main() {
	addr, err := homedir.Expand("~/Library/Containers/com.docker.docker/Data/com.docker.driver.amd64-linux/tty")
	if err != nil {
		log.Fatal(err)
	}

	c := &serial.Config{
		Address:  addr,
		BaudRate: 9600,
	}
	s, err := serial.Open(c)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	go read(s)
	go write(s)
	sigh(s)
}
