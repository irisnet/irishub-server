# IRIS-API-SERVER
IRIS API Server that supports various light clients

# Structure

- `config`: config of project
- `docs`: api documents written by swagger
- `env`: environment of project
- `models`: database model which defined
- `modules`: modules of project
- `rests`: define routes、custom error and vo
- `services`: business logic, handle api request
- `utils`: define common constants and functions
- `main.go`: bootstrap project

# API Documents

1. execute cmd `make run`
2. visit endpoint `{host}:{port}/swagger/index.html` in explorer

# Build And Run

- Build: `make build`
- Run: `make run`
- Cross compilation: `make build-linux` or `make docker-build`
