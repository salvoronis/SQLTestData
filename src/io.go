package main

import(
	"encoding/json"
	"os"
	"io/ioutil"
	"text/template"
)

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

func Write_data(file *os.File, tmp *template.Template, data *Request){
	err := tmp.Execute(file, data)
	if err != nil {
		logger.Println(err)
	}
}
