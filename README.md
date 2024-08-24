# Task Management API

## Overview

This is a task management API built using Go. It provides a simple and efficient way to manage tasks, including creating, reading, updating, and deleting tasks.

## Features

* Create tasks with a name, status, project ID, and assigned user ID
* Read tasks by ID or retrieve a list of all tasks
* Validate task payloads to ensure required fields are present
* Use a JWT bearer token for authentication

## API Endpoints

* `POST /tasks`: Create a new task
* `GET /tasks/{id}`: Retrieve a task by ID
* `GET /tasks`: Retrieve a list of all tasks
* `POST /users/register`: Register a new user

## Authentication

This API uses a bearer token for authentication. To use the API, you must provide a valid bearer token in the `Authorization` header of your requests.

## Error Handling

This API returns errors in the following format:

* `400 Bad Request`: Invalid request
* `401 Unauthorized`: Unauthorized
* `403 Forbidden`: Forbidden
* `500 Internal Server Error`: Internal server error

## License

This project is licensed under the [MIT License](LICENSE).

## Contributing

Contributions are welcome! If you'd like to contribute to this project, please fork the repository and submit a pull request.

## Running the API

To run the API, you'll need to have Go installed on your machine. You can then run the API using the following command: `make run`