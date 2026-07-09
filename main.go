package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Hamadn/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error reading config")
		return
	}

	State := state{
		cfg: &cfg,
	}

	commands := commands{
		commands: map[string]func(*state, command) error{},
	}
	commands.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	command := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	err = commands.run(&State, command)
	if err != nil {
		log.Fatal(err)
	}
}
