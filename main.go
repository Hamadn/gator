package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Hamadn/gator/internal/config"
	"github.com/Hamadn/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	db, err := sql.Open("postgres", cfg.DBURL)
	dbQueries := database.New(db)
	if err != nil {
		fmt.Println("Error reading config")
		return
	}

	State := state{
		db:  dbQueries,
		cfg: &cfg,
	}

	commands := commands{
		commands: map[string]func(*state, command) error{},
	}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)

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
