package models

import (
    "fmt"
    "log"
    "time"
    "github.com/fatih/structs"
    "github.com/go-redis/redis"
    "github.com/mitchellh/mapstructure"
)

type Article struct {
    Id              int64
    Title           string
    SubTitle        string
    Category        string
    Content         string
    Created         time.Time
    Updated         time.Time
    Views           int64
    Author          string
}

func NewClient() *redis.Client{
    client := redis.NewClient(&redis.Options{
        Addr:       "localhost:6379",
        Password:   "foobar",
        DB:         0,
    })

    return client
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
        article := &Article{}
        mapstructure.Decode(articlePtr.Val(), article)
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
    post.Id = id
    err = client.LPush("article:ids", id).Err()
    if err != nil  {
        return err
    }

    key := fmt.Sprintf("article:%d", id)
    err = client.HMSet(key, structs.Map(post)).Err()
    return err
}


