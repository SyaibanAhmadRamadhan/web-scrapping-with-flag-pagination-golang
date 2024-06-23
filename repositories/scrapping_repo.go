package repositories

import (
	"context"

	"SyaibanAhmadRamadhan/web-scrapping-with-flag-pagination-golang/entities"

	"github.com/qiniu/qmgo"
)

type ScrappingRepository interface {
	Creates(ctx context.Context, db *qmgo.Collection, entity []entities.Scrapping) ([]entities.Scrapping, error)
}

type ScrappingRepositoryImpl struct{}

func NewScrappingRepositoryImpl() ScrappingRepository {
	return &ScrappingRepositoryImpl{}
}
