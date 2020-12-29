package service

import (
	"fmt"
	"gin-vue-admin/config"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)
type PhotoAlbumUploader struct {
	Title string `json:title`
	Desc string `json:desc`
	Keywords []string `json:keywords`
	Status bool `json:status`
	Type string `json type`
	Md5 string `json md5`
	FileName string `json filename`
	IsCover bool `json isCover`
	DstName string `json dstName`
}



//@author: [piexlmax](https://github.com/piexlmax)
//@function: SaveChunk
//@description: 保存文件切片路径
//@param: uploader model.ExaSimpleUploader
//@return: err error

func SaveChunk(uploader model.ExaSimpleUploader) (err error) {
	return global.GVA_DB.Create(uploader).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CheckFileMd5
//@description: 检查文件是否已经上传过
//@param: md5 string
//@return: err error, uploads []model.ExaSimpleUploader, isDone bool

//func PhotoAlbumCheckMd5(md5 string) (err error, uploads []model.ExaSimpleUploader, isDone bool) {
//	err = global.GVA_DB.Find(&uploads, "identifier = ? AND is_done = ?", md5, false).Error
//	isDone = errors.Is(global.GVA_DB.First(&model.ExaSimpleUploader{}, "identifier = ? AND is_done = ?", md5, true).Error, gorm.ErrRecordNotFound)
//	return err, uploads, !isDone
//}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: PhotoMerge
//@description: 合并文件
//@param: md5 string, fileName string
//@return: err error

func PhotoMerge(md5 string, fileName string) (dstName string, err error) {
	var uploadConfig  config.Upload
	var now = global.MyTime(time.Now())
	finishDir := uploadConfig.ImgUploadPath + now.String() + "/"
	chunkDir := uploadConfig.ImgChunkPath + md5
	dstName = md5 + path.Ext(fileName)

	//打开切片文件夹
	rd, err := ioutil.ReadDir(chunkDir)
	_ = os.MkdirAll(finishDir, os.ModePerm)
	//创建目标文件
	fd, _ := os.OpenFile(finishDir+dstName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	//将切片文件按照顺序写入
	for k := range rd {
		content, _ := ioutil.ReadFile(chunkDir + "/" + fileName + strconv.Itoa(k+1))
		fd.Write(content)
	}
	//关闭文件
	defer fd.Close()

	//if err != nil {
	//	return dstName,err
	//}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		//删除切片信息
		if err = tx.Delete(&model.ExaSimpleUploader{}, "identifier = ? AND is_done = ?", md5, false).Error; err != nil {
			fmt.Println(err)
			return err
		}
		data := model.ExaSimpleUploader{
			Identifier: md5,
			IsDone:     true,
			FilePath:   finishDir + fileName,
			Filename:   fileName,
		}
		// 添加文件信息
		if err = tx.Create(&data).Error; err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	})

	err = os.RemoveAll(chunkDir) //清除切片
	return dstName,err
}

func PhotoSaveToDB(up *PhotoAlbumUploader) (err error) {
	var photoAlbum model.PhotoAlbum
	var data model.PhotoAlbum
	var IsCreate = true
	var now = global.MyTime(time.Now())
	res := global.GVA_DB.Where("title", up.Title).Where( "description", up.Desc).First(&photoAlbum)
	if res.RowsAffected > 0{
		global.GVA_LOG.Info(">0")
		global.GVA_LOG.Info(photoAlbum.Title)
		IsCreate = false
	}

	if IsCreate {
		if up.IsCover {
			data.Cover = up.DstName
		} else {
			data.Content = up.DstName
		}
		data.Title = up.Title
		data.Description =up.Desc
		data.Type = 1
		data.Status = &up.Status
		data.Reads = 123
		data.Is_top = &up.Status
		data.Keyword = strings.Join(up.Keywords, ",")
		data.CreatedAt = now
		data.UpdatedAt = now
		err = global.GVA_DB.Create(&data).Error
		return err
	}

	//更新
	if photoAlbum.Content == "" {
		err = global.GVA_DB.Model(&photoAlbum).Update("content", up.DstName).Error
	} else {
		err = global.GVA_DB.Model(&photoAlbum).Update("content", photoAlbum.Content + ",\n" + up.DstName).Error
	}


	return err
}





