package main

import (
	"fmt"
	"path/filepath"
)

type Task struct {
	Title    string
	Schedule []Command
	Params   []string
	Vars     TplVars
	Parser   Parser
	Base     string
}

type Command struct {
	Directory string
	Filename  string
	Template  string
	Output    string
}

type TplVars map[string]string

func (t Task) Run() {

	for _, current := range t.Schedule {
		current = t.parseSchedule(current)
		t.runCommand(current, t.Vars)
	}
}

func (t Task) runCommand(c Command, tplvars TplVars) {
	parsed := t.Parser.Parse(c.Template, tplvars)

	c.Output = filepath.Join(c.Directory, c.Filename)
	fmt.Println(parsed)
	WriteFile(c, []byte(parsed))
	fmt.Printf("created: %s\n", c.Output)
}

func (t Task) parseSchedule(c Command) Command {
	//c.Directory = raymond.MustRender(c.Directory, t.Vars)
	//c.Filename = raymond.MustRender(c.Filename, t.Vars)
	c.Template = filepath.Join(GENIE.BASE, c.Template)

	return c
}
