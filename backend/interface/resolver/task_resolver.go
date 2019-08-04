package resolver

import (
	"context"
	"github.com/KouT127/gin-sample/backend/interface/graphql/graph"
)

type taskResolver struct{ *Resolver }

func (r *taskResolver) ID(ctx context.Context, obj *graph.Task) (string, error) {
	return obj.ID, nil
}
func (r *taskResolver) User(ctx context.Context, obj *graph.Task) (*graph.User, error) {
	id := int(obj.UserRefer)
	return r.Query().User(ctx, &id)
}