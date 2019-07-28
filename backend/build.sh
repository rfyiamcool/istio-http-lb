export CGO_ENABLED=0
export GOOS=linux
go build -ldflags '-s' -o backend
docker build -t xiaorui/backend .
