package controllers

import (
    "log"
    "strconv"
    "github.com/adairxie/blog/models"
    "github.com/astaxie/beego"
)

type PostController struct {
    beego.Controller
}

func (this *PostController) Get() {
    this.TplName = "post.html"

    id, err := strconv.Atoi(this.Input().Get("article_id"))
    //根据article_id获取文章
    article := models.GetArticleById(int64(id))
    if err != nil {
        log.Fatalf("Get all articles err: %s", err)
    }

    this.Data["Article"] = article
}
