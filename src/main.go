package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

var index int

type Type struct {
	name        string
	description string
	note        string
	fields      []Field
}

type Field struct {
	name        string
	typeField   string
	description string
}

type Method struct {
	name        string
	description string
	note        string
	parameters  []Parameter
}

type Parameter struct {
	name        string
	typeField   string
	required    bool
	description string
}

func main() {
	data, err := os.ReadFile("./input/doca.html")
	if err != nil {
		panic("Не могу прочитать файл")
	}
	dataString := string(data)

	// var result string

	for {
		// получение <h4> ... </h4>
		var tagString string
		index, tagString = getInnerData(dataString, "<h4>", "</h4>")
		if index == -1 {
			break
		}

		indexH4End := index

		// получение имени
		var name string
		index, name = getInnerData(tagString, "</a>", "")
		fmt.Println("Название: " + name)

		// получение описания
		var desc string
		index, desc = getInnerData(dataString, "<p>", "</p>")
		fmt.Println("Описание: " + desc + "\n")

		// проверка на наличие таблицы
		nextH4 := strings.Index(dataString[indexH4End:], "<h4>")
		nextTbody := strings.Index(dataString[indexH4End:], "<tbody>")
		if nextH4 != -1 && nextTbody > nextH4 {
			dataString = dataString[nextH4:]
			continue
		}

		// получение полей
		var fields string
		index, fields = getInnerData(dataString, "<tbody>", "</tbody>")
		// fmt.Println("Поля: " + fields)

		// создание структур Type | Method
		if unicode.IsUpper(rune(name[0])) {
			typeTG := Type{}
			typeTG.name = name
			typeTG.description = desc

			for {
				var i int

				var name string
				i, name = getInnerData(fields, "<td>", "</td>")
				if i == -1 {
					break
				}
				fields = fields[i:]

				var typeField string
				i, typeField = getInnerData(fields, "<td>", "</td>")

				if i == -1 {
					break
				}
				fields = fields[i:]

				var desc string
				i, desc = getInnerData(fields, "<td>", "</td>")

				if i == -1 {
					break
				}

				field := Field{
					name:        name,
					typeField:   typeField,
					description: desc,
				}

				fmt.Print("field: ", field.name, ", ")
				fmt.Print("type: ", field.typeField, ", ")
				fmt.Print("desc: ", field.description, "\n")
				fields = fields[i:]
			}

		} else {
			methodTG := Method{}
			methodTG.name = name
		}

		dataString = dataString[index:]
	}

	// f, err := os.Create("./output/result.html")
	// if err != nil {
	// 	panic("Не получилось создать файл")
	// }

	// _, err = f.WriteString(result)
	// if err != nil {
	// 	panic("Не получилось записать в файл")
	// }
}

func getInnerData(dataString, tagStart, tagEnd string) (int, string) {
	indexStart := strings.Index(dataString, tagStart) + len(tagStart)
	if indexStart-len(tagStart) == -1 {
		return -1, ""
	}
	var indexEnd int
	if tagEnd == "" {
		indexEnd = len(dataString)
	} else {
		indexOffset := strings.Index(dataString[indexStart:], tagEnd)
		indexEnd = indexStart + indexOffset
	}

	return indexEnd, dataString[indexStart:indexEnd]
}
