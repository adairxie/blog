package models

import (
    "fmt"
    "log"
    "strconv"
    "github.com/fatih/structs"
    "github.com/go-redis/redis"
)

type Article struct {
    Id              int64
    Title           string
    SubTitle        string
    Category        string
    Content         string
    Created         string
    Updated         string
    Views           int64
    Author          string
}

func MapToStruct(articleMap map[string]string) *Article {

    article := &Article{}
    id := articleMap["Id"]
    title := articleMap["Title"]
    subTitle := articleMap["SubTitle"]
    content := articleMap["Content"]
    created := articleMap["Created"]
    updated := articleMap["Updated"]
    author := articleMap["Author"]
    views := articleMap["Views"]

    if id != "" {
       article.Id, _ = strconv.ParseInt(id, 10, 64)
    }

    if title != "" {
        article.Title = title
    }

    if subTitle != "" {
        article.SubTitle = subTitle
    }

    if content != "" {
        article.Content = content
    }

    if created != "" {
        article.Created = created
    }

    if updated != "" {
        article.Updated = updated
    }

    if author != "" {
        article.Author = author
    }

    if views != "" {
        article.Views, _= strconv.ParseInt(views, 10, 64)
    }

    return article
}

func NewClient() *redis.Client{
    client := redis.NewClient(&redis.Options{
        Addr:       "localhost:6379",
        Password:   "foobar",
        DB:         0,
    })

    return client
}

func GetArticleById(id int64) *Article {
    client := NewClient()
    key := fmt.Sprintf("article:%d", id)

    pipe := client.Pipeline()
    pipe.HIncrBy(key, "Views", 1)
    articleMap := pipe.HGetAll(key)
    _, err := pipe.Exec()
    if err != nil {
        log.Fatalf("Get Article failed, article'id: %d", id)
    }

    return MapToStruct(articleMap.Val())
}

func GetAllArticles() ([]*Article, error) {
    client := NewClient()
    ids, err := GetArticlesByRange(client, 0, -1)
    if err != nil {
        log.Fatalf("%s", err)
    }
    //log.Printf("Articles'ids : %s", ids)

    var articleMap []*redis.StringStringMapCmd
    pipe := client.Pipeline()

    for _, id := range ids {
        articlePtr := pipe.HGetAll(fmt.Sprintf("article:%s", id))
        //log.Printf("pipe.HGetAll: %s", articleMap)
        articleMap = append(articleMap, articlePtr)
    }

    _, err = pipe.Exec()
    if err != nil {
        log.Printf("Redis pipeline exec cmd failed: %s", err)
    }

    var articles []*Article
    for _, articlePtr := range articleMap {
        article := MapToStruct(articlePtr.Val())
        articles = append(articles, article)
    }
    return articles, err
}

func GenerateArticleId() (int64, error) {
    client := NewClient()
    return client.Incr("article:count").Result()
}

func GetArticlesByRange(client *redis.Client, start, end int64) ([]string, error){
    res, err := client.LRange("article:ids", start, end).Result()
    return res, err
}

func AddArticle(post *Article) error {
    client := NewClient()
    id, err := GenerateArticleId()
    if err != nil {
        panic(err)
    }
    post.Id = int64(id)
    err = client.LPush("article:ids", id).Err()
    if err != nil  {
        return err
    }

    key := fmt.Sprintf("article:%d", id)
    err = client.HMSet(key, structs.Map(post)).Err()
    return err
}

func ModifyArticle(fields map[string]interface{}) error {
    client := NewClient()
    key := fmt.Sprintf("article:%d", fields["Id"])
    err := client.HMSet(key, fields).Err()
    return err
}

func DeleteArticle(id int64) error {
    client := NewClient()
    pipe := client.Pipeline()
    // article:count sub 1
    pipe.Decr("article:count")
    // delete article:id 
    key := fmt.Sprintf("article:%d", id)
    pipe.Del(key)

    // delete id from article:ids
    res := pipe.LRem("article:ids", 1, id)
    _, err := pipe.Exec()
    if res.Val() != 1 {
        log.Printf("delete article failed, id: %d", id)
    }

    return err

}
