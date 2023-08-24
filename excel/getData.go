package excel

import (
	"fmt"
	"log"
	"os"

	"github.com/alejandrowaiz98/garbage-reductor/constants"
	"github.com/xuri/excelize/v2"
)

var LastColumnIndex int

func (e *Excel) BuildFinalExcel() error {

	log.Println(LastColumnIndex)

	log.Println("opening file...")

	garbageFile, err := excelize.OpenFile(os.Getenv("excel_filename"))
	defer garbageFile.Close()

	if err != nil {
		log.Printf("Err opening file: %v", err)
		return nil
	}

	log.Println("getting rows...")

	columns, err := garbageFile.GetRows("Hoja1")

	if err != nil {
		log.Printf("Err getting rows: %v", err)
		return nil
	}

	cols, err := garbageFile.GetCols("Hoja1")

	if err != nil {
		log.Printf("Err getting cols: %v", err)
		return nil
	}

	for i, column := range columns {

		if i == 0 {
			continue
		}

		if i == 1 {

			log.Printf("columnssss: %v", column)

			for j, columnName := range column {

				if contains(constants.WantedNames, columnName) {

					log.Printf("Inserting column: %v", column)

					cellCoord := fmt.Sprintf("%v%v", constants.WantedNamesWithLetter[columnName], LastColumnIndex)

					log.Printf("Into: %v", cellCoord)

					column = cols[j]

					finalCol := remove(column, 0)

					err = e.finalFile.SetSheetCol("Sheet1", cellCoord, &finalCol)

					if err != nil {
						log.Printf("err setting shell in coord %v:%v", cellCoord, err)
					}

				}

			}

			LastColumnIndex = LastColumnIndex + len(column) - 1

		} else {

			break
		}

	}

	if err != nil {
		return err
	}

	err = e.finalFile.Save()

	if err != nil {
		return err
	}

	return nil
}

func contains(names []string, wanted string) bool {
	for _, name := range names {
		if name == wanted {
			return true
		}
	}
	return false
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
