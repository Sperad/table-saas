package req

type ViewAddReq struct {
	ButtonReq []ButtonReq `json:"buttonReq"`
	EnumReq []EnumReq  `json:"enumReq"`
	FilterReq []FilterReq  `json:"filterReq"`
	HeaderReq []HeaderReq  `json:"headerReq"`
	OperatorReq []OperatorReq  `json:"operatorReq"`
	ViewReq ViewReq  `json:"viewReq"`
}

type ButtonReq struct {
	Key string
	Title string
	Event string
	Pid uint64
}

type EnumReq struct {
	Key string
	Label string
	Value string
}

type FilterReq struct {
	Key string
	Type string
	Multiple int
	EnumKey string
}

type HeaderReq struct {
	Title string
	Key string
	Type string
	Sortable int
	Filterable int
	Display string
}

type OperatorReq struct {
	Key uint64
	Title uint64
	Event uint64
}

type ViewReq struct {
	GroupKey string
	Title string
	PageSize string
}