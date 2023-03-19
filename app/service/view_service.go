package service

import (
	"table-saas/app/data/mysql"
	"table-saas/app/dto/vo"
)

type ViewService struct {

}

func (ser ViewService) GetViewAll(uid uint64, groupKey string) []*vo.ViewVo {
	viewModel := mysql.View{}
	dataView := viewModel.SelectAll(uid, groupKey)
	viewVo := vo.NewViewVo(dataView)
	return viewVo
}

func (ser ViewService) GetViewByViewId(id uint64 ) *vo.ViewVo {
	viewModel := mysql.View{}
	dataView := viewModel.SelectById(id)
	viewVo := vo.NewViewVoSingle(dataView)
	return viewVo
}

func (ser ViewService) GetHeaderByViewId(id uint64) []vo.HeaderVo {
	headerModel := mysql.HeaderList{}
	dataHeader := headerModel.SelectByViewId(id)
	headerVo := vo.NewHeaderVo(dataHeader)
	return headerVo
}

func (ser ViewService) GetButtonByViewId(id uint64) []vo.ButtonVo {
	buttonModel := mysql.ButtonList{}
	dataButton := buttonModel.SelectByViewId(id)
	buttonVo := vo.NewButtonVo(dataButton)
	return buttonVo
}

func (ser ViewService) GetFilterByViewId(id uint64) []vo.FilterVo {
	filterModel := mysql.FilterList{}
	dataFilter := filterModel.SelectByViewId(id)
	filterVo := vo.NewFilterVo(dataFilter)
	return filterVo
}

func (ser ViewService) GetEnumByViewId(id uint64) []vo.EnumVo {
	enumModel := mysql.EnumList{}
	dataEnum := enumModel.SelectByViewId(id)
	enumVo := vo.NewEnumVo(dataEnum)
	return enumVo
}

func (ser ViewService) GetOperatorByViewId(id uint64) []vo.OperatorVo {
	operatorModel := mysql.OperatorList{}
	dataOperator := operatorModel.SelectByViewId(id)
	operatorVo := vo.NewOperatorVo(dataOperator)
	return operatorVo
}
