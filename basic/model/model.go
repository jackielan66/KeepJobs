package model

import "gopkg.in/mgo.v2/bson"

const (
	//DictCity 城市
	DictCity = iota + 1
	//DictPosition 职位
	DictPosition
	//DictIndustry 行业
	DictIndustry
	//DictWorkExp 工作经验
	DictWorkExp
	//DictDegree 学历
	DictDegree
	//DictSalary 薪资
	DictSalary
	//DictWorkStatus 工作状态
	DictWorkStatus
)

//Dict 字典代码表
type Dict struct {
	ID        bson.ObjectId `bson:"_id" json:"id,omitempty"`
	Classify  int32         `bson:"classify" json:"classify,omitempty"`
	Code      string        `bson:"code" json:"code,omitempty"`
	Parent    string        `bson:"parent" json:"parent,omitempty"`
	Name      string        `bson:"name" json:"name,omitempty"`
	FirstChar string        `bson:"firstChar" json:"firstChar,omitempty"`
	Rank      string        `bson:"rank" json:"rank,omitempty"`
}
