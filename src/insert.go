package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"text/template"
	"time"
)

var (
	logfile, _ = os.Open("./log.txt")
	logger = log.New(logfile, "db_test_generator", log.LstdFlags|log.Lshortfile)
	alphabet []string = []string{}
	tables []Request = []Request{}
	dependencies map[string]int = make(map[string]int)
	n int
)

func main() {
	var template_str string = "insert into {{.Table_name}}({{range .Fields}}{{.Field}}{{end}}) values {{range .Data}}\n\t{{.}}{{end}}\n"
	template, err := template.New("sql_test_inserts").Parse(template_str)
	if err != nil {
		logger.Println(err)
	}

	n, err = strconv.Atoi(os.Args[1])
	if err != nil {
		logger.Println(err)
	}

	rand.Seed(time.Now().UnixNano())

	Parse_alphabet("./alphabet.json")
	Parse_scheme("./test_table.json")

	sql_script_file, err := os.Create("./insertss.sql")
	if err != nil {
		logger.Println(err)
	}

	for i := 0; i < len(tables); i++{
		Update_variables(tables[i])
		Generate_table(tables[i], sql_script_file, template)
	}
}

func Update_variables(request Request) {
	for i := 0; i < len(request.Fields); i++ {
		if i != len(request.Fields) - 1{
			request.Fields[i].Field += ","
		}
	}
}
