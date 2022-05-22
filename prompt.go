package main

import (
	"github.com/manifoldco/promptui"
)

func SelectTask(config Config) (int, error){
  options := make([]string, len(config))

  for i := 0; i < len(config); i++ {
    options[i] = config[i].Title
  }

  prompt := promptui.Select{
		Label: "Select a Command",
		Items: options,
	}

	idx, _, err := prompt.Run()

  return idx, err
}

func SetVars(t Task) Task{
	t.Vars = make(CtxVars, len(t.Ctx))
	for i := 0; i < len(t.Ctx); i++ {
		label := t.Ctx[i]
		t.Vars[label] = prompt(label)
	}
	return t
}

func prompt(label string) string {
	prompt := promptui.Prompt{
		Label:   label,
	}

	result, err := prompt.Run()

	Check(err)

	return result
}