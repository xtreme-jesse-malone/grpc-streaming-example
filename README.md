# README

### Server
server listens on port 3000.

```
pushd server
go get .
go build .
./server
popd
```

### Client

The client connects to `<url>:<port>`, opeining connections in `<threadcound>` threads

```
pushd client
go get .
go build .
./client <url>:<port> <threadcount>
```
