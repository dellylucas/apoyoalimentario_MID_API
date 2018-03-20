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
	/*	var reglasbase string
			reglasbase = `estadoPrograma(1).

		estrato(1,10).
		estrato(2,10).
		estrato(3,10).
		estrato(X,Y):-
			X @> 3,
			Y is 0 .

		matricula(X,Y):-
		    between(0,200000,X,20,Y);
		    between(200001,400000,X,16,Y);
		    between(400001,600000,X,12,Y);
		    between(600001,800000,X,8,Y);
		    between(800001,900000,X,4,Y);
		    between(900001,X,X,0,Y).

		ingresos(X,X1,Y):-
		    between(0,X1,X,30,Y);
		    between(X1,X1*2,X,20,Y);
		    between(X1*2,X1*3,X,10,Y);
		    between(X1*3,X1*4,X,5,Y);
		    between(X1*4,X,X,0,Y).

		between(V,W,X,Y,Z):-
		    X @>= V,
		    X @=< W,
		    Z is Y.
		betweenresultado(V,W,X,Y,Z):-
		    X @>= V,
		    X @=< W,
		    Z = Y.
		sostenimientoPropio(si,5).
		sostenimientoPropio(no,0).

		sostieneHogar(si,5).
		sostieneHogar(no,0).

		viveFueraNucleo(si,4).
		viveFueraNucleo(no,0).

		personasAcargo(si,6).
		personasAcargo(no,0).

		empleadorOarriendo(si,5).
		empleadorOarriendo(no,0).

		fueraDeBogota(si,5).
		fueraDeBogota(no,0).

		poblacionEspecial(X,Y):-
		    X \== n,
		    Y is 5.
		poblacionEspecial(n,0).

		discapacidad(si,5).
		discapacidad(no,0).

		patologiaAlimenticia(si,5).
		patologiaAlimenticia(no,0).

		suma(A,B,C,D,E,F,G,H,I,J,K,L,X):-
		    X is A+B+C+D+E+F+G+H+I+J+K+L.
		total(X,Y):-
		    betweenresultado(0,24,X,ss,Y);
		    betweenresultado(25,44,X,b,Y);
		    betweenresultado(45,59,X,a,Y);
		    betweenresultado(60,X,X,t,Y).

		resultado(AA,A,B,C,C1,D,E,F,G,H,I,J,K,L,Z):-
		    estadoPrograma(AA),
		   	estrato(A,Y1),
		    matricula(B,Y2),
		    ingresos(C,C1,Y3),
		    sostenimientoPropio(D,Y4),
		    sostieneHogar(E,Y5),
		    viveFueraNucleo(F,Y6),
		    personasAcargo(G,Y7),
		    empleadorOarriendo(H,Y8),
		    fueraDeBogota(I,Y9),
		    poblacionEspecial(J,Y10),
		    discapacidad(K,Y11),
		    patologiaAlimenticia(L,Y12),
		    suma(Y1,Y2,Y3,Y4,Y5,Y6,Y7,Y8,Y9,Y10,Y11,Y12,X),
		    total(X,Z).
		`
	*/ //carga de reglas desde el ruler
	var reglasbase string = ``
	var v []Predicado
	//	if err := getJson("http://10.20.0.254/ruler/v1/predicado/?query=dominio.Nombre:"+dominio+"&limit=-1", &v); err == nil {
	if err := getJson("http://localhost:5434/v1/predicado/?query=dominio.Nombre:"+dominio+"&limit=-1", &v); err == nil {

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
