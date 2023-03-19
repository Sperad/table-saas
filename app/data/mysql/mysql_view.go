package mysql

import (
	mysql2 "table-saas/app/components/mysql"
)

type View struct {
	Id uint64
	ViewId uint64
	Gid uint64
	Uid uint64
	Title string
	PageSize string
	State int
	IsDeleted int
}

func (view View) SelectById(id uint64) View  {
	db, _ := mysql2.GetMysqlPool()
	dataView := View{}
	db.Table("t_view").Where("view_id = ?", id).Find(&dataView)
	return dataView
}