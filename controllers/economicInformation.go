package controllers

import (
	"apoyoalimentario_MID_API/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// TipoApoyoController operations for result Type apoyo
type TipoApoyoController struct {
	beego.Controller
}

// URLMapping ...
func (c *TipoApoyoController) URLMapping() {
	c.Mapping("ObtenerInfoCoordinador", c.Put)

}

// Put ...
// @Title Put
// @Description Evaluate rules
// @Param code query string true "Estudiante al cual se le aplica reglas"
// @Success 201 {int} models.Economic
// @router / [put]
func (j *TipoApoyoController) Put() {

	//Codigo := this.Ctx.Input.Param(":code")

	var InfoEcono models.Economic

	err := json.Unmarshal(j.Ctx.Input.RequestBody, &InfoEcono)

	if err == nil {
		str := models.GetResult(InfoEcono)
		j.Data["json"] = str
	} else {

		j.Data["json"] = err.Error()
	}

	j.ServeJSON()

}
