#/bin/sh

rm -rf micro1 micro2

go build -o micro1 micro_cli_1.go
go build -o micro2 micro_cli_2.go

go run root.go micro1