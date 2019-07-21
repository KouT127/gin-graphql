package model

type PageInfo struct {
	StartCursor     string
	EndCursor       string
	HasNextPage     bool
	HasPreviousPage bool
}

//type Node struct {
//	id string
//}

//func (f *FriendsConnection) TotalCount() int {
//	return len(f.Ids)
//}
//
//func (f *FriendsConnection) PageInfo() PageInfo {
//	return PageInfo{
//		StartCursor: utt(f.From),
//		EndCursor:   EncodeCursor(f.To - 1),
//		HasNextPage: f.To < len(f.Ids),
//	}
//}
