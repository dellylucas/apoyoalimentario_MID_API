package models

import (
	"ss_mid_api/golog"
	"strconv"
	"strings"
)

type Economic struct {
	Idc             string
	Estrato         string
	Ingresos        int
	SostePropia     string
	SosteHogar      string
	Nucleofam       string
	PersACargo      string
	EmpleadArriendo string
	ProvBogota      string
	Ciudad          string
	PobEspecial     string
	Discapacidad    string
	PatAlimenticia  string
	SerPiloPaga     string
	Sisben          string
	Periodo         int
	Semestre        int
	Matricula       int
	TipoSubsidio    string
	Tipoapoyo       string
	Estado          int
}

func GetResult(InfoEcono Economic) string {
	reglasbase := CargarReglasBase("AdministrativaContratacion")
	TypeApoyo := `resultado(` + strconv.Itoa(InfoEcono.Estado) + `,` +
		strings.ToLower(InfoEcono.Estrato) + `,` +
		strconv.Itoa(InfoEcono.Matricula) + `,` +
		strconv.Itoa(InfoEcono.Ingresos) + `,` +
		/* aca SMLV + `,` +*/
		strings.ToLower(InfoEcono.SostePropia) + `,` +
		strings.ToLower(InfoEcono.SosteHogar) + `,` +
		strings.ToLower(InfoEcono.Nucleofam) + `,` +
		strings.ToLower(InfoEcono.PersACargo) + `,` +
		strings.ToLower(InfoEcono.EmpleadArriendo) + `,` +
		strings.ToLower(InfoEcono.ProvBogota) + `,` +
		strings.ToLower(InfoEcono.PobEspecial) + `,` +
		strings.ToLower(InfoEcono.Discapacidad) + `,` +
		strings.ToLower(InfoEcono.PatAlimenticia) + `,Y).`
	Result := golog.GetOneString(reglasbase, TypeApoyo, "Y")

	return Result
}
