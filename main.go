package main

import (
	"database/sql"
	"github.com/Geraetefreund/gator/internal/config"
	"github.com/Geraetefreund/gator/internal/database"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	config, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", config.DBurl)
	if err != nil {
		log.Fatalf("error opening database: %s", err)
	}

	dbQueries := database.New(db)

	programState := &state{
		cfg: &config,
		db:  dbQueries,
	}

	cmds := commands{
		registeredCommand: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)

	if len(os.Args) < 2 {
		log.Fatalf("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
