# SQLTestData

SQLTestData is a program, which create test data to your database. You have to write your database scheme and using alphabet in json file. In scheme file you can write dependencies, min or max and strings options. Also you have to write N integer for standart count of cells in **one** table.

## Supported types
- string
    - supports Options
- int
    - supports Min and Max options
    - supports Dependency
- boolean
- date

## JSON structure
JSON structure is an array of tables, which contains:
   - Table_name for table name
   - Fields for array of fields in the table
   - Ratio for resize count of cells
In Fields you can put Field name, which is a field's name, dependency, which will limit generator below last ID in the dependency (references) table, Min and Max, which will limit integer as you want, Options, which is a list of options
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
                "Operator": "*"
                "Number": "123"
            }
        }
    ]
```

## Installation and running
You can install it using `make all`
After all this programm have standart flags, so it will work, even without any flags, just put your scheme into the `table.json` file and alphabet into the `alphabet.json` file
The program supports flags, there they are:

> -a string
> JSON file containing alphabet (default "./alphabet.json")

>-n int
> A standart count of ceils (default 10)

>-o string
>Where to put result (default "./inserts.sql")

>-t string
>JSON file containing db scheme (default "./table.json")

>-h
>Help

## Why we don't use standart alphabet?

  - Your can make your custom alphabet, which can sounds more naturalelly
  - We will, but later


## Dependency
This program was written in Go language, so you have to install it first.

### Todos
If you have an idea how to make SQLTestData better, please write it into issues

 - standart alphabet (english)
 - support more types
 - directly fill database

### License

GPL-3.0 License 

---
I use arch btw
