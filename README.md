# Go Gin-Gonic CRUD Application

This is a Go Gin web application that performs CRUD operations on a User model. The application is Dockerized for easy deployment and management.

---

## Running the Application with Docker

The application is Dockerized, which means you can run it easily using Docker and Docker Compose.
Prerequisites

    Docker
    Docker Compose

Steps to Run 
---

    # Clone the Repository


    git clone https://github.com/Harshraj13/gin-gonic.git

    cd gin-gonic

Create a .env File
---
```sh 
Create a .env file in the root directory with the following content:

    env

    MONGO_URI: The connection string for your MongoDB instance.
    DB_NAME: The name of your MongoDB database.
```

Build and Run the Application
---
```sh
    docker build -t go.0.1 .

    docker-compose up -d

    This command will build the Docker image and start the container.

    Access the Application

    Once the container is running, you can access the application at 
    
    http://localhost:4455.
```


Project Structure
---

            gin-gonic
            ├── configs
            │   ├── controllers
            │   │   └── user_controller.go
            │   ├── env.go
            │   ├── models
            │   │   └── user_model.go
            │   ├── responses
            │   │   └── user_response.go
            │   └── setup.go
            ├── docker-compose.yml
            ├── Dockerfile
            ├── go.mod
            ├── go.sum
            ├── main.go
            ├── README.md
            └── routes
                └── user_route.go


    Dockerfile: The Dockerfile for building the application image.
    docker-compose.yml: The Docker Compose file for running the application.
    .env: The environment variables file.
    go.mod: The Go module file.
    go.sum: The Go dependencies file.
    main.go: The main application file.
    README.md: This README file.

API Endpoints
---

Here are the available API endpoints for CRUD operations on the User model:

    POST /user - Create a new user
    GET /user/:userId - Retrieve a user by ID
    PUT /user/:userId - Update a user by ID
    DELETE /user/:userId - Delete a user by ID
    GET /users - Retrieve a list of all users

### Example Requests
---
```sh
# Retrieve a specific user
curl --location 'http://localhost:4455/user/665c067a4124173c2f0d5bc7/'

# Create a user with specified details
curl --location 'http://localhost:4455/user/' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Gopal",
    "location": "Russia",
    "title": "narcotic specialist"
}'

# Update a user by ID
curl --location --request PUT 'http://localhost:4455/user/665b66c6978897f2c6057d12' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Hermabh",
    "location": "Canada"
}'

# Delete a user by ID
curl --location --request DELETE 'http://localhost:4455/user/665c103e4124173c2f0d5bc9'

# Retrieve a list of all users
curl --location 'http://localhost:4455/users/'
```


# Contributing

Feel free to fork this repository and submit pull requests. If you encounter any issues, please open an issue on GitHub.