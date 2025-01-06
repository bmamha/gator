package main

import (
	"github.com/bmamha/gator/internal/config"
	"os"
	"log"
)

type state struct {
	cfg *config.Config 
}


func main(){
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error in reading user:%w\n", err)
		return
	}
	s := &state{
		&cfg,
	}
	
	cmds := commands{
		make(map[string]func(*state, command) error), 
	}

	cmds.register("login", handlerLogin)

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
