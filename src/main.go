package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

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
	parameter   []Parameter
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
		tagH4Open := "<h4>"
		tagH4Close := "</h4>"

		indexH4Start := strings.Index(dataString, tagH4Open)
		if indexH4Start == -1 {
			break
		}

		indexEndOffsetByStart := strings.Index(dataString[indexH4Start:], tagH4Close)
		indexH4End := indexH4Start + indexEndOffsetByStart + len(tagH4Close)

		tagString := dataString[indexH4Start:indexH4End]

		// получение имени
		indexNameStart := strings.Index(tagString, "</a>") + 4
		indexNameEnd := strings.Index(tagString, "</h4>")
		name := tagString[indexNameStart:indexNameEnd]

		fmt.Println("Название: " + name)

		// получение описания
		indexDescStart := strings.Index(dataString, "<p>") + 3
		indexDescEndOffestByStart := strings.Index(dataString[indexDescStart:], "</p>")
		indexDescEnd := indexDescStart + indexDescEndOffestByStart
		desc := dataString[indexDescStart:indexDescEnd]
		fmt.Println("Описание: " + desc + "\n")

		nextH4 := strings.Index(dataString[indexH4End:], "<h4>")
		nextTbody := strings.Index(dataString[indexH4End:], "<tbody>")

		if nextH4 != -1 && nextTbody > nextH4 {
			dataString = dataString[nextH4:]
			continue
		}

		indexTbodyStart := strings.Index(dataString, "<tbody>") + 7
		indexTbodyOffsetByStart := strings.Index(dataString[indexTbodyStart:], "</tbody>")
		indexTbodyEnd := indexTbodyStart + indexTbodyOffsetByStart
		fields := dataString[indexTbodyStart:indexTbodyEnd]

		fmt.Println("Поля: " + fields)

		// получение полей
		if unicode.IsUpper(rune(name[0])) {
			typeTG := &Type{}
			typeTG.name = name
		} else {
			methodTG := &Method{}
			methodTG.name = name
		}

		dataString = dataString[indexTbodyEnd:]
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
