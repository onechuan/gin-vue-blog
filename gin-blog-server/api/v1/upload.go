package v1

import (
	"gin-blog/service"
	"gin-blog/utils"
	"gin-blog/utils/r"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Upload struct {
	serv service.Upload
}

func (api *Upload) UploadFile(c *gin.Context) {
	_, fileHeader, err := c.Request.FormFile("file")
	// 处理文件接收错误
	if err != nil {
		code := r.EEROR_FILE_RECEIVE
		utils.Logger.Error(r.GetMsg(code), zap.Error(err))
		r.SendCode(c, code)
	}
	// 上传文件获取文件路径
	url, code := api.serv.UploadFile(fileHeader)
	r.SendData(c, code, url)
}