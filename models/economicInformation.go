package models

import (
	"apoyoalimentario_MID_API/golog"
	"strconv"
	"strings"
)

type Economic struct {
	Id                   string
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

func GetResult(InfoEcono Economic) string {
	reglasbase := CargarReglasBase("APOYOALIMENTARIO")
	TypeApoyo := `resultado(` + strconv.Itoa(InfoEcono.Estadoprograma) + `,` +
		strings.ToLower(InfoEcono.Estrato) + `,` +
		strconv.Itoa(InfoEcono.Matricula) + `,` +
		strconv.Itoa(InfoEcono.Ingresos) + `,` +
		`780000` + `,` +
		strings.ToLower(InfoEcono.Sostenibilidadpropia) + `,` +
		strings.ToLower(InfoEcono.Sostenibilidadhogar) + `,` +
		strings.ToLower(InfoEcono.Nucleofamiliar) + `,` +
		strings.ToLower(InfoEcono.Personasacargo) + `,` +
		strings.ToLower(InfoEcono.Empleadoroarriendo) + `,` +
		strings.ToLower(InfoEcono.Provienefuerabogota) + `,` +
		strings.ToLower(InfoEcono.Poblacionespecial) + `,` +
		strings.ToLower(InfoEcono.Discapacidad) + `,` +
		strings.ToLower(InfoEcono.Patologiaalimenticia) + `,Y).`
	Result := golog.GetOneString(reglasbase, TypeApoyo, "Y")

	return Result
}
