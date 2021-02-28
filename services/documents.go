package services

import (
	"context"
	c "document-service/config"
	m "document-service/models"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/bson"
)

func GetDocument(documentID string, ownerID uint64, share uint64) (*m.Documents, error) {
	client := c.MongoInit()
	ctx := context.Background()
	defer client.Disconnect(ctx)
	cursor := client.Database(os.Getenv("MONGO_DNS")).Collection("documents")
	document := cursor.FindOne(ctx, bson.M{
		"id": documentID,
		"$and": []interface{}{
			bson.M{
				"$or": []interface{}{
					bson.M{"ownerid": ownerID},
					bson.M{"share": bson.M{"$in": []uint64{share}}},
				},
			},
		},
	})

	var res m.Documents
	err := document.Decode(&res)
	if err != nil {
		return nil, errors.New("document is not found")
	}

	return &res, nil
}

func AddDocument(req m.Documents) error {
	client := c.MongoInit()

	ctx := context.Background()
	defer client.Disconnect(ctx)
	cursor := client.Database(os.Getenv("MONGO_DNS")).Collection("documents")
	_, err := cursor.InsertOne(context.Background(), req)
	if err != nil {
		return err
	}

	return nil
}

func UpdateDocument(documentSetRequest m.Documents) error {
	client := c.MongoInit()
	ctx := context.Background()
	defer client.Disconnect(ctx)
	cursor := client.Database(os.Getenv("MONGO_DNS")).Collection("documents")

	_, err := cursor.UpdateOne(
		ctx,
		bson.M{"id": documentSetRequest.ID, "type": "document"},
		bson.M{
			"$set": bson.M{
				"name":      documentSetRequest.Name,
				"folderid":  documentSetRequest.FolderID,
				"content":   documentSetRequest.Content,
				"timestamp": documentSetRequest.Timestamp,
				"ownerid":   documentSetRequest.OwnerID,
				"share":     documentSetRequest.Share,
				"companyid": documentSetRequest.CompanyID,
			},
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func SetNewDocument(req m.Documents) m.Documents {
	var newDocument m.Documents
	newDocument.ID = req.ID
	newDocument.Name = req.Name
	newDocument.Type = "document"
	newDocument.FolderID = req.FolderID
	newDocument.Content = req.Content
	if req.OwnerID == 0 {
		newDocument.OwnerID = m.UserToken.UserID
	} else {
		newDocument.OwnerID = req.OwnerID
	}
	newDocument.Share = req.Share
	newDocument.Timestamp = req.Timestamp
	newDocument.CompanyID = req.CompanyID

	return newDocument
}
