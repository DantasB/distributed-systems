package signals

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

// BusyWait receives the signal channel and the exit channel and prints messages when receive some signals.
// It uses a goroutine, a go way to use concurrent programming.
// To sync the concurrency it's necessary to use a channel (in this case a Signal Channel).
// The while true loop have a select-case that checks if the signal was received and waits in this condition.
// But because it has a default sentence it will enter in this section every time the loop runs.
// The signals treated are SIGHUP (1), SIGTERM(15), SIGQUIT(3).
// The exitChan receives a 0 message exiting of the goroutine.
// It has no return.
func busyWait(signalChan chan os.Signal, exitChan chan int) {
	go func() {
		for {
			select {
			case signal := <-signalChan:

				if signal == syscall.SIGHUP {
					fmt.Println("[RECEIVER] SIGHUP received.")
				}
				if signal == syscall.SIGTERM {
					fmt.Println("[RECEIVER] SIGTERM received.")
				}
				if signal == syscall.SIGQUIT {
					fmt.Println("[RECEIVER] SIGQUIT received.")
					exitChan <- 0
				}
			default:
			}
		}
	}()
}

// BlockingWait receives the signal channel and the exit channel and prints messages when receive some signals.
// It uses a goroutine, a go way to use concurrent programming.
// To sync the concurrency it's necessary to use a channel (in this case a Signal Channel).
// The variable signal blocks the process until the signalChann variable receives a signal.
// The while true loop have a switch-case that checks every signal received.
// The signals treated are SIGHUP (1), SIGTERM(15), SIGQUIT(3).
// The exitChan receives a 0 message exiting of the goroutine.
// It has no return.
func blockingWait(signalChan chan os.Signal, exitChan chan int) {
	go func() {
		for {
			signal := <-signalChan
			switch signal {
			case syscall.SIGHUP:
				fmt.Println("[RECEIVER] SIGHUP received.")

			case syscall.SIGTERM:
				fmt.Println("[RECEIVER] SIGTERM received.")

			case syscall.SIGQUIT:
				fmt.Println("[RECEIVER] SIGQUIT received.")
				exitChan <- 0
			}
		}
	}()
}

// InstantiateChannels receives nothing as parameters and returns a tuple containing 2 channels.
// It instantiate the signalChannel that is a channel that receives the os.Signal values.
// It instantiate the exitChannel that is a channel that receives the int values.
// It receives the notification of every signal sent, avoiding any signal utility unless the defined by the previous functions.
// It has no return.
func instantiateChannels() (chan os.Signal, chan int) {
	signalChannel := make(chan os.Signal, 1)

	signal.Notify(signalChannel)

	exitChannel := make(chan int)

	return signalChannel, exitChannel
}

func SignalReceiver() {
	if len(os.Args) != 2 {
		fmt.Println("[RECEIVER] Missing parameter. Please, choose 0 for blocking wait or 1 for busy wait.")
		return
	}

	fmt.Printf("[RECEIVER] Process pid : %d\n", os.Getpid())

	signalChan, exitChan := instantiateChannels()

	mode, _ := strconv.Atoi(os.Args[1])
	switch mode {
	case 0:
		blockingWait(signalChan, exitChan)
	case 1:
		busyWait(signalChan, exitChan)
	default:
		fmt.Println("[RECEIVER] Unknown parameter. Please, choose 0 for blocking wait or 1 for busy wait.")
		return
	}

	code := <-exitChan

	fmt.Println("[RECEIVER] Process finished.")

	os.Exit(code)
}
