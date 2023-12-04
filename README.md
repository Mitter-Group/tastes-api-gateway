# My API Gateway

This project is an API gateway built with Go, using the Fiber framework and following the principles of clean architecture.

## Project Structure

The project is structured as follows:

- `cmd/server/main.go`: Entry point of the application. It starts the server and sets up the routes and middleware.
- `internal/app/handler`: Contains handlers for handling HTTP requests.
- `internal/app/middleware`: Contains middleware setup for the application.
- `internal/app/router`: Contains route setup for the application.
- `internal/domain`: Contains domain logic.
- `internal/repository`: Contains methods for interacting with the storage.
- `internal/usecase`: Contains methods for performing operations on the domain objects.
- `pkg/utils`: Contains utility functions that can be used throughout the application.

## Getting Started

To run this project, you will need to have Go installed on your machine. You can download it from the official website: https://golang.org/dl/

Once you have Go installed, you can clone this repository and navigate to the project directory:

```bash
git clone https://github.com/yourusername/my-api-gateway.git
cd my-api-gateway
```

Then, you can run the application:

```bash
go run cmd/server/main.go
```

The server will start and listen on port 8080.

## Testing

To run the tests, you can use the `go test` command:

```bash
go test ./...
```

This will run all the tests in the project.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.