package main

import (
	"context"
	"flag"
	"github.com/Dima191/SmartWay-Task/internal/app"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	configPath string
	isDebug    bool
)

func init() {
	flag.StringVar(&configPath, "config", "./config/config.env", "path to config file")
	flag.BoolVar(&isDebug, "debug", false, "enable debug log level")
}

func logger() *slog.Logger {
	var logLevel slog.Level

	isDebug = true
	switch isDebug {
	case true:
		logLevel = slog.LevelDebug
	default:
		logLevel = slog.LevelWarn
	}

	l := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel,
	}))

	return l
}

func main() {
	flag.Parse()

	l := logger()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	a, err := app.New(ctx, configPath, l)
	if err != nil {
		l.Error(err.Error())
		panic(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err = a.Run(); err != nil {
			l.Error("failed to run or an error occurred  application", slog.String("error", err.Error()))
			stop()
			l.Error(err.Error())
		}
	}()

	<-ctx.Done()
	stop()

	if err = a.Stop(); err != nil {
		l.Error(err.Error())
	}

	wg.Wait()
}
