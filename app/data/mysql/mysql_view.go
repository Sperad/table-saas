package mysql

import (
	mysql2 "table-saas/app/components/mysql"
)

var tableName = "t_view"

type View struct {
	Id uint64
	ViewId uint64
	Gid uint64
	Uid uint64
	GroupKey string
	Title string
	PageSize string
	State int
	IsDeleted int
}

type ViewList []View

func (view View) SelectById(id uint64) View  {
	db, _ := mysql2.GetMysqlPool()
	dataView := View{}
	db.Table(tableName).Where("view_id = ?", id).Find(&dataView)
	return dataView
}

func (view View) SelectAll(uid uint64, groupKey string) ViewList  {
	db, _ := mysql2.GetMysqlPool()
	dataList := ViewList{}
	db.Table(tableName).
				Where("uid = ?", uid).
				Where("group_key = ?", groupKey).
				Find(&dataList)

	return dataList
}