package dataloader

import "errors"

type Query struct {
	First   *int
	After   *string
	Last    *int
	Before  *string
	Keyword *string
}

func registerQuery(first *int, after *string, last *int, before *string, keyword *string) *Query {
	q := &Query{
		First:   first,
		After:   after,
		Last:    last,
		Before:  before,
		Keyword: keyword,
	}
	return q
}

func NewQuery(first *int, after *string, last *int, before *string, keyword *string) (*Query, error) {
	if first == nil && last == nil {
		return nil, errors.New("invalid query")
	}
	q := &Query{
		First:   first,
		After:   after,
		Last:    last,
		Before:  before,
		Keyword: keyword,
	}
	return q, nil
}
