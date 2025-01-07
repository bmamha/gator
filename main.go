package main
import _ "github.com/lib/pq"
import (
	"github.com/bmamha/gator/internal/config"
	"database/sql"
	"github.com/bmamha/gator/internal/database"
	"os"
	"log"
)

type state struct {
	db *database.Queries 
	cfg *config.Config
}


func main(){
  cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error in reading user:%w\n", err)
		return
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	dbQueries := database.New(db) 
 
	
	s := &state{
		dbQueries,
		&cfg,
	}
	
	cmds := commands{
		make(map[string]func(*state, command) error), 
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", registerHandler)
  cmds.register("reset", resetHandler)
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

