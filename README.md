# Taskfile example

This project serves as a showcase for the [Taskfile](https://taskfile.dev/) tool, which is a task runner and build automation tool that helps manage complex projects with ease. The project defines five tasks - build, test, clean, run, and stop - for a small Golang project.

- The **build** task compiles the Golang source code into an executable binary and places it in the `bin` directory.
- The **test** task runs the unit tests for the project using the `go test` command.
- The **clean** task removes any build artifacts from the project, including the `bin` directory and any test output.
- The **run** task runs the compiled binary and sends it to the background.
- The **stop** task stops any running instances of the application.

To run any of these tasks, simply run the `task {taskname}` command in the project directory, where `{taskname}` is one of the five defined tasks. 

## Dependencies

- [Taskfile](https://taskfile.dev/#/installation)
- [Golang](https://golang.org/doc/install)
