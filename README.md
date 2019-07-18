# Client and Server

# create
```bash
npx create-react-app frontend --typescript
```

```bash
go generate
```

```
作成されていない初回のみ
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
go run github.com/vektah/dataloaden UserLoader string *github.com/KouT127/gin-sample/backend/domain/model.User
go run github.com/vektah/dataloaden UserSliceLoader string []*github.com/KouT127/gin-sample/backend/domain/model.User
```