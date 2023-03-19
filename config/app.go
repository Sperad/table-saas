package config

type GinApp struct {
	Port int32
}

var App = &GinApp{
	Port: 8880,
}