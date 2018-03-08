package model

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	//AuthenticationWeixin 微信开放平台
	AuthenticationWeixin = iota + 1
)

const (
	//AuthCreated 新建未处理
	AuthCreated = iota + 1
	//AuthPass 通过
	AuthPass
	//AuthReject 拒绝
	AuthReject
)

const (
	//AuthorityUser 用户
	AuthorityUser = iota + 1
	//AuthorityMaintain 运营
	AuthorityMaintain
	//AuthorityAdmin 管理员
	AuthorityAdmin
)

const (
	//UserStatusNormal 正常
	UserStatusNormal = iota + 1
	//UserStatusLocked 锁定
	UserStatusLocked
	//UserStatusClosed 关闭
	UserStatusClosed
)

// Account 用户账号信息
type Account struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	AccountID   string        `bson:"-" json:"id"`
	UserName    string        `bson:"username" json:"username,omitempty"`
	Password    string        `bson:"password" json:"password,omitempty"`
	Salt        string        `bson:"salt" json:"-"`
	CreatedAt   time.Time     `bson:"createdAt" json:"createdAt,omitempty"`
	Token       string        `bson:"token" json:"-,omitempty"`
	UserInfoRef mgo.DBRef     `bson:"infoRef" json:"-"`
	UserInfo    UserInfo      `bson:"userInfo" json:"userInfo,omitempty"`
	Authority   []int32       `bson:"authority" json:"authority,omitempty"`
	Status      int32         `bson:"status" json:"status,omitempty"`
}

//Authentication 第三方授权认证登记表
type Authentication struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	AccountID string        `bson:"accountId" json:"accountId,omitempty"`
	OpenID    string        `bson:"openId" json:"openId,omitempty"`
	Channal   int32         `bson:"plateform" json:"plateform,omitempty"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt,omitempty"`
}

//UserInfo ...
type UserInfo struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	AccountID string        `bson:"accountId" json:"accountId"`
	AvatarURL string        `bson:"avatarURL" json:"avatarURL,omitempty"`
	City      string        `bson:"city" json:"city,omitempty"`
	Country   string        `bson:"country" json:"country,omitempty"`
	Gender    int           `bson:"gender" json:"gender,omitempty"`
	NickName  string        `bson:"nickName" json:"nickName,omitempty"`
	Province  string        `bson:"province" json:"province,omitempty"`
}
