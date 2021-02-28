package controllers

import (
	"document-service/helper"
	"document-service/models"
	"document-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	httpStatus := http.StatusOK

	folderList, documentList := services.GetAll(models.UserToken.UserID)
	var result []interface{}
	for _, resFolder := range folderList {
		result = append(result, resFolder)
	}
	for _, resDocument := range documentList {
		result = append(result, resDocument)
	}
	res := helper.Response(false, "", result)
	c.JSON(httpStatus, res)
}

func SetFolder(c *gin.Context) {
	var req models.Folder
	if err := c.ShouldBindJSON(&req); err != nil {
		res := helper.Response(true, err.Error(), map[string]interface{}{})
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	dataReq := services.SetNewFolder(req)
	if res, _ := services.GetFolder(dataReq.ID, models.UserToken.UserID); res == nil {
		if err := services.AddFolder(dataReq); err != nil {
			res := helper.Response(true, err.Error(), map[string]interface{}{})
			c.JSON(http.StatusUnprocessableEntity, res)
			return
		}
	} else {
		if err := services.UpdateFolder(dataReq); err != nil {
			res := helper.Response(true, err.Error(), map[string]interface{}{})
			c.JSON(http.StatusUnprocessableEntity, res)
			return
		}
	}

	res := helper.Response(false, "folder created", dataReq)
	c.JSON(http.StatusOK, res)
}

func DeleteFolder(c *gin.Context) {
	var req models.Folder
	if err := c.ShouldBindJSON(&req); err != nil {
		res := helper.Response(true, err.Error(), map[string]interface{}{})
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}
	_, err := services.GetFolder(req.ID, models.UserToken.UserID)
	if err != nil {
		res := helper.Response(true, err.Error(), map[string]interface{}{})
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	if err := services.DeleteDocumentFolder(req.ID); err != nil {
		res := helper.Response(true, err.Error(), map[string]interface{}{})
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	res := helper.Response(false, "Success delete folder", map[string]interface{}{})

	c.JSON(http.StatusOK, res)
}

func DocumentByFolderID(c *gin.Context) {
	folderID := c.Param("folder_id")
	documentList := services.GetDocumentByFolderID(folderID, models.UserToken.UserID, models.UserToken.UserID)
	res := helper.Response(false, "Success get document", documentList)
	c.JSON(http.StatusOK, res)
}
