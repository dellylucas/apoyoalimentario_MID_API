package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/astaxie/beego"
)

//CrudPath - path de crud api
const CrudPath = "http://localhost:8086/v1/information/"

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

//CallCRUDAPI - Hace un llamado al MID API con la informacion Socioeconomica y obtiene el tipo de apoyo
func CallCRUDAPI(url string, trequest string, datajson interface{}) (string, error) {
	b := new(bytes.Buffer)
	if datajson != nil {
		json.NewEncoder(b).Encode(datajson)
	}
	client := &http.Client{}
	req, err := http.NewRequest(trequest, url, b)
	r, err := client.Do(req)
	if err != nil {
		beego.Error("error", err)
		return "na", err
	}
	defer r.Body.Close()
	contents, err := ioutil.ReadAll(r.Body)
	return strings.Replace(string(contents), "\"", "", -1), err
}
