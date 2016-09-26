build:
	GO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o redirect
test:
	go test -cover