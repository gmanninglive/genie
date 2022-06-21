package main

import (
	"fmt"
	"os"
)

type Env struct {
	Config string
	BASE   string
}

type Flags struct {
	Config string
}

var GENIE Env

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

var helpInfo string = "Flag\t\tDescription\n-c --config\tSet location of config.json:\n\t\t- Include filename,\n\t\t- Custom filnames are accepted aslong as it retains json format,\n\t\t- Lead with ~ to refer to home directory\n"

func ReadFlags(args []string) Flags {
	var config string

	for i := 0; i < len(args); i++ {
		val := args[i]

		switch val {
		case "-c", "--config":
			config = args[i+1]
			i++
		case "-h":
			fmt.Print(helpInfo)

			os.Exit(0)
		default:
			fmt.Printf("Argument %s not recognised please refer to usage below:\n\n %s", val, helpInfo)

			os.Exit(0)
		}
	}

	return Flags{Config: config}
}

func main() {
	var flags Flags

	args := os.Args[1:]
	if args != nil {
		flags = ReadFlags(args)
	}

	config := LoadConfig(flags)

	//selected := PromptUser(config)

	//selected.Run()

	TuiRun(config)
}
