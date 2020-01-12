# Simple gRPC server

## Dependencies
1. protoc v3
2. golang >= 1.13
3. evans (client for check work)

**main.proto** - schema for rpc

## Build
```bash
$ protoc -I --go_out=plugins=grpc:. main.proto
$ go get -d -v && go build -o .
```

## Run
```bash
$ ./main
```

## Check
```bash
$ evans ./main.proto -p 5000

# evans cli
main.Hello@127.0.0.1:5000> call SayHello
name (TYPE_STRING) => Alice
```