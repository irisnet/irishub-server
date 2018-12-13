# IRISHub Server
A Thrift RPC server that allows light clients to interact with IRIS hub
# Structure

- `config`: config of project
- `env`: environment of project
- `errors`: define customer errors
- `models`: database model which defined
- `modules`: modules of project
- `rpc`: define rpc entrance
- `services`: business logic, handle api request
- `utils`: define common constants and functions
- `main.go`: bootstrap project

# RPC Structure

please see [irisnet-rpc](https://github.com/irisnet/irisnet-rpc) 

# Build And Run

- Build: `make build`
- Run: `make run`
- Cross compilation: `make build-linux`

# Run with docker

You can run application with docker.

Example:

```
# docker build -t irishub-server:v1 .
# docker run --name irishub-server -v /mnt/data/iris-log:/irishub-server/log -p 9080:9080 -e "DB_ADDR=127.0.0.1:27017"  -e "ENV=stage" -d irishub-server:v1
```