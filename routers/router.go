package routers

import (
    "github.com/adairxie/blog/controllers"
    "github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
