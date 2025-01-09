package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/bmamha/gator/internal/config"
	"github.com/bmamha/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
		return

	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatal(err)
		return
	}
	dbQueries := database.New(db)

	s := &state{
		dbQueries,
		&cfg,
	}

	cmds := commands{
		make(map[string]func(*state, command) error),
	}

	cmds.register("login", loginHandler)
	cmds.register("register", registerHandler)
	cmds.register("reset", resetHandler)
	cmds.register("users", usersHandler)
	cmds.register("agg", aggCommand)
	if len(os.Args) < 2 {
		log.Fatal("UsageL cli <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	cmd := command{
		cmdName,
		cmdArgs,
	}

	err = cmds.run(s, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
