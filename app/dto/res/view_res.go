package res

import "table-saas/app/dto/vo"

type ViewSettingRes struct {
	View *vo.ViewVo `json:"view"`
	Header[] vo.HeaderVo `json:"header"`
	Button[] vo.ButtonVo `json:"button"`
	Filter[] vo.FilterVo `json:"filter"`
	Enum[] vo.EnumVo `json:"enum"`
	Operator[] vo.OperatorVo `json:"operator"`
}

type ViewAll struct {
	ViewList []*vo.ViewVo `json:"viewList"`
}