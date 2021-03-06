package api

import (
	"bytes"
	"encoding/json"
	"gin-vue-admin/admin/service"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/response"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
)

// @Tags SimpleUploader
// @Summary 断点续传插件版示例
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "断点续传插件版示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"切片创建成功"}"
// @Router /simpleUploader/upload [post]
func PhotoAlbumUploaderUpload(c *gin.Context) {
	var chunk model.ExaSimpleUploader
	uploadConfig := global.GVA_CONFIG.Upload
	_, header, err := c.Request.FormFile("file")
	chunk.Filename = c.PostForm("filename")
	chunk.ChunkNumber = c.PostForm("chunkNumber")
	chunk.CurrentChunkSize = c.PostForm("currentChunkSize")
	chunk.Identifier = c.PostForm("identifier")
	chunk.TotalSize = c.PostForm("totalSize")
	chunk.TotalChunks = c.PostForm("totalChunks")

	var chunkDir = uploadConfig.ImgChunkPath + chunk.Identifier + "/"
	hasDir, _ := utils.PathExists(chunkDir)
	if !hasDir {
		if err := utils.CreateDir(chunkDir); err != nil {
			global.GVA_LOG.Error("创建目录失败!", zap.Any("err", err))
		}
	}
	chunkPath := chunkDir + chunk.Filename + chunk.ChunkNumber
	err = c.SaveUploadedFile(header, chunkPath)
	if err != nil {
		global.GVA_LOG.Error("切片创建失败!", zap.Any("err", err))
		response.FailWithMessage("切片创建失败", c)
		return
	}
	chunk.CurrentChunkPath = chunkPath
	err = service.SaveChunk(chunk)
	if err != nil {
		global.GVA_LOG.Error("切片创建失败!", zap.Any("err", err))
		response.FailWithMessage("切片创建失败", c)
		return
	} else {
		response.OkWithMessage("切片创建成功", c)
	}
}

// @Tags SimpleUploader
// @Summary 断点续传插件版示例
// @Security ApiKeyAuth
// @Produce  application/json
// @Param md5 query string true "md5"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /simpleUploader/checkFileMd5 [get]
func PhotoAlbumCheckMd5(c *gin.Context) {
	//md5 := c.Query("md5")
	response.OkWithDetailed(gin.H{
		"chunks": 0,
		"isDone": false,
	},"查询成功", c)
	//err, chunks, isDone := service.CheckFileMd5(md5)
	//if err != nil {
	//	global.GVA_LOG.Error("md5读取失败!", zap.Any("err", err))
	//	response.FailWithMessage("md5读取失败", c)
	//} else {
	//	response.OkWithDetailed(gin.H{
	//		"chunks": chunks,
	//		"isDone": isDone,
	//	},"查询成功", c)
	//}
}

func PhotoAlbumUploadMerge(c *gin.Context) {
	buf := make([]byte, 1024)
	data := new(service.PhotoAlbumUploader)
	n, _ := c.Request.Body.Read(buf)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf[:n]))
	j := buf[0:n]
	err := json.Unmarshal(j, &data)
	if err != nil {
		global.GVA_LOG.Error("error", zap.Any("err", err))
	}


	//.上传内容处理
	data.DstName, _ = service.PhotoMerge(data.Md5, data.FileName, data.Title)//4e9002c48226e9c1d99ca40a02ea925b.jpg

	//入库
	err = service.PhotoSaveToDB(data)
	if err != nil {
		global.GVA_LOG.Error("写真动漫入库失败!", zap.Any("err", err))
		response.FailWithMessage("写真动漫入库失败", c)
		return
	}
	response.OkWithMessage("写真动漫合并成功", c)
}