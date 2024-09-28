package main

import (
	"fmt"
	"os"

	"github.com/karl-thomas/gator/internal/config"
)

// its a joke
type Florida struct {
	Laws *config.Config
}

type command struct {
	Name string
	Args []string
}

type commands struct {
	Commands map[string]func(state *Florida, cmd command) error
}

func (c *commands) register(name string, f func(state *Florida, cmd command) error) {
	c.Commands[name] = f
}

func (c *commands) run(state *Florida, cmd command) error {
	if f, ok := c.Commands[cmd.Name]; ok {
		return f(state, cmd)
	}
	return fmt.Errorf("command %s not found", cmd.Name)
}

func main() {
	stuff := config.Read()
	// get it
	gatorState := Florida{
		Laws: &stuff,
	}
	cmds := commands{
		Commands: make(map[string]func(state *Florida, cmd command) error),
	}
	cmds.register("login", handleLogin)
	args := os.Args
	if len(args) < 2 {
		fmt.Println("need to provide a command")
		os.Exit(1)
	}
	error := cmds.run(&gatorState, command{Name: args[1], Args: args[2:]})
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
}

func handleLogin(state *Florida, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("need to provide a username")
	}
	error := config.SetUser(cmd.Args[0])
	if error != nil {
		return error
	}
	fmt.Println("logged in as", cmd.Args[0])
	return nil
}
