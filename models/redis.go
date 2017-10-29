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

    var articles []*Article

    for _, id := range ids {
        articleMap, err := client.HGetAll(fmt.Sprintf("article:%s", id)).Result()
        if err != nil {
            log.Fatalf("Cann't get %s's data, err: %s", id, err)
        }

        article := &Article{}
        mapstructure.Decode(articleMap, article)
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


