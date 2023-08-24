package constants

var Letters []string = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD", "AE", "AF", "AG"}

var WantedNames []string = []string{"Nombre", "RUT", "PUNTOS"}

var WantedNamesWithLetter map[string]string

func init() {

	WantedNamesWithLetter = make(map[string]string)
	WantedNamesWithLetter["NOMBRE"] = "A"
	WantedNamesWithLetter["RUT"] = "B"
	WantedNamesWithLetter["CARGO"] = "C"
	WantedNamesWithLetter["TIPO DE CONTRATO"] = "D"
	WantedNamesWithLetter["SISTEMA DE SALUD"] = "E"
	WantedNamesWithLetter["CENTRO COSTO"] = "F"

}
