## Building and Testing a REST API in Go using Gorilla Mux and MySQL

In this practice we will learn how to build and test a simple REST API in Go using Gorilla Mux router
and MySQL database. We will also create the application following the test-driven development (TDD)
methodology.

### Goals

 * Become familiar with the TDD methodology.
 * Become familiar with the Gorilla Mux package.
 * Learn how to use MySQL in Go.
 
### Prerequisites

 * Must have a working Go and MySQL environment.
 
### About the Application

The application is a simple REST API server that will provide endpoints to allow accessing and manipulating 
'users'.

### API Specification

 * Create a new user in response to a valid POST request at `/users`,
 * Update a user in response to a valid PUT request at `/users/{id}`,
 * Delete a user in response to a valid DELETE request at `/users/{id}`,
 * Fetch a user in response to a valid GET request at `/users/{id}`,
 * Fetch a list of users in response to a valid GET request at `/users`.
 
The `{id}` will determine which user the request will work with.

### Setup MySQL Environment

Please refer to [Install MySQL 8 on ubuntu 18.04](https://github.com/joneshsu/mysql8-ubuntu1804)

### Creating the Database

As our application is simple, we will create only one table called `users` with following fields:

 * `id` : is the primary key.
 * `name` : is the name of the user.
 * `age` : is the age of the user.
 
Let's use the following statement to create the database and the table.

```
CREATE DATABASE rest_api_example;
USE rest_api_example;
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    age INT NOT NULL
);
```

### Getting Dependencies

Before we start writing our application, we need to get some dependencies that we will use.
We need to get the two following packages:

 * `mux` : The Gorilla Mux router.
 * `mysql` : The MySQL driver.
 
We can easily use `go get` to get it:

```
go get github.com/gorilla/mux
go get github.com/go-sql-driver/mysql
```

### Testing

We can use following command to run all tests:

```
go test -v
```

```
Testing results:

=== RUN   TestEmptyTable
--- PASS: TestEmptyTable (0.01s)
=== RUN   TestGetNonExistentUser
--- PASS: TestGetNonExistentUser (0.01s)
=== RUN   TestCreateUser
--- PASS: TestCreateUser (0.04s)
=== RUN   TestGetUser
--- PASS: TestGetUser (0.07s)
=== RUN   TestUpdateUser
--- PASS: TestUpdateUser (0.08s)
=== RUN   TestDeleteUser
--- PASS: TestDeleteUser (0.07s)
```