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
	if err != nil {
		fmt.Println("Error reading config")
		return
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	State := state{
		db:  dbQueries,
		cfg: &cfg,
	}

	commands := commands{
		commands: map[string]func(*state, command) error{},
	}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)

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
