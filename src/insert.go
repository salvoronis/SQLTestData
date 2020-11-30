package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"text/template"
	"time"
)

var (
	logfile, _ = os.Create("./log.txt")
	logger = log.New(logfile, "db_test_generator", log.LstdFlags|log.Lshortfile)
	alphabet []string = []string{}
	tables []Request = []Request{}
	dependencies map[string]int = make(map[string]int)
	n int
)

func main() {
	var table_file, alphabet_file, output_file string
	Init_flag(&table_file,&alphabet_file, &output_file)

	var template_str string = "insert into {{.Table_name}}({{range .Fields}}{{.Field}}{{end}}) values {{range .Data}}\n\t{{.}}{{end}}\n"
	template, err := template.New("sql_test_inserts").Parse(template_str)
	if err != nil {
		logger.Fatalln(err)
	}

	rand.Seed(time.Now().UnixNano())

	Parse_alphabet(alphabet_file)
	Parse_scheme(table_file)

	sql_script_file, err := os.Create(output_file)
	if err != nil {
		logger.Fatalln(err)
	}

	for i := 0; i < len(tables); i++{
		Update_variables(tables[i])
		Generate_table(tables[i], sql_script_file, template)
	}
}

func Init_flag(table, alphabet, output *string){
	flag.IntVar(&n ,"n", 10, "A standart count of ceils")
	flag.StringVar(table, "t", "./table.json", "JSON file containing db scheme")
	flag.StringVar(alphabet, "a", "./alphabet.json", "JSON file containing alphabet")
	flag.StringVar(output, "o", "./inserts.sql", "Where to put result")
	flag.Parse()
}

func Update_variables(request Request) {
	for i := 0; i < len(request.Fields); i++ {
		if i != len(request.Fields) - 1{
			request.Fields[i].Field += ","
		}
	}
}
