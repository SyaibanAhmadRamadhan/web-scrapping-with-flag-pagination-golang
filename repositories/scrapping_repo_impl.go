package repositories

import (
	"context"

	"SyaibanAhmadRamadhan/web-scrapping-with-flag-pagination-golang/entities"

	"github.com/qiniu/qmgo"
)

func (repo *ScrappingRepositoryImpl) Creates(ctx context.Context, db *qmgo.Collection, entity []entities.Scrapping) ([]entities.Scrapping, error) {
	// panic("stop")
	_, err := db.InsertMany(ctx, entity)
	if err != nil {
		panic(err)
	}

	return entity, nil
}
