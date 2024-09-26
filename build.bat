set GOOS=linux
set GOARCH=amd64

cd cmd
go build -o server main.go