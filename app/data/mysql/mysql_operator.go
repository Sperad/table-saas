package mysql

import (
	mysql2 "table-saas/app/components/mysql"
)

type Operator struct {
	Id uint64
	OperatorId uint64
	ViewId uint64
	Key uint64
	Title uint64
	Event uint64
	State int
	IsDeleted int
}

type OperatorList []Operator

func (operator OperatorList) SelectByViewId(viewId uint64)  OperatorList {
	db, _ := mysql2.GetMysqlPool()
	operatorData := OperatorList{}
	db.Table("t_operator").Where("view_id = ?", viewId).Find(&operatorData)
	return operatorData
}