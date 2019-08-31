package dataloader

import (
	"context"
	"fmt"
	"github.com/KouT127/gin-sample/backend/domain/model"
	"github.com/KouT127/gin-sample/backend/infrastracture/database"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
	"time"
)

type ctxKeyType struct{ name string }

var ctxKey = ctxKeyType{"appCtx"}

type Loaders struct {
	UserById        *UserLoader
	TaskCountByUser *TaskCountLoader
	TaskByUser      *TaskSliceLoader
}

func LoaderMiddleware() echo.MiddlewareFunc{
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ldrs := Loaders{}
			wait := 250 * time.Microsecond
			ldrs.UserById = &UserLoader{
				wait:     wait,
				maxBatch: 100,
				fetch: func(keys []int) (users []*model.User, errors []error) {
					var keySql []string
					users = make([]*model.User, len(keys))
					errors = make([]error, len(keys))
					idx := 0
					for _, key := range keys {
						keySql = append(keySql, strconv.Itoa(key))
					}
					db := database.NewDB()
					time.Sleep(5 * time.Millisecond)
					query := db.Table("users").Where("id in (?)", strings.Join(keySql, ","))
					rows, err := query.Rows()
					if err != nil {
						errors = append(errors, err)
						return users, errors
					}
					for i, _ := range keys {
						rows.Next()
						u := &model.User{}
						err := db.ScanRows(rows, u)
						if err != nil {
							errors = append(errors, err)
						}
						users[i] = u
						idx += 1
					}
					return users, errors
				},
			}
			ldrs.TaskCountByUser = &TaskCountLoader{
				wait:     wait,
				maxBatch: 100,
				fetch: func(keys []int) (ints []*int, errors []error) {
					var cnt int
					ints = make([]*int, len(keys))
					db := database.NewDB()
					err := db.Model(&model.Task{}).Count(&cnt).Error
					if err != nil {
						panic(err)
					}
					for i, _ := range keys {
						ints[i] = &cnt
					}
					return ints, nil
				},
			}
			ldrs.TaskByUser = &TaskSliceLoader{
				wait:     wait,
				maxBatch: 100,
				fetch: func(keys []int) (tasks [][]*model.Task, errors []error) {
					var keySql []string
					tasks = make([][]*model.Task, len(keys))
					errors = make([]error, len(keys))
					db := database.NewDB()
					time.Sleep(5 * time.Millisecond)
					for _, key := range keys {
						keySql = append(keySql, strconv.Itoa(key))
					}
					query := db.Table("tasks").
						Where("user_refer in (?)", strings.Join(keySql, ","))
					rows, err := query.Rows()
					if err != nil {
						errors = append(errors, err)
						return tasks, errors
					}
					for i, key := range keys {
						for rows.Next() {
							t := &model.Task{}
							err := db.ScanRows(rows, t)
							if err != nil {
								panic(err)
							}
							if int(t.UserRefer) == key {
								tasks[i] = append(tasks[i], t)
							}
						}
					}
					return tasks, errors
				},
			}

			request := c.Request()
			ctx := context.WithValue(request.Context(), ctxKey, ldrs)
			c.SetRequest(request.WithContext(ctx))
			err := next(c)
			return err
		}
	}
}


func CtxLoaders(ctx context.Context) (Loaders, error) {
	gCtx := ctx.Value(ctxKey)
	if gCtx == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return Loaders{}, err
	}
	ldr := gCtx.(Loaders)
	return ldr, nil
}
