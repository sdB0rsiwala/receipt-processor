# Receipt Processor API

## Overview
The Receipt Processor API is a Go-based backend service that processes receipts and awards points based on predefined rules.

## Features
- RESTful API with JSON-based communication
- Stateless in-memory storage
- Dockerized for easy deployment
- Implements receipt processing rules for point calculation

## Prerequisites
### Required
- **Go 1.21.4+** (Required to run the project locally)
### Optional
- Docker (For containerized deployment)
- Postman or cURL for API testing

## **Installation & Setup**

### Running Locally (Recommended)
1. Clone the repository:
```sh
git clone https://github.com/sdB0rsiwala/receipt-processor.git
cd receipt-processor
```

2. Ensure **Go 1.21.4+** is installed on your system.
3. Install dependencies:
```sh
go mod tidy
```

4. Run the API:
```sh
go run main.go 
```

---

### Running with Docker (Optional)
1. Build the Docker image:
```sh
docker build -t receipt-processor . 
```
2. Run the container:
```sh
docker run -p 8080:8080 receipt-processor
```
