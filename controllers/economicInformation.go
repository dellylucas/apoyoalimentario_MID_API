package controllers

import (
	"apoyoalimentario_MID_API/models"
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
)

// TipoApoyoController services operations for result Type apoyo
type TipoApoyoController struct {
	beego.Controller
}

// URLMapping ...
func (c *TipoApoyoController) URLMapping() {
	c.Mapping("ObtenerInfoCoordinador", c.Put)

}

// CalculateScore - calculate score of student
// @Title CalculateScore
// @Description Evaluate rules for scrore of stundent
// @Param	smlv	path 	string	true	"salario minimo configurado por el administrador"
// @Param	code	path 	string	true	"codigo del estudiante"
// @Success 200 {object} models.Object
// @Failure 403 :smlv is empty
// @Failure 403 :code is empty
// @router /:smlv/:code [put]
func (c *TipoApoyoController) CalculateScore() {
	salario := c.Ctx.Input.Param(":smlv")
	var InfoEconomic models.Economic
	json.Unmarshal(c.Ctx.Input.RequestBody, &InfoEconomic)
	InfoEconomic.Codigo = c.Ctx.Input.Param(":code")
	Result := models.GetResult(&InfoEconomic, salario)
	if strings.Compare(Result, "na") == 0 {
		c.Data["json"] = 0 //error en mid api
	} else {
		c.Data["json"] = 1 // ruler ok
	}

	c.ServeJSON()
}
