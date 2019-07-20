package dataloader

import (
	"context"
	"fmt"
	. "github.com/KouT127/gin-sample/backend/domain/model"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)


type ctxKeyType struct{ name string }

var ctxKey = ctxKeyType{"userCtx"}

type Loaders struct {
	UserById   *UserLoader
	TaskByUser *TaskSliceLoader
}

func LoaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ldrs := Loaders{}
		wait := 250 * time.Microsecond

		ldrs.UserById = &UserLoader{
			wait:     wait,
			maxBatch: 100,
			fetch: func(keys []int) ([]*User, []error) {
				var keySql []string
				for _, key := range keys {
					keySql = append(keySql, strconv.Itoa(key))
				}

				fmt.Printf("SELECT * FROM address WHERE id IN (%s)\n", strings.Join(keySql, ","))
				time.Sleep(5 * time.Millisecond)

				users := make([]*User, len(keys))
				errors := make([]error, len(keys))
				for i, key := range keys {
					print(key)
					users[i] = &User{
						Name:     "a",
						BirthDay: "b",
						Gender:   "c",
						PhotoURL: "d",
						Active:   false,
					}
				}
				return users, errors
			},
		}

		ldrs.TaskByUser = &TaskSliceLoader{
			wait:     wait,
			maxBatch: 100,
			fetch: func(keys []int) ([][]*Task, []error) {
				var keySql []string
				for _, key := range keys {
					keySql = append(keySql, strconv.Itoa(key))
				}

				fmt.Printf("SELECT * FROM orders WHERE customer_id IN (%s)\n", strings.Join(keySql, ","))
				time.Sleep(5 * time.Millisecond)

				tasks := make([][]*Task, len(keys))
				errors := make([]error, len(keys))
				for i, key := range keys {
					print(key)
					tasks[i] = []*Task{}
				}

				return tasks, errors
			},
		}
		ctx := context.WithValue(c.Request.Context(), ctxKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func CtxLoaders(ctx context.Context) Loaders {
	return ctx.Value(ctxKey).(Loaders)
}
