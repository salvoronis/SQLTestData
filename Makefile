all:
	go build -o sql_tester insert.go

run:
	go run insert.go 10
