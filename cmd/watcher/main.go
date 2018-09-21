package main

import (
	"log"
	"os"

	"github.com/vterdunov/go-watcher"
)

func main() {
	params := watcher.ParseArgs(os.Args)

	w, err := watcher.MustRegisterWatcher(params)
	if err != nil {
		log.Fatalf("Could not register watcher: %s", err)
	}

	r := watcher.NewRunner()

	// wait for build and run the binary with given params
	go r.Run(params)
	b := watcher.NewBuilder(w, r)

	// build given package
	go b.Build(params)

	// listen for further changes
	go w.Watch()

	r.Wait()
}
