package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Resume 简历
type Resume struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	AccountID  string        `bson:"accountId" json:"accountId,omitempty"`
	Name       string        `bson:"name" json:"name,omitempty"`
	Phone      string        `bson:"phone" json:"phone,omitempty"`
	Email      string        `bson:"email" json:"email,omitempty"`
	Gender     string        `bson:"gender" json:"gender,omitempty"`
	Birth      time.Time     `bson:"birth" json:"birth,omitempty"`
	Bio        string        `bson:"bio" json:"bio,omitempty"`
	Status     string        `bson:"status" json:"status,omitempty"`
	Projects   []Project     `bson:"projects" json:"projects,omitempty"`
	Educations []Education   `bson:"educations" json:"educations,omitempty"`
	CreatedAt  time.Time     `bson:"createdAt" json:"createdAt"`
	UpdateAt   time.Time     `bson:"updateAt" json:"updateAt,omitempty"`
}

//Education 教育经历
type Education struct {
	Name        string `bson:"name" json:"name"`               //学校
	Major       string `bson:"major" json:"major"`             //专业
	Degree      string `bson:"degree" json:"degree"`           //学历
	StartDate   string `bson:"startAt" json:"startAt"`         //开始日期
	EndDate     string `bson:"endAt" json:"endAt"`             //结束日期
	Description string `bson:"description" json:"description"` //在校经历，描述
}

//Project 工作经历
type Project struct {
	StartDate string   `bson:"startAt" json:"startAt,omitempty"`   //开始日期
	EndDate   string   `bson:"endAt" json:"endAt,omitempty"`       //结束日期
	Position  string   `bson:"position" json:"position,omitempty"` //职位
	Content   string   `bson:"content" json:"content,omitempty"`   //工作内容
	Name      string   `bson:"name" json:"name,omitempty"`         //公司名称
	Tags      []string `bson:"tags" json:"tags,omitempty"`         //标签
}
