# Distributed Applications with Go

Welcome to the repository for Distributed Applications with Go! This repository is dedicated to practicing development and sharing of distributed systems and applications written in the Go programming language. My goal is to do practice in creating robust, scalable, and efficient software that leverages the concurrency model and simplicity of Go.

## Introduction

Distributed systems are a collection of independent components that work together to appear as a single coherent system to the end-user. They are crucial for building scalable services and are prevalent in many modern applications, from web services to databases.

Go, also known as Golang, is an open-source programming language that makes it easy to build simple, reliable, and efficient software. It is particularly well-suited for building distributed systems due to its lightweight goroutines, channels for communication, and its standard library which includes a wealth of packages to support networking and concurrency.

## Repository Structure

This repository is organized as follows:

- `/cmd`: Executable applications.
- `/log`: Logging service.
- `/registry`: Service registration

## Getting Started

To get started with this repository, you should have a working Go environment with version 1.16 or higher. Follow these steps:

1. Clone the repository:
git clone github.com/cyworld8x/go-distributed-applications.git

2. Navigate to the repository directory:
cd app

3. Build the applications:
To build Log Service: go build .\cmd\logservice\
To build Grading Service: go build .\cmd\gradingservice\
To build Registry Service: go build .\cmd\registryservice\
4. Run the applications:
To run Log Service: go run .\cmd\logservice\
To run Grading Service: go run .\cmd\gradingservice\
To run Registry Service: go run .\cmd\registryservice\
6. Run the tests to ensure everything is set up correctly:
go test ./...


## Contributing

We welcome contributions from the community! If you'd like to contribute to the project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Write tests for your changes.
4. Ensure that your code passes all tests.
5. Submit a pull request against the `main` branch.

Please read through our [CONTRIBUTING.md](CONTRIBUTING.md) document for more details on how to submit contributions.

Happy coding!
