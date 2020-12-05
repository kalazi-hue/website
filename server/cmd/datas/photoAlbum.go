package datas

import (
	"gin-vue-admin/model"
	"github.com/brianvoe/gofakeit/v5"
	"gorm.io/gorm"
)

func FakePhotoAlbum(db *gorm.DB) (err error) {
	return db.Transaction(func(tx *gorm.DB) error {
		 t := true
		for i := 0; i < 30; i++ {
			var movie = model.PhotoAlbum{
				Title:gofakeit.BeerName(),
				Description: gofakeit.AppAuthor(),
				Cover: gofakeit.ImageURL(100, 30),
				Content: gofakeit.ImageURL(300,800) + "," + gofakeit.ImageURL(300,800) + "," +gofakeit.ImageURL(300,800) + "," +
					gofakeit.ImageURL(300,800) + "," +gofakeit.ImageURL(300,800) + "," +gofakeit.ImageURL(300,800),
				Keyword: gofakeit.Animal(),
				Reads:gofakeit.Number(1000,9999),
				Type: gofakeit.Number(1,5),
				Status: &t,
			}
			if err := tx.Create(&movie).Error; err != nil { // 遇到错误时回滚事务
				return err
			}
		}
		return nil

	})
}