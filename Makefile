SRC = src/io.go src/types.go src/insert.go src/generators.go
OUTPUT = sqlTester

all:
	go build -o $(OUTPUT) $(SRC)

run:
	go run $(SRC) 10

clean:
	rm inserts.sql
