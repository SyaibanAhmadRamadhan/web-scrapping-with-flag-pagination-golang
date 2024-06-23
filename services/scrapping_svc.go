package services

import (
	"SyaibanAhmadRamadhan/web-scrapping-with-flag-pagination-golang/repositories"

	"github.com/qiniu/qmgo"
)

type ScrappingService interface {
	Post(maxPost int, maxPaging int) error
}

type ScrappingServiceImpl struct {
	DB   *qmgo.Collection
	Repo repositories.ScrappingRepository
}

func NewScrappingImpl(
	db *qmgo.Collection,
	repo repositories.ScrappingRepository,
) ScrappingService {
	return &ScrappingServiceImpl{
		DB:   db,
		Repo: repo,
	}
}
