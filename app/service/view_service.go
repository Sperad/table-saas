package service

import (
	"sync"
	"table-saas/app/data/mysql"
	"table-saas/app/dto/vo"
)

type ViewService struct {

}

var wg sync.WaitGroup

func (ser ViewService) GetSetting(viewId uint64 ) *vo.ViewSettingVo {
	var setting = vo.ViewSettingVo{}
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
		//setting.Filter = ser.GetFilterByViewId(viewId)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		//setting.Enum = ser.GetEnumByViewId(viewId)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		//setting.Operator = ser.GetOperatorByViewId(viewId)
	}()

	wg.Wait()
	return &setting
}

func (ser ViewService) GetViewByViewId(id uint64 ) *vo.ViewVo {
	viewModel := mysql.View{}
	dataView := viewModel.SelectById(id)
	viewVo := vo.NewViewVo(dataView)
	return viewVo
}

func (ser ViewService) GetHeaderByViewId(id uint64) []vo.HeaderVo {
	headerModel := mysql.HeaderList{}
	dataHeader := headerModel.SelectByViewId(id)
	headerVo := vo.NewHeaderVo(dataHeader)
	return headerVo
}

func (ser ViewService) GetButtonByViewId(id uint64) []vo.ButtonVo {
	buttonModel := mysql.Button{}
	dataButton := buttonModel.SelectByViewId(id)
	buttonVo := vo.NewButtonVo(dataButton)
	return buttonVo
}

func (ser ViewService) GetFilterByViewId(id uint64) *vo.FilterVo {
	filterModel := mysql.Filter{}
	dataFilter := filterModel.SelectByViewId(id)
	filterVo := vo.NewFilterVo(dataFilter)
	return filterVo
}

func (ser ViewService) GetEnumByViewId(id uint64) *vo.EnumVo {
	enumModel := mysql.Enum{}
	dataEnum := enumModel.SelectByViewId(id)
	enumVo := vo.NewEnumVo(dataEnum)
	return enumVo
}

func (ser ViewService) GetOperatorByViewId(id uint64) *vo.OperatorVo {
	operatorModel := mysql.Operator{}
	dataOperator := operatorModel.SelectByViewId(id)
	operatorVo := vo.NewOperatorVo(dataOperator)
	return operatorVo
}
