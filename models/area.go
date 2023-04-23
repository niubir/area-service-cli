package models

// 多个区域
type Areas []Area

// 区域
type Area struct {
	// 行政区名称
	Name string `json:"name"`
	// 行政区划级别
	Level Constant `json:"level"`
	// 区域编码
	Code string `json:"code"`
	// 区域中心点
	Center string `json:"center"`
	// 下级行政区列表
	Subs Areas `json:"subs"`
}
