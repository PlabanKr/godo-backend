# GODO

GoDo, a todo management rest api. It saves your todos in `mysql` database.

## Running The Server

This rest api is written is go lang. Follow these steps to run the server in your machine.

> Note: You need GO 1.19.5 or up and mysql installed in your machine.

- Clone this repo in your machine with `git clone`.
- Open a terminal/command prompt and then navigate to the root of this project.
- Run `go get` to install all the dependencies.
- Open `database.go` file in a text editor and then change the values for `dbConfig`.
- Open `mysql` and create database with same name as `dbConfig`.
- Then create a table. With the following sql query
  ```sql
  create table `todo` (`id` int not null, `title` varchar(255) not null, `body` text not null, `isDone` boolean, primary key(`id`));
  ```
- Now finally run `go build` in the root directory of the project.
- And the last step is to the run the executable binary. :fire:


## Routes

#### POST localhost:8080/todo

Send a post request, with body containing the todo information, on this route to create a todo.
Sample of the body.
```json
{
    "id": 2,
    "title": "First Todo",
    "body": "This is the very first todo",
    "isDone": 1
}
```

#### GET localhost:8080/todos

Send a get request on this route to get all todos.

#### GET localhost:8080/todo/:id

Send a get request on this route with todo id to get that specific todo.

#### PUT localhost:8080/todo/:id

Send a put request on this route with todo id to update that specific todo. In the body of the request you can put any field that you want to update except the `id`.

#### DELETE localhost:8080/todo/:id

Send a delete request on this route with todo id to delete that specific todo.



> INFO: You can't run `go run` in this project. This project contains multiple files under same package. `go run` will only compile a single file and you will get a undefined error.