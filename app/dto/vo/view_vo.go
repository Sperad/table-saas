package vo

import "table-saas/app/data/mysql"

type ViewVo struct {
	ViewId uint64 `json:"viewId"`
	Gid uint64 `json:"gid"`
	Uid uint64 `json:"uid"`
	Title string `json:"title"`
	PageSize string `json:"pageSize"`
	State int  `json:"state"`
}

type HeaderVo struct {
	HeaderId uint64 `json:"headerId"`
	ViewId uint64 `json:"viewId"`
	Title string `json:"title"`
	Key string `json:"key"`
	Type string `json:"type"`
	Sortable int  `json:"sortable"`
	Filterable int `json:"filterable"`
	Display string `json:"display"`
	Format string `json:"format"`
	State int  `json:"state"`
}

type ButtonVo struct {
	ButtonId uint64 `json:"buttonId"`
	ViewId uint64 `json:"viewId"`
	Key string `json:"key"`
	Title string `json:"title"`
	Event string `json:"event"`
	Pid uint64 `json:"pid"`
	Icon string `json:"icon"`
	State int `json:"state"`
}

type FilterVo struct {
	FilterId uint64  `json:"filterId"`
	ViewId uint64 `json:"viewId"`
	Key string `json:"key"`
	Type string `json:"type"`
	Multiple int `json:"multiple"`
	EnumKey string `json:"enumKey"`
	State int `json:"state"`
}

type EnumVo struct {
	EnumId uint64 `json:"enumId"`
	ViewId uint64 `json:"viewId"`
	Key string `json:"key"`
	Label string `json:"label"`
	Value string `json:"value"`
	State int  `json:"state"`
}

type OperatorVo struct {
	OperatorId uint64 `json:"operatorId"`
	ViewId uint64 `json:"viewId"`
	Key uint64 `json:"key"`
	Title uint64 `json:"title"`
	Event uint64 `json:"event"`
	State int  `json:"state"`
}

func NewViewVo(view mysql.View) *ViewVo {
	return &ViewVo {
		ViewId: view.ViewId,
		Gid: view.Gid,
		Uid: view.Uid,
		Title: view.Title,
		PageSize: view.PageSize,
		State: view.State,
	}
}

func NewHeaderVo(headerList mysql.HeaderList) []HeaderVo {
	headerVoList := make([]HeaderVo, len(headerList))
	for i, header := range headerList {
		headerVoList[i] = HeaderVo {
			HeaderId: header.HeaderId,
			ViewId: header.ViewId,
			Title: header.Title,
			Key: header.Key,
			Type: header.Type,
			Sortable: header.Sortable,
			Filterable: header.Filterable,
			Display: header.Display,
			Format: header.Format,
			State: header.State,
		}
	}
	return headerVoList
}

func NewButtonVo(buttonList mysql.ButtonList) []ButtonVo {
	buttonVoList := make([]ButtonVo, len(buttonList))
	for i, button := range buttonList {
		buttonVoList[i] = ButtonVo {
			ButtonId: button.ButtonId,
			ViewId : button.ViewId,
			Key :button.Key,
			Title :button.Title,
			Event : button.Event,
			Pid : button.Pid,
			Icon : button.Icon,
			State :button.State,
		}
	}
	return buttonVoList
}

func NewFilterVo(filterList mysql.FilterList) []FilterVo {
	filterVoList := make([]FilterVo, len(filterList))
	for i, filter := range filterList {
		filterVoList[i] = FilterVo {
			FilterId : filter.FilterId,
			ViewId : filter.ViewId,
			Key : filter.Key,
			Type : filter.Type,
			Multiple : filter.Multiple,
			EnumKey : filter.EnumKey,
			State : filter.State,
		}
	}
	return filterVoList
}

func NewEnumVo(enumList mysql.EnumList) []EnumVo {
	enumVoList := make([]EnumVo, len(enumList))
	for i, enum := range enumList {
		enumVoList[i] = EnumVo {
			EnumId : enum.EnumId,
			ViewId : enum.ViewId,
			Key : enum.Key,
			Label : enum.Label,
			Value : enum.Value,
			State : enum.State,
		}
	}
	return enumVoList
}

func NewOperatorVo(operatorList mysql.OperatorList) []OperatorVo {
	operatorVoList := make([]OperatorVo, len(operatorList))
	for i,operator := range operatorList {
		operatorVoList[i] = OperatorVo {
			OperatorId : operator.OperatorId,
			ViewId : operator.ViewId,
			Key :operator.Key,
			Title :operator.Title,
			Event : operator.Event,
			State : operator.State,
		}
	}
	return operatorVoList
}
