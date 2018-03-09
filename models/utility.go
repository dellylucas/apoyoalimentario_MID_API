package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
)

func sendJson(url string, trequest string, target interface{}, datajson interface{}) error {
	b := new(bytes.Buffer)
	if datajson != nil {
		json.NewEncoder(b).Encode(datajson)
	}
	client := &http.Client{}
	req, err := http.NewRequest(trequest, url, b)
	r, err := client.Do(req)
	//r, err := http.Post(url, "application/json; charset=utf-8", b)
	if err != nil {
		beego.Error("error", err)
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func getJsonWSO2(urlp string, target interface{}) error {
	b := new(bytes.Buffer)
	//proxyUrl, err := url.Parse("http://10.20.4.15:3128")
	//http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	client := &http.Client{}
	req, err := http.NewRequest("GET", urlp, b)
	req.Header.Set("Accept", "application/json")
	r, err := client.Do(req)
	//r, err := http.Post(url, "application/json; charset=utf-8", b)
	if err != nil {
		beego.Error("error", err)
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

//CargarReglasBase general
func CargarReglasBase(dominio string) (reglas string) {
	//carga de reglas desde el ruler
	var reglasbase string = ``
	var v []Predicado
	if err := getJson("http://10.20.0.254/ruler/v1/predicado/?query=Dominio.Nombre:"+dominio+"&limit=-1", &v); err == nil {

		reglasbase = reglasbase + FormatoReglas(v) //funcion general para dar formato a reglas cargadas desde el ruler
	} else {
		fmt.Println("err: ", err)
	}

	return reglasbase
}

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
