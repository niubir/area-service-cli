package models

type Constant int

const (
	UNKNOW Constant = 0

	// 国家
	AREA_LEVEL_COUNTRY Constant = 1
	// 省份（直辖市会在province显示）
	AREA_LEVEL_PROVINCE Constant = 2
	// 市（直辖市会在province显示）
	AREA_LEVEL_CITY Constant = 3
	// 区县
	AREA_LEVEL_DISTRICT Constant = 4
	// 街道
	AREA_LEVEL_STREET Constant = 5
)
