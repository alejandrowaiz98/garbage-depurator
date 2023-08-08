package excel

type Excel struct {
}

type IExcel interface {
}

func New() IExcel {
	return &Excel{}
}
