package main

type Command struct {
  Directory string
  Filename string
  Template string
  Output string
}

type CtxVars map[string]string

type Task struct {
  Title string
  Schedule []Command
  Ctx []string
  Vars CtxVars
  Base string
}

type Config []Task