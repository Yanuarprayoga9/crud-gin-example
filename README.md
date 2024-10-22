# My Go Application

This is a sample Go application that uses Docker and PostgreSQL for data storage.

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or later)
- [Docker](https://docs.docker.com/get-docker/) (version 20.10 or later)
- [Docker Compose](https://docs.docker.com/compose/install/) (version 1.27 or later)
- [Make](https://www.gnu.org/software/make/) (optional, for using the Makefile)

## Environment Variables

Create a `.env` file in the root of the project with the following variables:

```env
# Database configuration
DB_USER=myuser           # Change to your desired username
DB_PASSWORD=mypassword   # Change to your desired password
DB_NAME=mydatabase       # Change to your desired database name

# Optional: Port configuration
DB_PORT=5324             # Change to your desired host port
