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
```bash
cd backend
migrate -source file://infrastracture/database/migration/  -database 'mysql://root:@tcp(localhost:3306)/go_graph' up 1
```