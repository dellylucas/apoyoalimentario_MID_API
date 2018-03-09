package main

import (
	_ "apoyoalimentario_MID_API/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDataBase("amazonAdmin", "postgres", "postgres://"+beego.AppConfig.String("UsercrudAgora")+":"+beego.AppConfig.String("PasscrudAgora")+"@"+beego.AppConfig.String("HostcrudAgora")+"/"+beego.AppConfig.String("BdcrudAgora")+"?sslmode=disable&search_path="+beego.AppConfig.String("SchcrudAgora")+"")
	orm.RegisterDataBase("flywayAdmin", "postgres", "postgres://"+beego.AppConfig.String("UsercrudAdmin")+":"+beego.AppConfig.String("PasscrudAdmin")+"@"+beego.AppConfig.String("HostcrudAdmin")+"/"+beego.AppConfig.String("BdcrudAdmin")+"?sslmode=disable&search_path="+beego.AppConfig.String("SchcrudAdmin")+"")
	orm.RegisterDataBase("default", "postgres", "postgres://"+beego.AppConfig.String("UsercrudAgora")+":"+beego.AppConfig.String("PasscrudAgora")+"@"+beego.AppConfig.String("HostcrudAgora")+"/"+beego.AppConfig.String("BdcrudAgora")+"?sslmode=disable&search_path="+beego.AppConfig.String("SchcrudAgora")+"")
}

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
