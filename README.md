build with docker build -t goapp .
run it with docker run -p 8080:8080 goapp

build without docker go build -o ./cmd/goapp
run without docker ./cmd/goapp

for tests go test

