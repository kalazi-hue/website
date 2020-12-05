package datas

import (
	"gin-vue-admin/model"
	"strconv"
	"github.com/brianvoe/gofakeit/v5"
	"gorm.io/gorm"
)




func FakeMovie(db *gorm.DB) (err error) {
	return db.Transaction(func(tx *gorm.DB) error {
		for i := 0; i < 30; i++ {
			var movie = model.Movie{
				Title:gofakeit.BeerName(),
				Description: gofakeit.AppAuthor(),
				Cover: gofakeit.ImageURL(100, 30),
				PlayTime: strconv.Itoa(gofakeit.Number(0, 2)) + ":" + strconv.Itoa(gofakeit.Minute()) + ":" + strconv.Itoa(gofakeit.Second()),
				Type:gofakeit.Number(1,100),
				Tags:gofakeit.Animal() + "," + gofakeit.Animal() + "," + gofakeit.BeerName(),
				Star:gofakeit.Number(1000,9999),
				PlayCount:gofakeit.Number(1000,9999),
				ShelfTime:gofakeit.Date(),

			}
			if err := tx.Create(&movie).Error; err != nil { // 遇到错误时回滚事务
				return err
			}
		}
		return nil

	})
}
