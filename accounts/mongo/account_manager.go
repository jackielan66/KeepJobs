package mongo

import (
	"log"
	"time"

	"github.com/laidingqing/KeepJobs/accounts/model"
	"github.com/laidingqing/KeepJobs/common/config"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//AccountManager user manager struct
type AccountManager struct {
	session *mgo.Session
}

// NewAccountManager new manager
func NewAccountManager() *AccountManager {
	session, err := mgo.Dial(config.Database.HostURI)
	if err != nil {
		fatalError(err)
	}
	return &AccountManager{
		session: session,
	}
}

// CopySession for more flexible use
func (um *AccountManager) CopySession() *mgo.Session {
	return um.session.Copy()
}

//Insert 根据用户名查询用户
func (um *AccountManager) Insert(user model.Account) (model.Account, error) {
	copySession := um.CopySession()
	defer copySession.Close()
	query := copySession.DB(config.Database.DatabaseName).C(AccountCollectionName)
	user.ID = bson.NewObjectId()
	user.CreatedAt = time.Now()
	err := query.Insert(user)
	if err != nil {
		log.Printf("db manager err: %s", err.Error())
		return model.Account{}, err
	}
	return user, nil
}

//FindByID find account by id
func (um *AccountManager) FindByID(id string) (model.Account, error) {
	var acct model.Account
	copySession := um.CopySession()
	defer copySession.Close()
	coll := copySession.DB(config.Database.DatabaseName).C(AccountCollectionName)
	err := coll.FindId(bson.ObjectIdHex(id)).One(&acct)
	if err != nil {
		return model.Account{}, err
	}
	return acct, nil
}

//FindByUserName 根据用户名查询用户
func (um *AccountManager) FindByUserName(username string) (model.Account, error) {
	copySession := um.CopySession()
	defer copySession.Close()
	var user model.Account
	query := copySession.DB(config.Database.DatabaseName).C(AccountCollectionName)
	err := query.Find(bson.M{"username": username}).One(&user)
	if err != nil {
		return model.Account{}, err
	}
	return user, nil
}

//FindAccountByToken 根据当前TOKEN查询账号
func (um *AccountManager) FindAccountByToken(token string) (model.Account, error) {
	var acct model.Account
	copySession := um.CopySession()
	defer copySession.Close()
	coll := copySession.DB(config.Database.DatabaseName).C(AccountCollectionName)
	err := coll.Find(bson.M{"token": token}).One(&acct)
	if err != nil {
		return model.Account{}, err
	}
	return acct, nil
}

//UpdateCurrentToken 更新会话账号TOKEN
func (um *AccountManager) UpdateCurrentToken(id string, token string) error {
	copySession := um.CopySession()
	defer copySession.Close()
	coll := copySession.DB(config.Database.DatabaseName).C(AccountCollectionName)
	err := coll.UpdateId(bson.ObjectIdHex(id), bson.M{"$set": bson.M{"token": token}})
	if err != nil {
		return err
	}
	return nil
}
