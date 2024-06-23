package main

import (
	"flag"
	"fmt"

	"SyaibanAhmadRamadhan/web-scrapping-with-flag-pagination-golang/infrastructures/db"
	"SyaibanAhmadRamadhan/web-scrapping-with-flag-pagination-golang/repositories"
	"SyaibanAhmadRamadhan/web-scrapping-with-flag-pagination-golang/services"
)

func init() {
}

var (
	maxPost   int
	maxPaging int
)

func init() {
	flag.IntVar(&maxPost, "max-post", 0, "max post")
	flag.IntVar(&maxPaging, "max-paging", 0, "max paging")
	flag.Parse()
}

func main() {
	if maxPaging == 0 || maxPost == 0 {
		fmt.Println("you have to use the flag `--max-post=value --max-paging=value`")
		return
	}
	// fmt.Println(maxPost)
	// fmt.Println(maxPaging)

	cli := db.NewMongoConnection()
	repo := repositories.NewScrappingRepositoryImpl()
	svc := services.NewScrappingImpl(cli, repo)
	err := svc.Post(maxPost, maxPaging)
	if err != nil {
		panic(err)
	}
}
