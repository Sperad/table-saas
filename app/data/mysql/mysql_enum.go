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

func (enum Enum) SelectByViewId(viewId uint64) Enum {
	db, _ := mysql2.GetMysqlPool()
	enumData := Enum{}
	db.Table("t_enum").Where("view_id = ?", viewId).Find(&enumData)
	return enumData
}