# Book Finder

Book Finder is a Go-based web application for managing and searching books. It provides a RESTful API for adding, updating, deleting, retrieving, and searching books.

## Table of Contents

- [Installation](#installation)
- [Running the Application](#running-the-application)
- [Running Tests](#running-tests)

## Installation

To run this project locally, follow these steps:

1. Install dependencies:


    go mod tidy


## Running the Application

To run the application, use the following command:


    go run .


## Running Tests

To test the application, use the following command:


    go test

## Endpoints

The following endpoints are available for managing books:

GET /books: Retrieve a list of all books.

POST /books: Add a new book.

GET /books/:bookId: Retrieve a specific book by its ID.

PUT /books/:bookId: Update a specific book by its ID.

DELETE /books/:bookId: Delete a specific book by its ID.

GET /books/search: Search for books based on query parameters.


## Docker and Minikube

Docker: Run the application in a Docker container with the following settings:

Access the application at http://localhost:8080

Minikube: Deploy the application in a Minikube cluster with the following settings:

Access the application at http://<minikube ip>:30000

Make sure to replace <minikube ip> with the actual IP address of your Minikube instance.