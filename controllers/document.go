package controllers

import (
	"document-service/helper"
	"document-service/models"
	"document-service/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUpdateDocument(ctx *gin.Context) {
	var doc models.Documents
	if err := ctx.ShouldBindJSON(&doc); err != nil {
		res := helper.Response(true, err.Error(), map[string]interface{}{})
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	data := repositories.SetNewDocument(doc)
	if res, _ := repositories.GetFolder(data.FolderID, models.UserToken.UserID); res == nil {
		res := helper.Response(true, "Folder ID not found", map[string]interface{}{})
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	if res, _ := repositories.GetDocument(data.ID, models.UserToken.UserID, models.UserToken.UserID); res == nil {
		if err := repositories.AddDocument(data); err != nil {
			res := helper.Response(true, err.Error(), map[string]interface{}{})
			ctx.JSON(http.StatusUnprocessableEntity, res)
			return
		}
	} else {
		if err := repositories.UpdateDocument(data); err != nil {
			res := helper.Response(true, err.Error(), map[string]interface{}{})
			ctx.JSON(http.StatusUnprocessableEntity, res)
			return
		}
	}


	res := helper.Response(false, "Success set document", data)
	ctx.JSON(http.StatusOK, res)
}

func DetailDocument(c *gin.Context) {
	docID := c.Param("document_id")
	documentDetail, err := repositories.GetDocument(docID, models.UserToken.UserID, models.UserToken.UserID)
	if err != nil {
		res := helper.Response(true, err.Error(), map[string]interface{}{})
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}
	res := helper.Response(false, "Success get document", map[string]interface{}{"document": documentDetail})
	c.JSON(http.StatusOK, res)
}

func DeleteDocument(c *gin.Context) {
	var req models.Documents
	if err := c.ShouldBindJSON(&req); err != nil {
		res := helper.Response(true, err.Error(), map[string]interface{}{})
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}
	documentData, err := repositories.GetDocument(req.ID, models.UserToken.UserID, models.UserToken.UserID)
	if err != nil {
		res := helper.Response(true, err.Error(), map[string]interface{}{})
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}
	if documentData.OwnerID != models.UserToken.UserID {
		res := helper.Response(true, "You have no permission", map[string]interface{}{})
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	if err := repositories.DeleteDocumentFolder(documentData.ID); err != nil {
		res := helper.Response(true, err.Error(), map[string]interface{}{})
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	res := helper.Response(false, "Success delete document", map[string]interface{}{})
	c.JSON(http.StatusOK, res)
}
