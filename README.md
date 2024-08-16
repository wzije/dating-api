# Simple Dating API Demo

This API service designed to support the core functionalities of a dating application. 
Right now it's Provide endpoints to `Register` and `Login` only.

## Prerequisite

- Go: version go1.22.4 or newer
- Mysql: version 5.7.36 or newer
- Git: To Clone the repository

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
    - Ensure your application is Running 🔥
    - Open the project, then Go into the http_tests directory
    - Open `http_test.http` file using code editor like vscode / goland
    - Click Run Request / Run All Request

## Integration & Deployment

This automation work on `Gitlab` only.
Before starting, ensure you have set up your environment variables in the GitLab configuration.

Follow these steps to run the integration code:

1. Create a new branch with the prefix feature, e.g., `git checkout -b feature-1`.
2. Push the branch to the GitLab repository using `git push gitlab feature-1`.
3. After pushing the code, the GitLab pipeline will run automatically, and you can monitor it on the GitLab pipeline
   page under the build menu.

___