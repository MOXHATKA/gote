package main

import (
	"fmt"
	"os"
	"strings"
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

	var result string

	for {
		block := getInnerData(dataString, "<h4>", "<h4>")
		if block == nil {
			break
		}
		name := getInnerData(block.data, "</a>", "</h4>")
		result += name.data + "\n"
		// fmt.Println(name.data)

		description := getInnerData(block.data, "<p>", "</p>")
		result += description.data + "\n"
		// fmt.Println(description.data)

		table := getInnerData(block.data, "<tbody>", "</tbody>")
		if table == nil {
			dataString = dataString[block.indexEnd:]
			continue
		}

		for {
			row := getInnerData(table.data, "<tr>", "</tr>")
			if row == nil {
				break
			}

			var rowString string
			for {
				cell := getInnerData(row.data, "<td>", "</td>")
				if cell == nil {
					break
				}
				rowString += cell.data + "\n"
				row.data = row.data[cell.indexEnd:]
			}
			result += rowString + "\n"

			table.data = table.data[row.indexEnd:]
		}

		dataString = dataString[block.indexEnd:]
	}

	f, err := os.Create("./output/result.html")
	if err != nil {
		panic("Не получилось создать файл")
	}

	_, err = f.WriteString(result)
	if err != nil {
		panic("Не получилось записать в файл")
	}
}

type InnerDataResult struct {
	indexEnd int
	data     string
}

func getInnerData(dataString, tagStart, tagEnd string) *InnerDataResult {
	indexStart := strings.Index(dataString, tagStart)
	if indexStart == -1 {
		return nil
	}
	indexStart += len(tagStart)

	var indexEnd int
	if tagEnd == "" {
		indexEnd = len(dataString)
	} else {
		indexOffset := strings.Index(dataString[indexStart:], tagEnd)
		if indexOffset == -1 {
			return nil
		}

		indexEnd = indexStart + indexOffset
	}

	result := &InnerDataResult{
		indexEnd: indexEnd,
		data:     dataString[indexStart:indexEnd],
	}

	return result
}

func printType(typeTG Type) {
	fmt.Println("Название: " + typeTG.name)
	fmt.Println("Описание: " + typeTG.description)
	for _, f := range typeTG.fields {
		fmt.Print("field: ", f.name, ", ")
		fmt.Print("type: ", f.typeField, ", ")
		fmt.Print("desc: ", f.description, "\n")
	}
	fmt.Println()
}

func printMethod(methodTG Method) {
	fmt.Println("Название: " + methodTG.name)
	fmt.Println("Описание: " + methodTG.description)
	for _, f := range methodTG.parameters {
		fmt.Print("field: ", f.name, ", ")
		fmt.Print("type: ", f.typeField, ", ")
		fmt.Print("required: ", f.required, ", ")
		fmt.Print("desc: ", f.description, "\n")
	}
	fmt.Println()
}
