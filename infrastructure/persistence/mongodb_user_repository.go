package persistence

import (
	"context"

	"github.com/YuuHikida/GSC_backend_go/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

// user情報登録(1件)
func (r *UserRepository) Save(user *model.UserInfo) error {
	_, err := r.collection.InsertOne(context.TODO(), user)
	return err
}

// Initializeでコレクションをセットアップ
func NewMongoUserRepository() *UserRepository { // インターフェースを返す
	client, _ := mongo.Connect(context.TODO()) // 初期化の詳細は省略
	db := client.Database("gitInfoContributes")
	return &UserRepository{collection: db.Collection("user_info")} // ポインタ型を返す
}

// MongoDBから1件のドキュメントを取得して返す
func (r *UserRepository) FindOne(ctx context.Context, gitName string) (model.UserInfo, error) {
	var user model.UserInfo
	// bson.Mはドキュメントのキーと値をペアで保存するデータ型
	// キーはStinrg,値はinterface{}		var user model.UserInfo
	err := r.collection.FindOne(ctx, bson.M{"git_name": gitName}).Decode(&user)
	return user, err
}

// MongoDBから全件のドキュメントを取得して返す
func (r *UserRepository) FindAll(ctx context.Context) ([]model.UserInfo, error) {
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var users []model.UserInfo
	err = cursor.All(ctx, &users)
	return users, err
}
