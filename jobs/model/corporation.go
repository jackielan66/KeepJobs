package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Corporation 企业信息
type Corporation struct {
	ID        bson.ObjectId `bson:"_id" json:"id,omitempty"`
	AccountID string        `bson:"accountId" json:"accountId,omitempty"`
	Name      string        `bson:"name" json:"name,omitempty"`
	Alias     string        `bson:"alias" json:"alias,omitempty"`
	Intro     string        `bson:"intro" json:"intro,omitempty"`
	Address   string        `bson:"address" json:"address,omitempty"`
	Pictures  []string      `bson:"pics" json:"pics,omitempty"`
	Province  string        `bson:"province" json:"province,omitempty"`
	City      string        `bson:"city" json:"city,omitempty"`
	Area      string        `bson:"area" json:"area,omitempty"`
	Website   string        `bson:"website" json:"website,omitempty"`
	Industry  string        `bson:"industry" json:"industry,omitempty"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt,omitempty"`
	UpdatedAt time.Time     `bson:"updatedAt" json:"updatedAt,omitempty"`
}

//BusinessInfo 工商登记信息
type BusinessInfo struct {
	ID      bson.ObjectId `bson:"_id" json:"id,omitempty"`
	Code    string        `bson:"code" json:"code,omitempty"`
	Address string        `bson:"address" json:"address,omitempty"`
	Status  string        `bson:"status" json:"status,omitempty"`
	FormOfE string        `bson:"formOfE" json:"formOfE,omitempty"`
	FoundAt string        `bson:"foundAt" json:"foundAt,omitempty"`
	Legal   string        `bson:"legal" json:"legal,omitempty"`
	Capital string        `bson:"capital" json:"capital,omitempty"`
	Scope   string        `bson:"scope" json:"scope,omitempty"`
}

//Job detail.
type Job struct {
	ID          bson.ObjectId `bson:"_id" json:"id,omitempty"`
	CorpID      string        `bson:"corpId" json:"corpId,omitempty"`
	Title       string        `bson:"title" json:"title,omitempty"`
	City        string        `bson:"city" json:"city,omitempty"`
	Experience  string        `bson:"experience" json:"experience,omitempty"`
	Schooling   string        `bson:"schooling" json:"schooling,omitempty"`
	Salary      string        `bson:"salary" json:"salary,omitempty"`
	Description string        `bson:"description" json:"description,omitempty"`
	Tags        []string      `bson:"tags" json:"tags,omitempty"`
	CreatedAt   time.Time     `bson:"createdAt" json:"createdAt,omitempty"`
	UpdatedAt   time.Time     `bson:"updatedAt" json:"updatedAt,omitempty"`
	Status      int32         `bson:"status" json:"status,omitempty"`
}
