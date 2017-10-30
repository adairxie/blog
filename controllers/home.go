package controllers

import (
    "log"
    "github.com/adairxie/blog/models"
    "github.com/astaxie/beego"
)

type MainController struct {
    beego.Controller
}

func (this *MainController) Get() {
    this.TplName = "index.html"

    /*article := &models.Article{
        Title:"Man must explore, and this is exploration at its greatest",
        SubTitle:"Problems look mighty small from 150 miles up",
        Created: time.Now(),
    }

    err := models.AddArticle(article)
    if err != nil {
        log.Fatalf("Add article failed: %s", err)
    }*/

    articles, err := models.GetAllArticles()
    if err != nil {
        log.Fatalf("Get all articles err: %s", err)
    }

    this.Data["Articles"] = articles
}
