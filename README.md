# Go Web Server in Docker

This project demonstrates how to package a simple Go web server into a Docker container.

## Project Structure

- `main.go`: Contains the Go code for the web server.
- `Dockerfile`: Contains the instructions to build the Docker image.

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.20 or later)
- [Docker](https://docs.docker.com/get-docker/)

## Getting Started

### 1. Write a Simple Go Web Server

The `main.go` file contains a basic HTTP server that responds with "Hello from Dockerized Go App!".

### 2. Create a Dockerfile

The `Dockerfile` contains the following instructions:

- Use the `golang:alpine` base image.
- Set the working directory to `/app`.
- Copy `go.mod` and `main.go` files.
- Download dependencies.
- Copy the source code.
- Build the Go application.
- Expose port 8080.
- Specify the command to run the Go application.

### 3. Build and Run the Docker Image

Navigate to the directory containing `main.go` and `Dockerfile`, then run the following commands:

#### Build the Docker image:

```sh
docker build -t go-web-server .
```
#### Access the Web Server
Open your web browser and navigate to http://localhost:8080 to see the message "Hello from Dockerized Go App!".

#### Summary
- Go Web Server: A simple HTTP server that responds with "Hello from Dockerized Go App!".
- Dockerfile: Instructions to package the Go application into a Docker image.
- Docker Commands: Build and run the Docker image to deploy the Go application.

#### Learning Outcomes
Understand the fundamentals of Docker images and containers.
Learn how to package Go applications for deployment.
