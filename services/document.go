package services

import (
	"document-service/controllers"
	c "document-service/controllers"
	"document-service/models"
	"document-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUpdateDocument(ctx *gin.Context) {
	var doc models.Documents
	if err := ctx.ShouldBindJSON(&doc); err != nil {
		res := utils.Response(true, err.Error(), map[string]interface{}{})
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	data := c.SetNewDocument(doc)
	if res, _ := c.GetFolder(data.FolderID, models.UserToken.UserID); res == nil {
		res := utils.Response(true, "Folder ID not found", map[string]interface{}{})
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	if res, _ := c.GetDocument(data.ID, models.UserToken.UserID, models.UserToken.UserID); res == nil {
		if err := c.AddDocument(data); err != nil {
			res := utils.Response(true, err.Error(), map[string]interface{}{})
			ctx.JSON(http.StatusUnprocessableEntity, res)
			return
		}
	} else {
		if err := c.UpdateDocument(data); err != nil {
			res := utils.Response(true, err.Error(), map[string]interface{}{})
			ctx.JSON(http.StatusUnprocessableEntity, res)
			return
		}
	}

	//config.DeleteCache("DOCUMENT:" + data.ID)
	res := utils.Response(false, "Success set document", data)
	ctx.JSON(http.StatusOK, res)
}

func DetailDocument(c *gin.Context) {
	docID := c.Param("document_id")
	documentDetail, err := controllers.GetDocument(docID, models.UserToken.UserID, models.UserToken.UserID)
	if err != nil {
		res := utils.Response(true, err.Error(), map[string]interface{}{})
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}
	res := utils.Response(false, "Success get document", map[string]interface{}{"document": documentDetail})
	c.JSON(http.StatusOK, res)
}

func DeleteDocument(c *gin.Context) {
	var req models.Documents
	if err := c.ShouldBindJSON(&req); err != nil {
		res := utils.Response(true, err.Error(), map[string]interface{}{})
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}
	documentData, err := controllers.GetDocument(req.ID, models.UserToken.UserID, models.UserToken.UserID)
	if err != nil {
		res := utils.Response(true, err.Error(), map[string]interface{}{})
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}
	if documentData.OwnerID != models.UserToken.UserID {
		res := utils.Response(true, "You have no permission", map[string]interface{}{})
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	if err := controllers.DeleteDocumentFolder(documentData.ID); err != nil {
		res := utils.Response(true, err.Error(), map[string]interface{}{})
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}
	//config.DeleteCache("DOCUMENT:" + req.ID)
	res := utils.Response(false, "Success delete document", map[string]interface{}{})
	c.JSON(http.StatusOK, res)
}
