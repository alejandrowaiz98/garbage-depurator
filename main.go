package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func main() {

	f, err := excelize.OpenFile("test.xlsx")

	if err != nil {
		log.Println(err)
		return
	}

	cols, err := f.GetCols("Hoja1")

	if err != nil {
		log.Println(err)
		return
	}

	newF := excelize.NewFile()

	for _, col := range cols {

		cellCoord := fmt.Sprintf("%v%v", GetCellLetter(col[0]), strconv.Itoa(len(cols)))

		log.Printf("cell coord: %v", cellCoord)

		finalCol := remove(col, 0)

		err = newF.SetSheetCol("Sheet1", cellCoord, &finalCol)

		if err != nil {
			log.Printf("err setting: %v", err)
		}

	}

	err = newF.SaveAs("final.xlsx")

	if err != nil {
		log.Println("err saving", err)
		return
	}

}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func GetCellLetter(columnName string) string {
	if columnName == "Rut" {
		return "A"
	} else if columnName == "Name" {
		return "B"
	} else if columnName == "Age" {
		return "C"
	}
	return ""
}
