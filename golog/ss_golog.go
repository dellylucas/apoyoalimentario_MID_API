package golog

import (
	"fmt"

	. "github.com/mndrix/golog"
)

//GetOneString - get one variable type string
func GetOneString(reglas string, reglaInyectada string, variableObtener string) string {
	m := NewMachine().Consult(reglas)
	resultados := m.ProveAll(reglaInyectada)
	var returnedString string
	returnedString = ""
	for _, solution := range resultados {
		aux := fmt.Sprintf("%s", solution.ByName_(variableObtener))
		returnedString = aux
	}
	return returnedString
}
