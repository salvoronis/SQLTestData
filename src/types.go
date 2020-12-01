package main

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
	Unique bool;
	NotEqualID string;
}
