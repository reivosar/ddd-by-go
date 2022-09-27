# Overview 

DDD sample created in GO language. This repository was built using clean architecture.

# Required
* Java
* git
* Docker engine

# Getting started
```
1. Clone this repository
  git clone https://github.com/reivosar/ddd-by-go.git
2. Launch docker
  docker-compose up --build
```

# Call API
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
curl -X PATCH -H "Content-Type: application/json" -d "{"userId" : "created user id", "userNamea" : "new user name"}" http://localhost:8080/users
```

* Delete user
```
curl -X DELETE -H "Content-Type: application/json" -d "{"userId" : "created user id"}" http://localhost:8080/users
```