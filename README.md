This repository was created for a Medium post on [Go Modules and CircleCI](https://medium.com/@toddkeech/go-modules-and-circleci-c0d6fac0b000). It shows an example of Go module builds in CircleCI. 

### Running the Example
1. Fork the repository or copy the relevant parts to a different project
2. Configure the project in CircleCI. [Instructions](https://medium.com/r/?url=https%3A%2F%2Fcircleci.com%2Fdocs%2F2.0%2Flanguage-go%2F) for creating a Go project in CircleCI can be found on the CircleCI web site.
3. Commit a change

CircleCI will execute 3 jobs to build/test on Go 1.10, Go 1.11 and Go 1.11 using modules.

