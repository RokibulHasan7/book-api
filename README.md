# REST API Server in Golang

- This is the API Server for a **Book Store**. It allows users to perform Login, Logout into the site and **CRUD** (create, read, update, and delete) operations on the store's inventory of books.

## To Start API Server
```
$ git clone https://github.com/RokibulHasan7/bookApi-server.git
```
```
$ cd book-api
```

## Command to run unit test for API endpoints
```
$ cd api
```
```
$ go test
```

## Data Model

- User Model
``````
type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
``````
- Book Model
``````
type Book struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	ISBN        string   `json:"isbn"`
	PublishDate string   `json:"publishDate"`
	AuthorList  []Author `json:"author"`
}
``````
- Author Model
``````
type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
``````
## Available API Endpoints

|  Method | API Endpoint       | Authentication Type              | Description                                                           |
|---|--------------------|----------------------------------|-----------------------------------------------------------------------|
|POST| /api/v1/login      | Basic                            | Return jwt token in response for successful authentication            |
|POST| /api/v1/logout     | Basic or Bearer token            | Clean token                                                           |
|GET| /api/v1/books      | Basic or Bearer token or No Auth | Return a list of all Books in response                                | 
|GET| /api/v1/books/{id} | Basic or Bearer token or No Auth | Return the data of given Book id in response                          | 
|GET| /api/v1/books/isbn/{isbn} | Basic or Bearer token or No Auth | Return the data of given Book ISBN in response                        |
|POST| /api/v1/books      | Basic or Bearer token            | Add a Book in the database and return the added Book data in response | 
|PUT| /api/v1/books/{id} | Basic Bearer token               | Update the Book and return the updated Book info in response          | 
|DELETE| /api/v1/books/{id} | Basic or Bearer token            | Delete the Book and return the deleted user data in response          | 


## Sample Curl commands without authentication

Run API server without authentication

```shell
$ go run main.go
``` 
On the other CLI run these commands:-

Get all Books information

```shell
$ curl -X GET http://localhost:3333/api/v1/books
``` 

Get Book information with id 1

```shell
$ curl -X GET http://localhost:3333/api/v1/books/1
```

Get Book information with ISBN 14-14-14
```shell
$ curl -X GET http://localhost:3333/api/v1/books/isbn/14-14-14
```

## Sample Curl commands with Basic authentication

Run API server with authentication

```shell
$ go run main.go
``` 
On the other CLI run these commands:-

Get all Books information

```shell
$ curl -X GET --user rakib:1234 http://localhost:3333/api/v1/books
``` 

Get Book information with id 1

```shell
$ curl -X GET --user rakib:1234 http://localhost:3333/api/v1/books/1
```

Get Book information with ISBN 14-14-14
```shell
$ curl -X GET --user rakib:1234 http://localhost:3333/api/v1/books/isbn/14-14-14
```

Create a New Book

```shell
$ curl -X POST  --user rakib:1234 -d '{"id":"6","name":"testfirst","isbn":"testIsbn","publishDate":"12/12/12","author":[{"firstName":"testname","lastName":"testlast","email":"test@gmail.com"}]}' http://localhost:3333/api/v1/books
``` 

Update user data with given id

```shell
$ curl -X PUT  --user rakib:1234 -d '{"name":"testfirst","publishDate":"12/12/12"}' http://localhost:3333/api/v1/books/1
``` 

Delete user with given id

```shell
$ curl -X DELETE --user rakib:1234 http://localhost:3333/api/v1/books/1
``` 

## Sample Curl commands with Bearer token(JWT token) authentication

Run API server with authentication

```shell
$ go run main.go
``` 
On the other CLI run these commands:-

Get jwt token via login with basic authentication

```shell
$ curl -X POST --user rakib:1234  http://localhost:3333/login
``` 

Get all Books information

```shell
$ curl -X GET -H "Authorization: Bearer <jwt_token>"  http://localhost:3333/api/v1/books
``` 

Get Book information with id 1

```shell
$ curl -X GET -H "Authorization: Bearer <jwt_token>" http://localhost:3333/api/v1/books/1
``` 

Create a New Book

```shell
$ curl -X POST -H "Authorization: Bearer <jwt_token>" -d '{"id":"6","name":"testfirst","isbn":"testIsbn","publishDate":"12/12/12","author":[{"firstName":"testname","lastName":"testlast","email":"test@gmail.com"}]}' http://localhost:3333/api/v1/books
``` 

Update Book data with given id

```shell
$ curl -X PUT -H "Authorization: Bearer <jwt_token>" -d '{"name":"test","publishDate":"test"}' http://localhost:3333/api/v1/books/1 
``` 

Delete Book with given id

```shell
$ curl -X DELETE -H "Authorization: Bearer <jwt_token>" http://localhost:3333/api/v1/books/1
``` 