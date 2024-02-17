# Dependency injection example using samber/do
This repository contains an example of dependency injection in Go using samber/do, as explained in this [article](todo).

## Prerequisites
Ensure PostgreSQL is installed and running. Follow the [official PostgreSQL documentation](https://www.postgresql.org/docs/14/tutorial.html) for installation instructions. If you're using MacOS, you can use the following commands:
```bash
brew install postgres@14
brew services start postgresql
createdb mydb
```

## Running the Program
Start the program by running the following command:
```bash
go run cmd/web/main.go
```

You can then send a request to the running service in a separate terminal:
```bash
curl http://localhost:8080/user?username=joshua
```
