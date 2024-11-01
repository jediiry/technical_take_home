# Project technical_take_home

This project is a technical take-home assignment designed to demonstrate the development of a basic key-value store service in Go. The service exposes a RESTful API that allows clients to store, retrieve, and delete key-value pairs. It also supports graceful shutdown to ensure active requests are completed when stopping the server. The API is structured to be modular, making it easy to extend and integrate with other systems.


## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```
