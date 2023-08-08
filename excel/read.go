package excel

import (
	"log"
	"os"

	"github.com/xuri/excelize/v2"
)

func (e *Excel) Start(garbageFile, newFile *excelize.File) (data [][]string, len int) {

	rows, err := garbageFile.GetRows(os.Getenv("sheet_name"))

	if err != nil {
		log.Printf("Err getting rows: %v", err)
		return nil, 0
	}

	var columnNames []string

	for i, columns := range rows {

		if i == 0 {
			columnNames = columns
		} else {
			break
		}

	}

	var columnPosition []int

	for i, name := range columnNames {

		if contains(columnNames, name) {

			columnPosition = append(columnPosition, i)

		}

	}

	return

}

func contains(names []string, wanted string) bool {
	for _, name := range names {
		if name == wanted {
			return true
		}
	}
	return false
}
