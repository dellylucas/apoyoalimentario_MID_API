package models

import (
	"apoyoalimentario_MID_API/golog"
	"strconv"
	"strings"
)

//Economic - model of information of student
type Economic struct {
	Codigo               string
	Estrato              string
	Matricula            int
	Ingresos             int
	Sostenibilidadpropia string
	Sostenibilidadhogar  string
	Nucleofamiliar       string
	Personasacargo       string
	Empleadoroarriendo   string
	Provienefuerabogota  string
	Poblacionespecial    string
	Discapacidad         string
	Patologiaalimenticia string
	Tiposubsidio         string
	Estadoprograma       int
}

//GetResult - Get Result of type apyo of student
func GetResult(InfoEconomic *Economic, salario string) string {

	reglasbase := CargarReglasBase("APOYOALIMENTARIO")
	TypeApoyo := `resultado(` +
		strconv.Itoa(InfoEconomic.Estadoprograma) + `,` +
		strings.ToLower(InfoEconomic.Estrato) + `,` +
		strconv.Itoa(InfoEconomic.Matricula) + `,` +
		strconv.Itoa(InfoEconomic.Ingresos) + `,` +
		salario + `,` +
		strings.ToLower(InfoEconomic.Sostenibilidadpropia) + `,` +
		strings.ToLower(InfoEconomic.Sostenibilidadhogar) + `,` +
		strings.ToLower(InfoEconomic.Nucleofamiliar) + `,` +
		strings.ToLower(InfoEconomic.Personasacargo) + `,` +
		strings.ToLower(InfoEconomic.Empleadoroarriendo) + `,` +
		strings.ToLower(InfoEconomic.Provienefuerabogota) + `,` +
		strings.ToLower(InfoEconomic.Poblacionespecial) + `,` +
		strings.ToLower(InfoEconomic.Discapacidad) + `,` +
		strings.ToLower(InfoEconomic.Patologiaalimenticia) +
		`,Y).`

	Result := golog.GetOneString(reglasbase, TypeApoyo, "Y")
	if strings.Compare(Result, "") != 0 {
		InfoEconomic.Tiposubsidio = Result
		InfoEconomic.Estadoprograma = 2
	} else {
		Result = "na"
		InfoEconomic.Estadoprograma = 0
	}
	CallCRUDAPI(CrudPath+InfoEconomic.Codigo, "PUT", InfoEconomic)
	return Result
}
