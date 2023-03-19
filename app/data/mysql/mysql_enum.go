package mysql

import (
	mysql2 "table-saas/app/components/mysql"
)

type Enum struct {
	Id uint64
	EnumId uint64
	ViewId uint64
	Key string
	Label string
	Value string
	State int
	IsDeleted int
}

type EnumList []Enum

func (enum EnumList) SelectByViewId(viewId uint64) EnumList {
	db, _ := mysql2.GetMysqlPool()
	enumData := EnumList{}
	db.Table("t_enum").Where("view_id = ?", viewId).Find(&enumData)
	return enumData
}