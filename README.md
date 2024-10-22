# My Go Application

This is a sample Go application that uses Docker and PostgreSQL for data storage.

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or later)
- [Docker](https://docs.docker.com/get-docker/) (version 20.10 or later)
- [Docker Compose](https://docs.docker.com/compose/install/) (version 1.27 or later)
- [Make](https://www.gnu.org/software/make/) (optional, for using the Makefile)

## Environment Variables

Create a `.env` file in the root of the project with the following variables:

``` bash
# Database configuration
DATABASE_URL="host=go_db user=postgres password=postgres dbname=postgres sslmode=disable"
```

## How to Use

### 1. Build the Docker containers

To build the Go app and PostgreSQL containers, run:

```bash
make build
```
### 2. Start the PostgreSQL Database
To start only the PostgreSQL container
```bash
make up-db
```
### 3. Start the Go Application
To start only the Go app container:
```bash
make up-app
```
### 4. Start Both Go App and PostgreSQL
To start both containers:
```bash
make up
```
### 5. Stop the Running Containers
To stop the running containers without removing them: 
```bash
make stop
```
### 6. Stop and Remove the Containers, Networks, and Volumes
To bring down all services and clean up the environment: 
```bash
make down
```

