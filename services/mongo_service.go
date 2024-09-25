package services

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

// Initializeでコレクションをセットアップ
func Initialize(client *mongo.Client) {
	db := client.Database("gitInfoContributes")
	collection = db.Collection("user_info")
}

// FindOneDocument は、MongoDBから1件のドキュメントを取得して返す
func FindOneDocument(ctx context.Context, gitName string) (bson.M, error) {
	// bson.Mはドキュメントのキーと値をペアで保存するデータ型
	// キーはStinrg,値はinterface{}
	var result bson.M
	err := collection.FindOne(ctx, bson.M{"git_name": gitName}).Decode(&result) //.Decode(格納)
	if err != nil {
		log.Printf("Error finding document: %v", err)
		return nil, err
	}
	return result, nil
}

// FindAllDocuments は、MongoDBから全件のドキュメントを取得して返す
func FindAllDocuments(ctx context.Context) ([]bson.M, error) {
	var results []bson.M
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("Error finding documents: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &results)
	if err != nil {
		log.Printf("Error decoding documents: %v", err)
		return nil, err
	}
	return results, nil
}
