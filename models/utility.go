package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//getJSON - get Json format
func getJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

//CargarReglasBase since BD
func CargarReglasBase(dominio string) (reglas string) {

	var reglasbase string
	var v []Predicado
	if err := getJSON("http://10.20.0.254/ruler/v1/predicado/?query=dominio.Nombre:"+dominio+"&limit=-1", &v); err == nil {
		//if err := getJSON("http://localhost:5434/v1/predicado/?query=dominio.Nombre:"+dominio+"&limit=-1", &v); err == nil {

		reglasbase = reglasbase + FormatoReglas(v) //funcion general para dar formato a reglas cargadas desde el ruler
	} else {
		fmt.Println("err: ", err)
	}

	return reglasbase
}

//FormatoReglas - format of rules
func FormatoReglas(v []Predicado) (reglas string) {
	var arregloReglas = make([]string, len(v))
	reglas = ""
	//var respuesta []models.FormatoPreliqu
	for i := 0; i < len(v); i++ {
		arregloReglas[i] = v[i].Nombre
	}

	for i := 0; i < len(arregloReglas); i++ {
		reglas = reglas + arregloReglas[i] + "\n"
	}
	return
}
