package main

import (
	"math/rand"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func Generate_table(request Request, file *os.File, tmp *template.Template){
	request.Count = Count_comp(request.Ratio)

	dependencies[request.Table_name] = request.Count

	table_uniques := make(map[string]int)

	var b strings.Builder
	for i := 0; i < request.Count; i++ {
		b.WriteString("(")

		var str string
		notEquals := make(map[string]string)

		for j := 0; j < len(request.Fields); j++ {

			data := Gen_data(request.Fields[j])

			if request.Fields[j].NotEqualID != "" {
				var fuse int = 0
				for data == notEquals[request.Fields[j].NotEqualID] {
					fuse++
					if fuse > 100 {
						logger.Fatalln("can not create not equal value")
					}
					data = Gen_data(request.Fields[j])
				}
				notEquals[request.Fields[j].NotEqualID] = data
			}

			if request.Fields[j].Unique {
				var fuse int = 0
				for table_uniques[request.Fields[j].Field+data] != 0 {
					fuse++
					data = Gen_data(request.Fields[j])
					if fuse > 100 {
						logger.Fatalln("can not create unique value")
					}
				}
				table_uniques[request.Fields[j].Field+data] ++
			}

			str += data

			if j != len(request.Fields)-1 {
				str += ","
			}

		}
		b.WriteString(str + ")")

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

func Gen_data(field Field_container) string{
	switch field.Type {
	case "string":
		return "'" + String_type(field) + "'"
	case "int":
		return Itn_type(field)
	case "boolean":
		return Boolean_type(field)
	case "date":
		return Date_type(field)
	}
	return ""
}

func Date_type(field Field_container) string {
	layoutISO := "2006-01-02"
	min := 1000
	max := 9999
	if field.Max != 0 {
		max = field.Max
	}
	if field.Min != 0 && field.Min >= 1000 {
		min = field.Min
	}
	t := time.Date(min+rand.Intn(max-min), time.Month(1+rand.Intn(12)), 1+rand.Intn(28),0,0,0,0,time.UTC)
	return "'" + t.Format(layoutISO) + "'"
}

func String_type(field Field_container) string{
	if len(field.Options) != 0{
		return field.Options[rand.Intn(len(field.Options))]
	}
	return Gen_rand_string(2+rand.Intn(4))
}

func Itn_type(field Field_container) string {
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
}

func Boolean_type(field Field_container) string{
	if rand.Intn(2) == 1 {
		return "true"
	}
	return "false"
}

func Gen_rand_string(size int) string {
	var b strings.Builder
	for i := 0; i < size; i++ {
		b.WriteString(alphabet[rand.Intn(len(alphabet))])
	}
	return b.String()
}
