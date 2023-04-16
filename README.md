# DDD by Go

This is a sample application using Domain-Driven Design with Go.

## Architecture

### Domain Layer

- **Aggregate**: encapsulates related entities and value objects.
- **Entity**: a unique identifiable object.
- **Value Object**: an object whose equality is based on its value rather than identity.
- **Repository**: interface for persisting and retrieving aggregates.

### Application Layer

- **Command**: request to modify the state of an aggregate.
- **Command Handler**: receives a command and coordinates the use of domain objects to execute the command.
- **Query**: request to retrieve data.
- **Query Handler**: receives a query and retrieves data from the repository.

### Infrastructure Layer

- **Repository Implementation**: implementation of repository using a specific technology (e.g. database, message queue).
- **External Service Client**: client for communicating with external services (e.g. payment gateway, email service).

### Prerequisites

- Go 1.15 or later
- PostgreSQL 9.4 or later

### Installing

1. Clone this repository.
2. Create a database on PostgreSQL.
3. Copy `.env.example` to `.env` and edit it according to your environment.
4. Run the migration with the following command:


## Getting Started

* Clone this repository
```
git clone https://github.com/reivosar/ddd-by-go.git
```
* Launch docker
```
docker-compose up --build
```

## Call API
* Create user
```
curl -X POST -H "Content-Type: application/json" -d "{"userName" : "new user name"}" http://localhost:8080/users
```

* Search user
```
curl -X GET -H "Content-Type: application/json" -d "{"userName" : "created user name"}" http://localhost:8080/users
```

* Change user name
```
curl -X PATCH -H "Content-Type: application/json" -d "{"userId" : "created user id", "userName" : "new user name"}" http://localhost:8080/users
```

* Delete user
```
curl -X DELETE -H "Content-Type: application/json" -d "{"userId" : "created user id"}" http://localhost:8080/users
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.


