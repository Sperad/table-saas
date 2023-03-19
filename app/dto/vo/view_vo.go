package vo

import "table-saas/app/data/mysql"

type ViewVo struct {
	ViewId uint64
	Gid uint64
	Uid uint64
	Title string
	PageSize string
	State int
}

type HeaderVo struct {
	HeaderId uint64
	ViewId uint64
	Title string
	Key string
	Type string
	Sortable int
	Filterable int
	Display string
	Format string
	State int
}

type ButtonVo struct {
	ButtonId uint64
	ViewId uint64
	Key string
	Title string
	Event string
	Pid uint64
	Icon string
	State int
}

type FilterVo struct {
	FilterId uint64
	ViewId uint64
	Key string
	Type string
	Multiple int
	EnumKey string
	State int
}

type EnumVo struct {
	EnumId uint64
	ViewId uint64
	Key string
	Label string
	Value string
	State int
}

type OperatorVo struct {
	OperatorId uint64
	ViewId uint64
	Key uint64
	Title uint64
	Event uint64
	State int
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

func NewHeaderVo(headerList []mysql.Header) []HeaderVo {
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

func NewButtonVo(buttonList []mysql.Button) []ButtonVo {
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

func NewFilterVo(filter mysql.Filter) *FilterVo {
	return &FilterVo {
		FilterId : filter.FilterId,
		ViewId : filter.ViewId,
		Key : filter.Key,
		Type : filter.Type,
		Multiple : filter.Multiple,
		EnumKey : filter.EnumKey,
		State : filter.State,
	}
}

func NewEnumVo(enum mysql.Enum) *EnumVo {
	return &EnumVo {
		EnumId : enum.EnumId,
		ViewId : enum.ViewId,
		Key : enum.Key,
		Label : enum.Label,
		Value : enum.Value,
		State : enum.State,
	}
}

func NewOperatorVo(operator mysql.Operator) *OperatorVo {
	return &OperatorVo {
		OperatorId : operator.OperatorId,
		ViewId : operator.ViewId,
		Key :operator.Key,
		Title :operator.Title,
		Event : operator.Event,
		State : operator.State,
	}
}
