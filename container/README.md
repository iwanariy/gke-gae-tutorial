golang app
----

Simple app written in golang.

## Run locally
```go
go run app.go
```

## How to build
```go
docker build --tag iwanariy/goapp .
```

## Run as container
```go
docker run -d -p 8080:8080 iwanariy/goapp
```
