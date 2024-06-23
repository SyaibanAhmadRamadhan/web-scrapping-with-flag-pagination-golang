package services

import (
	"context"
	"fmt"
	"log"
	"os"

	"SyaibanAhmadRamadhan/web-scrapping-with-flag-pagination-golang/entities"
	"SyaibanAhmadRamadhan/web-scrapping-with-flag-pagination-golang/utils"

	"github.com/gocolly/colly/v2"
)

func (svc *ScrappingServiceImpl) Post(maxPost int, maxPaging int) error {
	var scrappingEntities []entities.Scrapping
	var scrappingEntity entities.Scrapping
	counter := 0
	pageScrapping := "https://nasional.sindonews.com/more/5"

	collector := colly.NewCollector(
		colly.AllowedDomains("nasional.sindonews.com", "www.nasional.sindonews.com"),
	)
	detailCollector := collector.Clone()

	collector.OnHTML(".sm-pl15", func(h *colly.HTMLElement) {
		if h.Attr("class") == "width-100 mb24 sm-pl15 sm-pr15" {
			scrappingEntity.TanggalArtikel = h.ChildText(".date-kanal")
			scrappingEntity.Judul = h.ChildText(".desc-kanal")
			detailCollector.Visit(h.ChildAttr("a", "href"))
		}
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	detailCollector.OnHTML(".detail-nama-redaksi", func(h *colly.HTMLElement) {
		scrappingEntity.NamaPenulis = h.Text
	})

	detailCollector.OnHTML(".detail-desc", func(h *colly.HTMLElement) {
		scrappingEntity.IsiArtikel = h.Text
	})

	detailCollector.OnHTML(".detail-img", func(h *colly.HTMLElement) {
		scrappingEntity.Gambar = h.ChildAttr("img", "data-src")
		file := utils.BuildFileName(h.ChildAttr("img", "data-src"))
		utils.PutFile(file, h.ChildAttr("img", "data-src"))
		scrappingEntities = append(scrappingEntities, scrappingEntity)
		if len(scrappingEntities) >= maxPost {
			_, err := svc.Repo.Creates(context.Background(), svc.DB, scrappingEntities)
			if err != nil {
				panic(err)
			}
			log.Println("scrapping website successfully")
			log.Printf("youre max paging : %d | max post : %d", maxPaging, maxPost)
			os.Exit(0)
		}
	})

	var url []string
	collector.OnHTML(".btn-pagination", func(h *colly.HTMLElement) {
		nextPage := h.Request.AbsoluteURL(h.ChildAttr("a", "href"))
		collector.Visit(nextPage)
		url = append(url, nextPage)
	})

	collector.OnResponse(func(r *colly.Response) {
		if counter > maxPaging {
			_, err := svc.Repo.Creates(context.Background(), svc.DB, scrappingEntities)
			if err != nil {
				panic(err)
			}
			log.Println("scrapping website successfully")
			log.Printf("youre max paging : %d | max post : %d", maxPaging, maxPost)
			os.Exit(0)
		}
		counter++
	})
	collector.Visit(pageScrapping)

	_, err := svc.Repo.Creates(context.Background(), svc.DB, scrappingEntities)
	if err != nil {
		panic(err)
	}
	return nil
}
