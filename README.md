[![CircleCI](https://circleci.com/gh/tkeech1/gomodules.svg?style=svg)](https://circleci.com/gh/tkeech1/gomodules)

This repository was created for a Medium post on [Go Modules and CircleCI](https://medium.com/@toddkeech/go-modules-and-circleci-c0d6fac0b000). It shows an example of Go module builds in CircleCI. 

### Running the Example
1. Fork the repository or copy the relevant parts to a different project
2. Configure the project in CircleCI. [Instructions](https://circleci.com/docs/2.0/language-go/) for creating a Go project in CircleCI can be found on the CircleCI web site.
3. Commit a change

CircleCI will execute 3 jobs to build/test on Go 1.10, Go 1.11 and Go 1.11 using modules.

