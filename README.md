# learn

A simple Go web service demonstrating the use of the Gin framework and Logrus logger.

## Description

This project provides a basic HTTP server built with [Gin](https://github.com/gin-gonic/gin). It exposes a single endpoint `/` that returns a "Hello World" JSON response. It also utilizes [Logrus](https://github.com/sirupsen/logrus) for logging server startup messages.

## Files

*   `cmd/main.go`: The main application entry point. It sets up the Gin router, defines the `/` route, and starts the HTTP server on port 8080.
*   `go.mod`: Defines the project module (`learn`), Go version (`1.24.1`), and lists project dependencies, including Gin and Logrus.
*   `go.sum`: Contains the checksums for the direct and indirect dependencies (not explicitly shown but implied by `go.mod`).

## Features

*   Simple HTTP server using Gin.
*   One endpoint: `GET /` returns `{"message": "Hello World"}`.
*   Uses Logrus for logging.
*   Listens on port `8080`.

## Prerequisites

*   Go version `1.24.1` or later.

## Installation

1.  Clone the repository (or ensure you have the project files).
2.  Navigate to the project directory.
3.  Download dependencies:
    ```bash
    go mod download