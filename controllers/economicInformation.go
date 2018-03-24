package controllers

import (
	"apoyoalimentario_MID_API/models"
	"encoding/json"

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
// @Success 201 {int} models.Economic
// @router / [put]
func (c *TipoApoyoController) CalculateScore() {

	var InfoEconomic models.Economic

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &InfoEconomic)

	if err == nil {
		Result := models.GetResult(InfoEconomic)
		c.Data["json"] = Result
	} else {

		c.Data["json"] = err.Error()
	}

	c.ServeJSON()

}
