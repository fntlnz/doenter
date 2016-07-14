package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/tarm/serial"
)

func read(s *serial.Port) {
	for {
		io.Copy(os.Stdout, s)
	}
}

func write(s *serial.Port) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		s.Write([]byte(text))
	}
}

// This functions handles signals to the main process and routes them to the VM
func sigh(s *serial.Port) {
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

// This function allows the user to leave the tty by pressing Ctrl-c two times in a two seconds time frame
func detach() {
	fmt.Fprintf(os.Stdout, "Send two times Ctrl-c (SIGINT) in order to leave doenter")
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan)

	frec := false
	for range sigchan {
		if frec {
			os.Exit(0)
		}

		timer := time.NewTimer(time.Second * 2)
		frec = true

		go func() {
			<-timer.C
			frec = false
		}()
	}
}

func main() {
	addr, err := homedir.Expand("~/Library/Containers/com.docker.docker/Data/com.docker.driver.amd64-linux/tty")
	if err != nil {
		log.Fatal(err)
	}

	c := &serial.Config{
		Name: addr,
		Baud: 9600,
	}

	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	go read(s)
	go write(s)
	go sigh(s)
	detach()
}
