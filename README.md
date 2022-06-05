## Genie - A Blazingly fast code generation tool

This is a small ongoing project to create a cli tool similar to [plopjs](https://plopjs.com/) to help speed up workflow with code generation.

Currently the project is a lot more rudimentary than plop, the capabilities include.

- Create config for tasks
- Include array of Handlebars style templates\* per task to generate multiple files 
- Define params and assigned variables in cli
- A small handful of template helper methods (add, toUpper, toLower, toTitle)
\* doesn't include full handlebars feature-set
- A custom built lexer and parser, That runs concurrently. 

The lexer is based on [Rob Pike's lexer](https://talks.golang.org/2011/lex.slide#1)

Genie was mainly built as a project for me to learn some go, It includes my experimentations with goroutines and channels.

## How to use

Pull down the repo, then either run `go build` `go run genie`. Then add the path to the binary to your shell environment.

I have also added an alias to use the --config flag and change location of the config / template files to another directory.

The outputs are currently relative to you cwd

More updates to come 🚀
