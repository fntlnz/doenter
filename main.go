package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/tarm/serial"
)

func main() {
	// read
	go func() {
		c := &serial.Config{Name: "/dev/ttys001", Baud: 9600}
		s, err := serial.OpenPort(c)
		if err != nil {
			log.Fatal(err)
		}

		for {
			buf := make([]byte, 128)
			n, err := s.Read(buf)
			if err != nil {
				//log.Fatal(err)
			}
			fmt.Fprintf(os.Stdout, "%s", buf[:n])
		}
	}()

	//Write
	go func() {
		c := &serial.Config{Name: "/dev/ttys001", Baud: 9600}
		s, err := serial.OpenPort(c)
		if err != nil {
			log.Fatal(err)
		}

		for {
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			_, err := s.Write([]byte(text))
			if err != nil {
				//log.Fatal(err)
			}
		}
	}()

	// Write signals
	go func() {
		c := &serial.Config{Name: "/dev/ttys001", Baud: 9600}
		s, err := serial.OpenPort(c)
		if err != nil {
			log.Fatal(err)
		}

		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, os.Interrupt)
		for sig := range sigchan {
			switch sig {
			case os.Interrupt:
				// TODO(fntlnz): I don't know how to send signals to the tty atm, blame on me
				break

			}
		}
	}()

	for {

	}

}
