package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"viper/config"
)

func main() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	env := os.Getenv("RUNOS")
	if env == "" {
		env = "local"
	}
	cfg := config.LoadConfig(env)

	fmt.Printf("config load succ..\n")
	cfg.Print()

	exit := <-sig
	fmt.Printf("recived signal %s\n", exit.String())

}
