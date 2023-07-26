package signal

import (
	"os"
	"os/signal"
)

type (
	SignalListener func(sig os.Signal)
)

var (
	sigChan   = make(chan os.Signal)
	listeners = make(map[os.Signal][]SignalListener)
)

func On(sig os.Signal, listener SignalListener) {
	if _, ok := listeners[sig]; !ok {
		signal.Notify(sigChan, sig)
		listeners[sig] = []SignalListener{listener}
		return
	}

	listeners[sig] = append(listeners[sig], listener)
}

func Listen() {
	go handle()
}

func Close() {
	signal.Stop(sigChan)
	close(sigChan)
}

func handle() {
	for sig := range sigChan {
		listeners := listeners[sig]
		for _, listener := range listeners {
			listener(sig)
		}
	}
}
