package reader

import (
	"log"

	"github.com/alejandrowaiz98/garbage-reductor/excel"
)

type Reader struct {
	e excel.IExcel
}

type IReader interface {
	BuildExcel() error
}

func New() (IReader, error) {

	excel, err := excel.New()

	if err != nil {
		log.Printf("Err creating Reader: %v", err)
		return nil, err
	}

	R := Reader{e: excel}

	return &R, nil

}
