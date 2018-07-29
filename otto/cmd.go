package main

// Basic command
type Cmd struct {
	Name string
	Args []string
	Cmds []Cmd
	Run func(cmd string, args []string)
}


