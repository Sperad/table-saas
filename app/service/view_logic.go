package service

import (
	"sync"
	"table-saas/app/dto/res"
)

type ViewLogic struct {

}

var ser = ViewService{}
var wg sync.WaitGroup

func (logic ViewLogic) GetAll(uid uint64, groupKey string)  res.ViewAll {
	var res = res.ViewAll{}
	res.ViewList = ser.GetViewAll(uid, groupKey)
	return res
}

func (logic ViewLogic) GetSetting(viewId uint64 ) *res.ViewSettingRes {
	var setting = res.ViewSettingRes{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		setting.View = ser.GetViewByViewId(viewId)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		setting.Header = ser.GetHeaderByViewId(viewId)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		setting.Button = ser.GetButtonByViewId(viewId)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		setting.Filter = ser.GetFilterByViewId(viewId)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		setting.Enum = ser.GetEnumByViewId(viewId)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		setting.Operator = ser.GetOperatorByViewId(viewId)
	}()

	wg.Wait()
	return &setting
}