export CGO_ENABLED=0
export GOOS=linux
go build -ldflags '-s' .
docker build -t xiaorui/frontend .
