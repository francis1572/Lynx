package models

import "github.com/mitchellh/mapstructure"

//User structure
type Enumerable struct {
	DataList []interface{} `bson:"dataList" json:"dataList"`
}

type IEnumerable interface{}

func (l *Enumerable) Decode(structList []interface{}) []interface{} {
	for i, obj := range l.DataList {
		mapstructure.Decode(obj, structList[i])
	}
	return structList
}
