set GOOS=linux
set GOARCH=amd64

go build -o checkvalidate checkvalidate.go
go build -o commonok commonok.go
go build -o checkstock checkstock.go