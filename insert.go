package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type Request struct{
	Table_name string;
	Fields []Field_container;
	Data []string;
	Ratio Rati;
	Count int;
}

type Rati struct{
	Operator string;
	Number int;
}

type Field_container struct{
	Field string;
	Type string;
	Dependent string;
	Options []string;
	Max, Min int;
}

var (
	logfile, _ = os.Open("./log.txt")
	logger = log.New(logfile, "db_test_generator", log.LstdFlags|log.Lshortfile)
	alphabet []string = []string{}
	tables []Request = []Request{}
	dependencies map[string]int = make(map[string]int)
	n int
)

func main() {
	var err error
	n, err = strconv.Atoi(os.Args[1])
	if err != nil {
		logger.Println(err)
		n = 100
	}
	rand.Seed(time.Now().UnixNano())
	var template_str string = "insert into {{.Table_name}}({{range .Fields}}{{.Field}}{{end}}) values {{range .Data}}\n\t{{.}}{{end}}\n"

	Parse_alphabet("./alphabet.json")
	Parse_scheme("./table.json")

	sql_script_file, err := os.Create("./inserts.sql")
	if err != nil {
		logger.Println(err)
	}

	template, err := template.New("sql_test_inserts").Parse(template_str)
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

func Generate_table(request Request, file *os.File, tmp *template.Template){
	request.Count = Count_comp(request.Ratio)

	dependencies[request.Table_name] = request.Count

	var b strings.Builder
	for i := 0; i < request.Count; i++ {
		b.WriteString("(")
		for j := 0; j < len(request.Fields); j++ {
			b.WriteString(Gen_data(request.Fields[j]))
			if j != len(request.Fields)-1 {
				b.WriteString(",")
			}
		}
		b.WriteString(")")
		if i != request.Count - 1{
			b.WriteString(",\n\t")
		} else {
			b.WriteString(";\n")
		}
	}
	request.Data = append(request.Data, b.String())

	Write_data(file, tmp, &request)
}

func Count_comp(req Rati) int {
	if req == (Rati{}) {
		return n
	}
	switch req.Operator{
	case "/":
		return n/req.Number
	case "*":
		return n*req.Number
	case "+":
		return n+req.Number
	case "-":
		return n-req.Number
	}
	return 0
}

func Write_data(file *os.File, tmp *template.Template, data *Request){
	err := tmp.Execute(file, data)
	if err != nil {
		logger.Println(err)
	}
}

func Gen_data(field Field_container) string{
	switch field.Type {
	case "string":
		if len(field.Options) != 0{
			return field.Options[rand.Intn(len(field.Options))]
		}
		var str string = "'"
		str += Gen_rand_string(2+rand.Intn(4))
		str += "'"
		return str
	case "int":
		if field.Dependent != "" {
			depend := dependencies[field.Dependent]
			return strconv.Itoa(1+rand.Intn(depend))
		}
		max := 10
		min := 0
		if field.Max != 0 {
			max = field.Max
		}
		if field.Min != 0 {
			min = field.Min
		}
		return strconv.Itoa(min+rand.Intn(max))
	case "boolean":
		if rand.Intn(2) == 1 {
			return "true"
		}
		return "false"
	}
	return ""
}

func Gen_rand_string(size int) string {
	var b strings.Builder
	for i := 0; i < size; i++ {
		b.WriteString(alphabet[rand.Intn(len(alphabet))])
	}
	return b.String()
}

func Parse_alphabet(alphabet_path string){
	alphabet_file, err := os.Open(alphabet_path)
	if err != nil {
		logger.Println(err)
	}
	alphabet_byte, err := ioutil.ReadAll(alphabet_file)
	if err != nil {
		logger.Println(err)
	}
	alphabet_file.Close()
	json.Unmarshal(alphabet_byte, &alphabet)
}

func Parse_scheme(scheme_path string){
	tables_file, err := os.Open("table.json")
	if err != nil {
		logger.Println(err)
	}
	tables_byte, err := ioutil.ReadAll(tables_file)
	if err != nil {
		logger.Println(err)
	}
	tables_file.Close()
	json.Unmarshal(tables_byte, &tables)
}
