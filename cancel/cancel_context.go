package main

import (
	"context"
	"flag"
	"fmt"
	"os"
)

var (
	verbose bool
)

func init() {
	flag.BoolVar(&verbose, "verbose", false, "turn up the level of noise")
}

func println(str ...string) {
	if verbose {
		fmt.Println(str)
	}
}

func main() {

	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println("Need a command to run")
		os.Exit(2)
	}
	cmd := flag.Arg(1)
	switch cmd {
	case "numbers":
		GeneratorPattern()
	}
}

// GeneratorPattern generates a sequence of integers in a Goroutine
// and sends them out the returned channel.  It is up to the
// consumer to send a cancel once they are done consuming the
// integers.
func GeneratorPattern() {

	// Gen points to the function that creates a channel, strats a
	// Goroutine then starts sending numbers to the returned channel.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1

		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning to NOT leak the Go Routine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	// Client sets up the "cancel" context to have the "server stop sending"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
