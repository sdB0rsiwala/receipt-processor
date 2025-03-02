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
- Docker  
### Optional (For Local Development)
- Go 1.21+
- Postman or cURL for testing  

## Installation & Setup

### Running Locally (Without Docker)
1. Clone the repository:

git clone https://github.com/your-username/receipt-processor.git 
cd receipt-processor

2. Install dependencies:

go mod tidy

3. Run the API:

go run main.go


### Running with Docker
1. Build the Docker image:
docker build -t receipt-processor .

2. Run the container:
docker run -p 8080:8080 receipt-processor


