# exoplanet-service

## Overview
This microservice is designed to manage different exoplanets, providing functionalities such as adding, listing, retrieving, updating, and deleting exoplanets. It also includes a fuel estimation feature for trips to specific exoplanets.


## Prerequisites
- [Go](https://golang.org/dl/) (version 1.16+)
- [Docker](https://www.docker.com/get-started) (for running the service in a container)

## Setup

### Clone the Repository
```
git clone <repository-url>
cd exoplanet-service
```

### Install Dependencies
```
go mod tidy
```

### Running the Service Locally
```
go run main.go
```
The service will be running at http://localhost:8080.



## Using Docker
```
docker build -t exoplanet-service .
docker run -p 8080:8080 exoplanet-service
```
The service will be running at http://localhost:8080.

## Postman

Import the given `postman-collection.json` in postman to test the api's.