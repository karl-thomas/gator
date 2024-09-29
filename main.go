package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/karl-thomas/gator/internal/config"
	"github.com/karl-thomas/gator/internal/database"
	_ "github.com/lib/pq"
)

// its a joke
type Florida struct {
	db   *database.Queries
	Laws *config.Config
}

func (f *Florida) OpenDB() {
	db, err := sql.Open("postgres", f.Laws.DBUrl)
	if err != nil {
		panic(err)
	}
	f.db = database.New(db)
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
	gatorState.OpenDB()

	cmds := commands{
		Commands: make(map[string]func(state *Florida, cmd command) error),
	}
	cmds.register("login", handleLogin)
	cmds.register("register", handleRegister)
	cmds.register("reset", handleReset)
	cmds.register("users", handleUsers)
	cmds.register("agg", handleAgg)
	cmds.register("addfeed", middlewareLoggedIn(handleAddFeed))
	cmds.register("feeds", handleFeeds)
	cmds.register("follow", middlewareLoggedIn(handleFollow))
	cmds.register("following", middlewareLoggedIn(handleFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handleUnfollow))

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
