package mysql

import (
	mysql2 "table-saas/app/components/mysql"
)

type Header struct {
	Id uint64
	HeaderId uint64
	ViewId uint64
	Title string
	Key string
	Type string
	Sortable int
	Filterable int
	Display string
	Format string
	State int
	IsDeleted int
}

type HeaderList []Header

func (header HeaderList) SelectByViewId(viewId uint64) []Header {
	db, _ := mysql2.GetMysqlPool()
	dataHeader := HeaderList{}
	db.Table("t_header").Where("view_id = ?", viewId).Find(&dataHeader)
	return dataHeader
}