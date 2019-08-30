# Client and Server

# create
```bash
npx create-react-app frontend --typescript
```

```bash
go generate
```
```bash
create database go_graph
```


```
作成されていない初回のみ
cd backend/interface/graph

gqlgen init

gqlgen -v

更新後は、implement methodで追加する。 
```

```
# GraphQLを試す場合
https://github.com/prisma/graphql-playground
```
```
Name InputType ReturnType
1-1
go run github.com/vektah/dataloaden UserLoader int *github.com/KouT127/gin-sample/backend/domain/model.User
go run github.com/vektah/dataloaden TaskCountLoader int *github.com/KouT127/gin-sample/backend/domain/model.Task
1-m m-m
go run github.com/vektah/dataloaden TaskSliceLoader int []*github.com/KouT127/gin-sample/backend/domain/model.Task
```

# migration
```zsh
現在のDBのテーブル情報を吐き出す
make show-migrations

実行されるクエリを確認する
make mysqldef-dry

差分を実行する
make mysqldef
```

```zsh
# 実際にマイグレーションを適用する
cd backend
migrate -source file://infrastracture/database/migration/  -database 'mysql://root:@tcp(localhost:3306)/go_graph' up 1
```
### DBテーブル変更手順
1. make show-migrationsでschema.sqlを吐き出す。
2. 吐き出したschema.sqlを編集する。
3. make mysqldef-dryで吐き出されるクエリを確認する。
4. 吐き出したクエリをMigrationのSqlとして、別のsqlファイルで保存する。
5. make migrate で実際にmigrateを行う。