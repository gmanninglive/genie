package main

import (
	"fmt"
	"os"
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
	file, err := os.ReadFile(c.Template)
	if err != nil {
		panic(err)
	}

	template := string(file)
	parsed := t.Parser.Parse(template, tplvars)

	c.Output = filepath.Join(c.Directory, c.Filename)

	//fmt.Println(parsed)
	WriteFile(c, []byte(parsed))
	fmt.Printf("created: %s\n", c.Output)
}

func (t Task) parseSchedule(c Command) Command {
	c.Directory = t.Parser.Parse(c.Directory, t.Vars)
	c.Filename = t.Parser.Parse(c.Filename, t.Vars)
	c.Template = filepath.Join(GENIE.BASE, c.Template)

	return c
}
