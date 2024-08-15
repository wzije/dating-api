# Dating API Demo

This API allow you to find interesting partner for you. This API Provide endpoints to `Register` and `Login` only.

## Prerequisite

- Go: version go1.22.4 or newer
- Mysql: version 5.7.36 or newer
- Git: To Clone the repository
- Golangci-lint (Optional): version 1.59.1 for lint test

## Project Structure

   ``` 
   dating-api/
   ├── http_tests/
   │   └── http_test.http
   ├── src/
   │   ├── configs.go
   │   ├── database.go
   │   ├── entities.go
   │   ├── handlers.go
   │   ├── json.go
   │   ├── request_data.go
   │   ├── router.go
   │   └── securities.go
   ├── go.mod
   ├── go.sum
   └── main.go
   
   ```

### Directory and Files

- `http_test/`: Contains the rest client test / end-to-end test.
    - `http_test.http`: Defines the testing endpoints.
- `src/`: Contains the core business logic and modules.
    - `configs.go`: Defines any configs that used by the application like db connection.
    - `database.go`: Defines any queries that access and mapping the database.
    - `entities.go`: Defines the entities / data that correspondent with the database table.
    - `handlers.go`: Defines HTTP handlers or other interface for interacting with client.
    - `json`: Contains the response data that sent to the client.
    - `request_data`: Contains attributes data for client request mapping and validation.
    - `routers`: Contains the routers service, including middleware.
    - `securities`: Contains security functions like hashing and verifying password, generate token, etc.
- `go.mod`: Specifies the module's dependencies and versioning information.
- `go.sum`: Contains checksums for module dependencies to ensure consistent build.
- `main.go`: The main entry point for the application, where the service is initialized and run.

## Highlight Dependencies

The service uses several Go packages to facilitate various functionalities:

- fiber: Go framework that make easy to build the application.
- go-sql-driver: For interact with database.
- golang-jwt: For implement Json Web Token.
- testify: For testing purpose.

## Installation

Follow these steps to install and run service:

1. Clone Repository

2. Copy or rename .env.example to .env, then adjust the environment variable as needed.
3. Install dependency and build the application

    ```bash 
    go build -o dating-api main.go
    ```

4. Run the application
   ```bash
   ./dating-api
   ```

## Run Test

1. To run tests, use the following command:

    ```bash
    cd dating-api
    go test $(go list ./...) -v
    ```
2. To run Rest HTTP test (End-to-End test)
    - Ensure your application is running
    - Open the project, then Go into the http_tests directory
    - Open `http_test.http` file using code editor like vscode / goland
    - Click Run Code 