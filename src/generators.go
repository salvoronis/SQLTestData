package main

import(
	"math/rand"
	"strconv"
	"strings"
	"os"
	"text/template"
)

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

func Gen_data(field Field_container) string{
	switch field.Type {
	case "string":
		return String_type(field)
	case "int":
		return Itn_type(field)
	case "boolean":
		return Boolean_type(field)
	}
	return ""
}

func String_type(field Field_container) string{
	if len(field.Options) != 0{
		return field.Options[rand.Intn(len(field.Options))]
	}
	var str string = "'"
	str += Gen_rand_string(2+rand.Intn(4))
	str += "'"
	return str
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
