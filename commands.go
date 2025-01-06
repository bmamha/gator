package main

import (
	"fmt"
)

type command struct {
	Name string
	Args []string 
}


type commands struct {
	commandsMap map[string]func(*state, command) error
}



func (c *commands) register(name string, f func(*state, command) error) {
	c.commandsMap[name] = f 
}

func (c *commands) run(s *state, cmd command) error {

	com, ok := c.commandsMap[cmd.Name]
  if !ok {
		return fmt.Errorf("The command %s is not in our commands list", cmd.Name)
	}

  err := com(s,cmd)
  if err != nil {
	 return err 
  }

  return nil 

 } 
