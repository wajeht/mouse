package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/wajeht/mouse/internal/mouse"
)

func main() {
	var cfg mouse.Config
	flag.IntVar(&cfg.Size, "size", 100, "Size of square movement")
	flag.IntVar(&cfg.Delay, "delay", 500, "Delay between movements (ms)")
	flag.BoolVar(&cfg.DryRun, "dry-run", false, "Simulate actions only")
	flag.Parse()

	m := mouse.New(cfg)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sig
		fmt.Println("\nStopping...")
		cancel()
	}()

	if err := m.Run(ctx); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
