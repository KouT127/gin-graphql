package model

import (
	"errors"
	"github.com/KouT127/gin-sample/backend/util"
	"strconv"
	"strings"
)

const CursorKey = "cursor:"

type Query struct {
	First   int
	After   int
	Last    int
	Before  int
	Keyword string
}

func ParseString(v string) int {
	s, _ := util.Base64Decode(v)
	s = strings.Replace(s, CursorKey, "", 1)
	i, _ := strconv.Atoi(s)
	i++
	return i
}

func NewQuery(first *int, after *string, last *int, before *string, keyword *string) (*Query, error) {
	var f, l, af, bf int
	var kw string
	if first == nil && last == nil {
		return nil, errors.New("invalid query")
	}
	if first != nil {
		f = *first
	}
	if last != nil {
		l = *last
	}
	if after != nil {
		af = ParseString(*after)
	}
	if before != nil {
		bf = ParseString(*before)
	}
	if keyword != nil {
		kw = *keyword
	}
	q := &Query{
		First:   f,
		After:   af,
		Last:    l,
		Before:  bf,
		Keyword: kw,
	}
	return q, nil
}

