package service

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/utils"
	"gorm.io/gorm"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)
type MovieUploader struct {
	Desc string `json:desc`
	FileName string `json filename`
	IsCover bool `json isCover`
	IsRecommend bool `json isRecommend`
	IsTop bool `json isTop`
	Md5 string `json md5`
	Status bool `json:status`
	Tag []string `json:tag`
	Title string `json:title`
	Type string `json type`
	DstName string `json dstName`
}



//@author: [piexlmax](https://github.com/piexlmax)
//@function: SaveChunk
//@description: 保存文件切片路径
//@param: uploader model.ExaSimpleUploader
//@return: err error

func MovieSaveChunk(uploader model.ExaSimpleUploader) (err error) {
	return global.GVA_DB.Create(uploader).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CheckFileMd5
//@description: 检查文件是否已经上传过
//@param: md5 string
//@return: err error, uploads []model.ExaSimpleUploader, isDone bool

func MovieCheckFileMd5(md5 string) (err error, uploads []model.ExaSimpleUploader, isDone bool) {
	err = global.GVA_DB.Find(&uploads, "identifier = ? AND is_done = ?", md5, false).Error
	isDone = errors.Is(global.GVA_DB.First(&model.ExaSimpleUploader{}, "identifier = ? AND is_done = ?", md5, true).Error, gorm.ErrRecordNotFound)
	return err, uploads, !isDone
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: PhotoMerge
//@description: 合并文件
//@param: md5 string, fileName string
//@return: err error

func MovieMerge(md5 string, fileName string, title string) (dstName string, err error) {
	var finishDir,finishName string

	uploadConfig := global.GVA_CONFIG.Upload
	var today = time.Now().Format("2006-01-02")
	global.GVA_LOG.Error(path.Ext(fileName))
	if path.Ext(fileName) == ".mp4" {
		finishDir = uploadConfig.VideoUploadPath + today + "/" + utils.MD5V([]byte(title)) +"/origin/"
		finishName = finishDir + "org" +path.Ext(fileName)
	} else {
		finishDir = uploadConfig.VideoUploadPath + today + "/" + utils.MD5V([]byte(title))
		finishName = finishDir + "/cover.jpg"
	}
	//判断文件类型

	chunkDir := uploadConfig.VideoChunkPath + md5
	dstNameWithDomain := uploadConfig.VideoDomain + finishName

	//打开切片文件夹
	rd, err := ioutil.ReadDir(chunkDir)
	err = os.MkdirAll(finishDir, os.ModePerm)
	if err != nil {
		global.GVA_LOG.Info("mkdir fail:" + finishDir)
	}
	//创建目标文件
	fd, _ := os.OpenFile(finishName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
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
	return dstNameWithDomain,err
}

func MovieSaveToDB(up *MovieUploader) (err error) {
	var movie model.Movie
	var data model.Movie
	//var IsCreate = true
	var now = global.MyTime(time.Now())
	res := global.GVA_DB.Where("title", up.Title).Where( "type", up.Type).First(&movie)
	if res.RowsAffected > 0{
		//global.GVA_LOG.Info(">0")
		//global.GVA_LOG.Info(movie.Title)
		//IsCreate = false
	}

	if up.IsCover {
		data.Cover = up.DstName
		data.Title = up.Title
		data.Description =up.Desc
		data.Type = up.Type
		data.Status = &up.Status
		data.PlayCount = 123
		data.IsRecommend = &up.IsRecommend
		data.IsTop = &up.IsTop
		data.Tags = strings.Join(up.Tag, ",")
		data.CreatedAt = now
		data.UpdatedAt = now
		err = global.GVA_DB.Create(&data).Error
		return err
	}


	err = global.GVA_DB.Model(&movie).Updates(map[string]interface{}{"play_url": getPlayUrl(up.DstName), "down_url": up.DstName}).Error
	//info := strconv.Itoa(int(movie.ID)) + "," + getDownPath(up.DstName)
	info := strconv.Itoa(int(movie.ID)) + "," + up.DstName

	global.GVA_LOG.Info(info)

	global.GVA_REDIS.LPush("video-process", info).Result()
	defer global.GVA_REDIS.Close()
	//utils.SendMsgToGroup(info)
	return err

}

func getPlayUrl(downUrl string) string {
	s := strings.Split(downUrl, "origin")
	//res, _ := url.Parse(downUrl)
	return s[0] + "m3u8/index.m3u8"
}

func getDownPath(downUrl string) string {
	res,_ := url.Parse(downUrl)
	return res.Path
}








