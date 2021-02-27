package controllers

import (
	"context"
	c "document-service/config"
	"document-service/models"
	m "document-service/models"
	"errors"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAll(userID uint64) ([]m.Folder, []m.Documents) {
	client := c.MongoInit()
	ctx := context.Background()
	defer client.Disconnect(ctx)
	cursorFolder := client.Database(os.Getenv("MONGO_DNS")).Collection("folder")
	listFolder, err := cursorFolder.Find(ctx, bson.M{
		"$or": []interface{}{
			bson.M{"ispublic": true},
			bson.M{"ownerid": userID},
			bson.M{"share": bson.M{"$in": []uint64{userID}}},
		},
	})

	if err != nil {
		log.Fatal(err.Error())
	}
	defer listFolder.Close(ctx)

	var resultF []models.Folder
	for listFolder.Next(ctx) {
		var row models.Folder
		if err := listFolder.Decode(&row); err != nil {
			log.Fatal(err.Error())
		}
		resultF = append(resultF, row)
	}

	cursorDocument := client.Database(os.Getenv("MONGO_DNS")).Collection("documents")
	listDocument, err := cursorDocument.Find(ctx, bson.M{
		"$or": []interface{}{
			bson.M{"ownerid": userID},
			bson.M{"share": bson.M{"$in": []uint64{userID}}},
		},
	})

	if err != nil {
		log.Fatal(err.Error())
	}
	defer listDocument.Close(ctx)
	var resultD []models.Documents
	for listDocument.Next(ctx) {
		var row models.Documents
		if err := listDocument.Decode(&row); err != nil {
			log.Fatal(err.Error())
		}
		resultD = append(resultD, row)
	}
	return resultF, resultD
}

func GetFolder(ID string, ownerID uint64) (*m.Folder, error) {
	client := c.MongoInit()
	ctx := context.Background()
	defer client.Disconnect(ctx)

	cursor := client.Database(os.Getenv("MONGO_DNS")).Collection("folder")
	folder := cursor.FindOne(ctx, bson.M{
		"id":      ID,
		"ownerid": ownerID,
	})
	var result models.Folder
	err := folder.Decode(&result)
	if err != nil {
		return nil, errors.New("folder is not found")
	}
	return &result, nil
}

func AddFolder(req m.Folder) error {
	client := c.MongoInit()
	ctx := context.Background()
	defer client.Disconnect(ctx)
	cursor := client.Database(os.Getenv("MONGO_DNS")).Collection("folder")

	_, err := cursor.InsertOne(context.Background(), req)
	if err != nil {
		return err
	}

	return nil
}

func UpdateFolder(req m.Folder) error {
	client := c.MongoInit()
	ctx := context.Background()
	defer client.Disconnect(ctx)

	cursor := client.Database(os.Getenv("MONGO_DNS")).Collection("folder")
	_, err := cursor.UpdateOne(
		ctx,
		bson.M{"id": req.ID},
		bson.M{
			"$set": bson.M{
				"name":      req.Name,
				"timestamp": req.Timestamp,
			},
		},
	)

	if err != nil {
		return err
	}
	return nil
}

func DeleteDocumentFolder(ID string) error {
	client := c.MongoInit()
	ctx := context.Background()
	defer client.Disconnect(ctx)

	cursor := client.Database(os.Getenv("MONGO_DNS")).Collection("folder")
	_, err := cursor.DeleteOne(ctx, bson.M{"id": ID})

	if err != nil {
		return err
	}
	return nil
}

func GetDocumentByFolderID(ID string, ownerID uint64, share uint64) []m.Documents {
	client := c.MongoInit()
	ctx := context.Background()
	defer client.Disconnect(ctx)
	cursor := client.Database(os.Getenv("MONGO_DNS")).Collection("documents")

	listDocument, err := cursor.Find(ctx, bson.M{
		"folderid": ID,
		"type":     "document",
		"$and": []interface{}{
			bson.M{
				"$or": []interface{}{
					bson.M{"ownerid": ownerID},
					bson.M{"share": bson.M{"$in": []uint64{share}}},
				},
			},
		},
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer listDocument.Close(ctx)

	result := make([]models.Documents, 0)
	for listDocument.Next(ctx) {
		var row models.Documents
		if err := listDocument.Decode(&row); err != nil {
			log.Fatal(err.Error())
		}
		result = append(result, row)
	}
	return result
}

func SetNewFolder(req models.Folder) models.Folder {
	data := models.Folder{
		ID:        req.ID,
		Name:      req.Name,
		Type:      "folder",
		IsPublic:  true,
		OwnerID:   models.UserToken.UserID,
		Share:     []uint64{},
		Timestamp: req.Timestamp,
		CompanyID: models.UserToken.CompanyID,
	}
	return data
}
