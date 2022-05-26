package main

import (
	"github.com/manifoldco/promptui"
)

func PromptUser(c Config) Task {
	selected := selectTask(c)

	taskWithVars := setVars(c[selected])
	return taskWithVars
}

func selectTask(config Config) int {
  var options []string

  for _, task := range config {
    options = append(options, task.Title)
  }

  prompt := promptui.Select{
		Label: "Select a Command",
		Items: options,
	}

	idx, _, err := prompt.Run()
	Check(err)

  return idx
}

func setVars(t Task) Task{
	t.Vars = make(TplVars, len(t.Params))

	for _, label := range t.Params{
		t.Vars[label] = prompt(label)
	}
	return t
}

func prompt(label string) string {
	prompt := promptui.Prompt{
		Label: label,
	}

	result, err := prompt.Run()
	Check(err)

	return result
}