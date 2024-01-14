# thttp (TestHTTP)

A simple HTTP server is written in Go for testing purposes.

## Features

- Listens on a specified port.
- Handles requests with random timeouts.
- Logs all incoming requests.

## Usage

```bash
go run main.go -port 8000 -timeout 5000
```

This will start the server on port 8000 with a maximum timeout of 5000 milliseconds for the  endpoint.

### Endpoints

- `timeout`: Handles requests with a random timeout between 0 and the maximum timeout specified at startup.
- `/`: Handles all other requests and logs them.

## Installation

1. Clone the repository.
2. Navigate to the project directory.
3. Run `go build` to compile the application.
4. Run the application with the desired flags.

## Requirements

- Go 1.18 or later.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)

Please replace the [MIT](https://choosealicense.com/licenses/mit/) with the actual link to your license file if you have one.
