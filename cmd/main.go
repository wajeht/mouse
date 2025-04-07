package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/wajeht/mouse/internal/mouse"
)

func main() {
	var cfg mouse.Config
	flag.IntVar(&cfg.Size, "size", 100, "Size of square movement")
	flag.IntVar(&cfg.Delay, "delay", 500, "Delay between movements (ms)")
	flag.BoolVar(&cfg.DryRun, "dry-run", false, "Simulate actions only")
	flag.Parse()

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	ctrl := &mouse.RobotController{}
	mover := mouse.NewMover(cfg, ctrl, rng)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sig
		fmt.Println("\nStopping...")
		cancel()
	}()

	if err := mover.Run(ctx); err != nil && err != context.Canceled {
		fmt.Printf("Error: %v\n", err)
	}
}
