<h1 align="center">FizzBuzz API</h1>

<div align="center">

![STATUS](https://img.shields.io/badge/status-active-brightgreen?style=for-the-badge)
![LICENSE](https://img.shields.io/badge/license-MIT-blue?style=for-the-badge)

</div>

---

## 🔎 About <a name = "about"></a>

The objective of this project is to showcase the concepts of SWE and DevOps that I have learned so far. This project features CI/CD workflows and automations, unittesting, and containerization. Git & Github were used as the VCS, and a Docker image is published for container deployment. I dub this as "FizzBuzz-as-a-Service".

This program aims to create a Go-based microservice that solves the classic FizzBuzz programming problem. The FizzBuzz problem is a common coding challenge often used in interviews to evaluate a candidate's basic programming skills. The task is to write a program that prints numbers from 1 to n; however, for multiples of 3, it should print "Fizz" instead of the number, for multiples of 5, it should print "Buzz," and for numbers that are multiples of both 3 and 5, it should print "FizzBuzz". The FizzBuzz Microservice API provides a simple HTTP-based API to solve the FizzBuzz problem. It allows users to make requests to the microservice and receive the FizzBuzz sequence as a response. This microservice can be easily integrated into other applications or used for testing and learning purposes.

## 🏁 Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Before running the FizzBuzz Microservice API, make sure you have the following prerequisites installed:

- Go 1.21 or higher

### Installing

To get started with this project, clone the repository to your local machine and install the required dependencies.

```bash
git clone https://github.com/jgfranco17/fizzbuzz-api.git
cd fizzbuzz-api
go mod tidy
```

## 🚀 Usage <a name = "usage"></a>

### Dev mode

To run the API server in dev mode, simply execute either of the following commands:

```bash
# Default Go execution
go run service/cmd/main.go --port=8000 --dev=true

# Or, after editing the Makefile to set a port number
just run-debug 8000
```

### Build with Docker

To run the microservice in a container, the package comes with both a Dockerfile and a Compose YAML configuration. Run either of the following to get the API launched in a container; by default, the API will be set to listen on port 5050 for the Compose.

```bash
# Plain Docker build
docker build -t fizzbuzz-api .
docker run -p 8000:8000 fizzbuzz-api

# Docker-Compose build
docker compose up
```

### Calling the API

Once the microservice is launched, the API is now reachable. To get a FizzBuzz sequence, simply send a `GET` request to the server, with the number as a parameter.

```bash
curl http://localhost:<PORT>/fizzbuzz/<number>
```

## 🔧 Testing <a name = "testing"></a>

### Running unittest suite

To run the full test suite, run the Justfile command as follows:

```bash
just test
```

## ✒️ Authors <a name = "authors"></a>

- [Chino Franco](https://github.com/jgfranco17)
