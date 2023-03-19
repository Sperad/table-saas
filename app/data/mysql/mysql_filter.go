package mysql

import (
	mysql2 "table-saas/app/components/mysql"
)

type Filter struct {
	Id uint64
	FilterId uint64
	ViewId uint64
	Key string
	Type string
	Multiple int
	EnumKey string
	State int
	IsDeleted int
}

type FilterList []Filter

func (filter FilterList) SelectByViewId(viewId uint64) FilterList {
	db, _ := mysql2.GetMysqlPool()
	filterData :=FilterList{}
	db.Table("t_filter").Where("view_id = ?", viewId).Find(&filterData)
	return filterData
}