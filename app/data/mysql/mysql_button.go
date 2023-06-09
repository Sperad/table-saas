package mysql

import (
	mysql2 "table-saas/app/components/mysql"
)

type Button struct {
	Id uint64
	ButtonId uint64
	ViewId uint64
	Key string
	Title string
	Event string
	Pid uint64
	State int
	IsDeleted int
}

type ButtonList []Button

func (button ButtonList) SelectByViewId(viewId uint64) ButtonList {
	db, _ := mysql2.GetMysqlPool()
	buttonData := ButtonList{}
	db.Table("t_button").Where("view_id = ?", viewId).Find(&buttonData)
	return buttonData
}