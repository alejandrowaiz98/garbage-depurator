package excel

import (
	"log"

	"github.com/xuri/excelize/v2"
)

type Excel struct {
	finalFile *excelize.File
}

type IExcel interface {
	BuildFinalExcel() error
}

func New() (IExcel, error) {

	f := excelize.NewFile()

	err := f.SaveAs("Final.xlsx")

	if err != nil {

		log.Printf("Err creating new excel file: %v", err)

		return nil, err

	}

	return &Excel{finalFile: f}, nil
}
