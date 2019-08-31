package resolver

import (
	"context"
	"github.com/KouT127/gin-sample/backend/application/graphql/graph"
)

type itemResolver struct{ *Resolver }

func (r *itemResolver) Price(ctx context.Context, obj *graph.Item) (float64, error) {
	return float64(obj.Price), nil
}
