- build with docker

```
docker build -t goapp .
```

- run it with docker

```
docker run -p 8080:8080 goapp
```

- build without docker 

```
go build -o ./cmd/goapp
```

- run without docker

```
go run ./cmd/goapp
```

- for tests
 
```
go test
```
