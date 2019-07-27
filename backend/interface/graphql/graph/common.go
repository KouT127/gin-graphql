package graph

const CursorKey = "cursor:"

type PageInfo struct {
	StartCursor     string
	EndCursor       string
	HasNextPage     bool
	HasPreviousPage bool
}

