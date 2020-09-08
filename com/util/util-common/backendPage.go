package util_common

import "container/list"

type BackendPage struct {
	Code int
	Msg string
	Count int
	Data list.List
}

func (bp *BackendPage) BackendPage(count int,data list.List) {
	bp.Code=0
	bp.Msg=""
	bp.Count=count
	bp.Data=data
	return
}

