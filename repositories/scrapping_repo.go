package repositories

import (
	"context"

	"technical-test-pt-semesta-arus-technology/entities"

	"github.com/qiniu/qmgo"
)

type ScrappingRepository interface {
	Creates(ctx context.Context, db *qmgo.Collection, entity []entities.Scrapping) ([]entities.Scrapping, error)
}

type ScrappingRepositoryImpl struct{}

func NewScrappingRepositoryImpl() ScrappingRepository {
	return &ScrappingRepositoryImpl{}
}
