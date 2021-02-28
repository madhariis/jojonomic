package services

import (
	"document-service/controllers"
	"document-service/models"
	"document-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	httpStatus := http.StatusOK

	folderList, documentList := controllers.GetAll(models.UserToken.UserID)
	var result []interface{}
	for _, resFolder := range folderList {
		result = append(result, resFolder)
	}
	for _, resDocument := range documentList {
		result = append(result, resDocument)
	}
	res := utils.Response(false, "", result)
	c.JSON(httpStatus, res)
}

func SetFolder(c *gin.Context) {
	var req models.Folder
	if err := c.ShouldBindJSON(&req); err != nil {
		res := utils.Response(true, err.Error(), map[string]interface{}{})
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	dataReq := controllers.SetNewFolder(req)
	if res, _ := controllers.GetFolder(dataReq.ID, models.UserToken.UserID); res == nil {
		if err := controllers.AddFolder(dataReq); err != nil {
			res := utils.Response(true, err.Error(), map[string]interface{}{})
			c.JSON(http.StatusUnprocessableEntity, res)
			return
		}
	} else {
		if err := controllers.UpdateFolder(dataReq); err != nil {
			res := utils.Response(true, err.Error(), map[string]interface{}{})
			c.JSON(http.StatusUnprocessableEntity, res)
			return
		}
	}
	//config.DeleteCache("FOLDER:" + dataReq.ID)
	res := utils.Response(false, "folder created", dataReq)
	c.JSON(http.StatusOK, res)
}

func DeleteFolder(c *gin.Context) {
	var req models.Folder
	if err := c.ShouldBindJSON(&req); err != nil {
		res := utils.Response(true, err.Error(), map[string]interface{}{})
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}
	_, err := controllers.GetFolder(req.ID, models.UserToken.UserID)
	if err != nil {
		res := utils.Response(true, err.Error(), map[string]interface{}{})
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	if err := controllers.DeleteDocumentFolder(req.ID); err != nil {
		res := utils.Response(true, err.Error(), map[string]interface{}{})
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	res := utils.Response(false, "Success delete folder", map[string]interface{}{})
	//config.DeleteCache("FOLDER:" + req.ID)
	c.JSON(http.StatusOK, res)
}

func DocumentByFolderID(c *gin.Context) {
	folderID := c.Param("folder_id")
	documentList := controllers.GetDocumentByFolderID(folderID, models.UserToken.UserID, models.UserToken.UserID)
	res := utils.Response(false, "Success get document", documentList)
	c.JSON(http.StatusOK, res)
}
