
### collection操作関数一覧と使い方

### 1. **ドキュメントの挿入**
   - **関数**: `InsertOne`, `InsertMany`
   - **返り値の型**:
     - `InsertOne`: `(*mongo.InsertOneResult, error)`
     - `InsertMany`: `(*mongo.InsertManyResult, error)`
   - **書式指定子**: `InsertOneResult` と `InsertManyResult` の `InsertedID` フィールドにアクセスし、`ObjectID` は `%v` で表示可能。
   - **使い方**:
     ```go
     result, err := collection.InsertOne(ctx, bson.D{{"name", "John"}})
     fmt.Printf("Inserted document ID: %v\n", result.InsertedID)
     ```

### 2. **ドキュメントの取得（1件）**
   - **関数**: `FindOne`
   - **返り値の型**:
     - `(*mongo.SingleResult, error)` (実際には`Decode`メソッドで使用する)
   - **書式指定子**: `bson.M`（マップ）や`bson.D`（スライス）にデコードして `%v` で表示。
   - **使い方**:
     ```go
     var result bson.M
     err := collection.FindOne(ctx, bson.D{{"name", "John"}}).Decode(&result)
     fmt.Printf("Retrieved document: %v\n", result)
     ```

### 3. **ドキュメントの取得（複数）**
   - **関数**: `Find`
   - **返り値の型**:
     - `(*mongo.Cursor, error)`
   - **書式指定子**: `bson.M`（スライス）にデコードして `%v` で表示。
   - **使い方**:
     ```go
     cursor, err := collection.Find(ctx, bson.D{})
     var results []bson.M
     err = cursor.All(ctx, &results)
     fmt.Printf("Retrieved documents: %v\n", results)
     ```

### 4. **ドキュメントの更新（1件）**
   - **関数**: `UpdateOne`
   - **返り値の型**:
     - `(*mongo.UpdateResult, error)`
   - **書式指定子**: `UpdateResult` の `MatchedCount`, `ModifiedCount`, `UpsertedID` にアクセスし、それぞれ `%v` で表示。
   - **使い方**:
     ```go
     update := bson.D{{"$set", bson.D{{"name", "Johnny"}}}}
     result, err := collection.UpdateOne(ctx, bson.D{{"name", "John"}}, update)
     fmt.Printf("Matched %v documents and modified %v documents\n", result.MatchedCount, result.ModifiedCount)
     ```

### 5. **ドキュメントの更新（複数）**
   - **関数**: `UpdateMany`
   - **返り値の型**:
     - `(*mongo.UpdateResult, error)`
   - **書式指定子**: `UpdateResult` の `MatchedCount`, `ModifiedCount`, `UpsertedID` にアクセスし、それぞれ `%v` で表示。
   - **使い方**:
     ```go
     update := bson.D{{"$set", bson.D{{"active", true}}}}
     result, err := collection.UpdateMany(ctx, bson.D{{"status", "inactive"}}, update)
     fmt.Printf("Matched %v documents and modified %v documents\n", result.MatchedCount, result.ModifiedCount)
     ```

### 6. **ドキュメントの削除（1件）**
   - **関数**: `DeleteOne`
   - **返り値の型**:
     - `(*mongo.DeleteResult, error)`
   - **書式指定子**: `DeleteResult` の `DeletedCount` にアクセスして `%v` で表示。
   - **使い方**:
     ```go
     result, err := collection.DeleteOne(ctx, bson.D{{"name", "John"}})
     fmt.Printf("Deleted %v documents\n", result.DeletedCount)
     ```

### 7. **ドキュメントの削除（複数）**
   - **関数**: `DeleteMany`
   - **返り値の型**:
     - `(*mongo.DeleteResult, error)`
   - **書式指定子**: `DeleteResult` の `DeletedCount` にアクセスして `%v` で表示。
   - **使い方**:
     ```go
     result, err := collection.DeleteMany(ctx, bson.D{{"status", "inactive"}})
     fmt.Printf("Deleted %v documents\n", result.DeletedCount)
     ```

### 8. **ドキュメントの数を数える**
   - **関数**: `CountDocuments`
   - **返り値の型**:
     - `(int64, error)`
   - **書式指定子**: `int64` を `%d` で表示。
   - **使い方**:
     ```go
     count, err := collection.CountDocuments(ctx, bson.D{})
     fmt.Printf("Documents count: %d\n", count)
     ```

### 9. **ドキュメントの存在確認**
   - **関数**: `FindOne`
   - **返り値の型**:
     - `(*mongo.SingleResult, error)`（エラーが `mongo.ErrNoDocuments` かどうかをチェック）
   - **書式指定子**: 直接エラーメッセージを表示する場合は `%v`。
   - **使い方**:
     ```go
     err := collection.FindOne(ctx, bson.D{{"name", "John"}}).Err()
     if err == mongo.ErrNoDocuments {
         fmt.Println("Document not found")
     }
     ```

### 10. **インデックスの作成**
   - **関数**: `Indexes().CreateOne`, `Indexes().CreateMany`
   - **返り値の型**:
     - `(string, error)`（インデックス名が返される）
   - **書式指定子**: `string` を `%s` で表示。
   - **使い方**:
     ```go
     indexModel := mongo.IndexModel{
         Keys:    bson.D{{"email", 1}},
         Options: options.Index().SetUnique(true),
     }
     result, err := collection.Indexes().CreateOne(ctx, indexModel)
     fmt.Printf("Created index: %s\n", result)
     ```

### 11. **アグリゲーション（集計）**
   - **関数**: `Aggregate`
   - **返り値の型**:
     - `(*mongo.Cursor, error)`
   - **書式指定子**: `bson.M` にデコードして `%v` で表示。
   - **使い方**:
     ```go
     pipeline := mongo.Pipeline{
         bson.D{{"$match", bson.D{{"status", "active"}}}},
         bson.D{{"$group", bson.D{{"_id", "$department"}, {"total", bson.D{{"$sum", 1}}}}}},
     }
     cursor, err := collection.Aggregate(ctx, pipeline)
     var results []bson.M
     err = cursor.All(ctx, &results)
     fmt.Printf("Aggregation results: %v\n", results)
     ```

### 12. **コレクションのドロップ（削除）**
   - **関数**: `Drop`
   - **返り値の型**:
     - `(error)`（エラーがなければ`nil`が返される）
   - **書式指定子**: エラーメッセージを表示する場合は `%v`。
   - **使い方**:
     ```go
     err := collection.Drop(ctx)
     if err != nil {
         fmt.Printf("Failed to drop collection: %v\n", err)
     } else {
         fmt.Println("Collection dropped successfully")
     }
     ```

このリストを使って、どの操作がどの型を返すのか、そしてどの書式指定子を使えばいいのかを確認できるよ。これで操作が少しでも楽になると嬉しいな😊もし他に質問があれば、遠慮なく聞いてね！


# docker コマンド
 1.Dockerイメージのビルド
 docker build -t my-go-app .
 2.コンテナ実行
 docker run -d -p 8080:8080 my-go-app
 3.コンテナの確認
 docker ps
4.確認
curl http://localhost:8080
