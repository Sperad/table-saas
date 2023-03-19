package service

import (
	"sync"
	"table-saas/app/dto/res"
)

type ViewLogic struct {
}


var wg sync.WaitGroup

func (logic ViewLogic) GetSetting(viewId uint64 ) *res.ViewSettingRes {
	ser := ViewService{}
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