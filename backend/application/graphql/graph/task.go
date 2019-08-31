package graph

import (
	"github.com/KouT127/gin-sample/backend/domain/model"
	"github.com/KouT127/gin-sample/backend/util"
	"strconv"
)

const taskKey = "task:"

type TaskConnection struct {
	TotalCount int
	PageInfo   *PageInfo
	Edges      []*TaskEdge
}
type TaskEdge struct {
	Cursor string
	Node   *Task
}

type Task struct {
	ID          string
	UserRefer   int
	Title       string
	Description string
}

func (c *TaskConnection) registerConnection(cnt int, es []*TaskEdge) *TaskConnection {
	c.PageInfo = &PageInfo{}
	c.Edges = es
	c.TotalCount = cnt
	return c
}

func (e *TaskEdge) registerEdge(t *Task, offset int) *TaskEdge {
	e.Cursor = util.Base64Encode(model.CursorKey + strconv.Itoa(offset))
	e.Node = t
	return e
}

func (t *Task) registerTask(m *model.Task) *Task {
	t.ID = util.Base64Encode(taskKey + strconv.Itoa(int(m.ID)))
	t.Title = m.Title
	t.Description = m.Description
	t.UserRefer = m.UserRefer
	return t
}

func (c *ItemConnection) registerPageInfo(e []*ItemEdge) *PageInfo {
	p := &PageInfo{
		//StartCursor:     util.Base64Encode(c.Edges[0].Cursor),
		//EndCursor:       util.Base64Encode(c.Edges[len(c.Edges)-1].Cursor),
		HasNextPage:     c.TotalCount < len(c.Edges),
		HasPreviousPage: c.TotalCount > len(c.Edges),
	}
	return p
}

func NewItemEdge(m *model.Item, offset int) *ItemEdge {
	t := &Item{}
	t.registerTask(m)
	edge := &ItemEdge{}
	edge = edge.registerEdge(t, offset)
	return edge
}

func NewItemConnection(cnt int, edge []*ItemEdge) *ItemConnection {
	conn := &ItemConnection{}
	conn.TotalCount = cnt
	conn.PageInfo = conn.registerPageInfo(edge)
	conn.Edges = edge
	return conn
}
