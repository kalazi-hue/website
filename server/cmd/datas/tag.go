package datas

import (
	"gin-vue-admin/model"
	"github.com/brianvoe/gofakeit/v5"
	"gorm.io/gorm"
)

func FakeTag(db *gorm.DB) (err error) {
	var tmp bool
	return db.Transaction(func(tx *gorm.DB) error {
		for i := 0; i < 50; i++ {
			num := gofakeit.Number(0,1)

			if num == 1 {
				tmp = false
			} else {
				tmp = true
			}
			var tag = model.Tag{
				Name:gofakeit.Animal(),
				Sort: gofakeit.Number(1, 100),
				Status: &tmp,
			}
			if err := tx.Create(&tag).Error; err != nil { // 遇到错误时回滚事务
				return err
			}
		}
		return nil

	})
}