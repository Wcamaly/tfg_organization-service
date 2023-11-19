package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"tfg/organization-service/cmd/config"
	"tfg/organization-service/cmd/server"
	"tfg/organization-service/pkg/config/logs"
)

func main() {
	println("[Config] Loading")
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	if err := logs.InitDefaultLogger(); err != nil {
		panic("error initializing default logger")
	}

	deps, err := config.BuildDependencies(cfg)
	if err != nil {
		logs.Error(context.Background(), "error building dependencies")
		panic(err)
	}

	go func() {
		//logs.Info(context.Background(), fmt.Sprintf("[Server] Listening on %s", cfg.Port))
		if err := server.StartServer(cfg, deps); err != nil {
			logs.Error(context.Background(), "error starting server")
			panic(err)
		}
	}()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	<-exit

}
