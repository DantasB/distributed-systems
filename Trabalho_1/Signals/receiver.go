package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func busy_wait(signal_chan chan os.Signal, exit_chan chan int) {
	go func() {
		for {
			select {
			case signal := <-signal_chan:

				if signal == syscall.SIGHUP {
					fmt.Println("[RECEIVER] SIGHUP received.")
				}
				if signal == syscall.SIGTERM {
					fmt.Println("[RECEIVER] SIGTERM received.")
				}
				if signal == syscall.SIGQUIT {
					fmt.Println("[RECEIVER] SIGQUIT received.")
					exit_chan <- 0
				}
			default:
			}
		}
	}()
}

func blocking_wait(signal_chan chan os.Signal, exit_chan chan int) {
	go func() {
		for {
			signal := <-signal_chan
			switch signal {
			case syscall.SIGHUP:
				fmt.Println("[RECEIVER] SIGHUP received.")

			case syscall.SIGTERM:
				fmt.Println("[RECEIVER] SIGTERM received.")

			case syscall.SIGQUIT:
				fmt.Println("[RECEIVER] SIGQUIT received.")
				exit_chan <- 0
			}
		}
	}()
}

func main() {
	fmt.Printf("pid: %d\n", os.Getpid())
	signal_chan := make(chan os.Signal, 1)

	signal.Notify(signal_chan)

	exit_chan := make(chan int)

	busy_wait(signal_chan, exit_chan)

	code := <-exit_chan

	fmt.Println("[RECEIVER] Process finished.")

	os.Exit(code)
}
