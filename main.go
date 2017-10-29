package main

import (
    "github.com/adairxie/blog/controllers"
    "github.com/astaxie/beego"
)

func main() {

    //register router
    beego.Router("/", &controllers.HomeController{})

    //start beego
    beego.Run()
}
