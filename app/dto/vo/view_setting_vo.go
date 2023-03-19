package vo

type ViewSettingVo struct {
	View *ViewVo `json:"view"`
	Header[] HeaderVo
	Button[] ButtonVo
	Filter[] FilterVo
	Enum[] EnumVo
	Operator[] OperatorVo
}
