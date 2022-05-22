## Genie - A Blazingly fast code generation tool

This is a small ongoing project to create a cli tool similar to [plopjs](https://plopjs.com/) to help speed up workflow with code generation.

Currently the project is a lot more rudimentary than plop, the capabilities include.

- Create config for tasks
- Include array of Handlebars templates per task to generate multiple files
- Define params and assigned variables in cli
- A small handful of template helper methods (add, toUpper, toLower, toTitle)

This project relies on [raymond](https://github.com/aymerick/raymond/) heavily for parsing the handlebars templates.

## How to use

Pull down the repo, then either run `go build` `go run genie` or use the binary already included. Then add the path to the binary to your shell environment.

I have also added an alias to use the --config flag and change location of the config / template files to another directory.

The outputs are currently relative to you cwd

More updates to come 🚀
