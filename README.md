# Jobs Manager Service - WIP

Service to manage jobs

## How to Install

For now, you can run with go run `cmd/manager/main.go`

## How to Use

Define the jobs schema in `config.json`

The service will load and run the jobs in background.

Through the **rest api** you can **get,** **run**, **stop** and **start** the jobs.

The base_url is **http://localhost:8080** and has a postman collection to help in the use.

## Tech Stack

- golang 1.18
- fiber
- gocron