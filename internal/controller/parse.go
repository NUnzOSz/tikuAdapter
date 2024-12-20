package controller

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
	"tikuAdapter/internal/service"
	"tikuAdapter/pkg/global"

	"code.sajari.com/docconv/v2"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// Parser 解析文件接口 支持解析docx和xlsx
func Parser(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, global.ErrorParam)
		return
	}
	uploadedFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, global.ErrorParam)
		return
	}
	defer func(uploadedFile multipart.File) {
		err := uploadedFile.Close()
		if err != nil {
			c.JSON(http.StatusBadRequest, global.ErrorParam)
			return
		}
	}(uploadedFile)
	contentType := file.Header.Get("Content-Type")
	if contentType == "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
		f, err := excelize.OpenReader(uploadedFile)
		if err != nil {
			c.JSON(http.StatusBadRequest, global.ErrorParseFile)
			return
		}
		var options = make([]string, 0)
		err = json.Unmarshal([]byte(c.PostForm("options")), &options)
		if err != nil {
			c.JSON(http.StatusBadRequest, global.ErrorParam)
			return
		}
		opt := service.XLSXOptions{
			SheetName: c.PostForm("sheetName"),
			Question:  c.PostForm("question"),
			Answer:    c.PostForm("answer"),
			Option:    options,
		}

		c.JSON(http.StatusOK, service.ParseXls(f, opt))
	} else {
		convert, err := docconv.Convert(uploadedFile, contentType, true)
		if err != nil {
			c.JSON(http.StatusBadRequest, global.ErrorParseFile)
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, global.ErrorParseFile)
			return
		}
		c.JSON(http.StatusOK, convert)
	}
}
