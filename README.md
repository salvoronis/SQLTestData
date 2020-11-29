 # SQLTestData

SQLTestData is a program, which create test data to your database. You have to write your database scheme and using alphabet in json file. In scheme file ypu can write dependencies, min or max and strings options. Also you have to write N integer for standart count of cells in **one** table.

## JSON structure
JSON structure is an array of tables, which contains
   - Table_name for table name
   - Fields for array of fields in the table
   - Ratio for resize count of cells
In Fields you can put Field, which is a field's name, dependency, which will limin generator below last id in the dependency table, Min and Max, which will limit integer as you want and list of options
In Ratio you can resize count of cells in the table, it will resize your standart N, options:
 - " * "
 - " / "
 - " + "
 - " - "
```json
    [
        {
            "Table_name": "table",
            "Fields":[
                {
                    "Field": "field",
                    "Type": "string",
                    "Dependency": "Other_Table_Name",
                    "Min": 10,
                    "Max": 110,
                    "Options": [
                        "M",
                        "F"
                    ]
                }
            ],
            "Ratio":{
                "Operator": "*",
                "Number": 123
            }
        }
    ]
```

## Installation and running
You can install it using `make all`
The program at the moment doesn't support flags, so you can run it like `bin N`, where N is your standart count number

## Why we don't use standart alphabet?

  - Your can make your custom alphabet, which can sounds more naturalelly (see "japanese psewdo-words")
  - We will, but later


## Dependency
This program was written in Go language, so you have to install it first.

### Todos
If you have an idea how to make SQLTestData better, please write it into issues

 - standart alphabet
 - support more types
 - directly fill database

### License

GPL-3.0 License 

---
I use arch btw
