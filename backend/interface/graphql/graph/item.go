package graph

import (
	"github.com/KouT127/gin-sample/backend/domain/model"
	"github.com/KouT127/gin-sample/backend/util"
	"strconv"
)

const itemKey = "item:"

type ItemConnection struct {
	TotalCount int
	PageInfo   *PageInfo
	Edges      []*ItemEdge
}
type ItemEdge struct {
	Cursor string
	Node   *Item
}

type Item struct {
	ID          string
	Name        string
	Description string
	Price       float32
}

type Cart struct {
	CartItems []*Item
}

func (c *ItemConnection) registerConnection(cnt int, es []*ItemEdge) *ItemConnection {
	c.PageInfo = &PageInfo{}
	c.Edges = es
	c.TotalCount = cnt
	return c
}

func (e *ItemEdge) registerEdge(t *Item, offset int) *ItemEdge {
	e.Cursor = util.Base64Encode(model.CursorKey + strconv.Itoa(offset))
	e.Node = t
	return e
}

func (t *Item) registerTask(m *model.Item) *Item {
	t.ID = util.Base64Encode(taskKey + strconv.Itoa(int(m.ID)))
	t.Name = m.Name
	t.Description = m.Description
	t.Price = m.Price
	return t
}

func (c *TaskConnection) registerPageInfo(e []*TaskEdge) *PageInfo {
	p := &PageInfo{
		//StartCursor:     util.Base64Encode(c.Edges[0].Cursor),
		//EndCursor:       util.Base64Encode(c.Edges[len(c.Edges)-1].Cursor),
		HasNextPage:     c.TotalCount < len(c.Edges),
		HasPreviousPage: c.TotalCount > len(c.Edges),
	}
	return p
}

func NewTaskEdge(m *model.Task, offset int) *TaskEdge {
	t := &Task{}
	t.registerTask(m)
	edge := &TaskEdge{}
	edge = edge.registerEdge(t, offset)
	return edge
}

func NewTaskConnection(cnt int, edge []*TaskEdge) *TaskConnection {
	conn := &TaskConnection{}
	conn.TotalCount = cnt
	conn.PageInfo = conn.registerPageInfo(edge)
	conn.Edges = edge
	return conn
}
