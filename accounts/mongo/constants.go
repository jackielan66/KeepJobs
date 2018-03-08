package mongo

import (
	"log"
	"os"

	"github.com/laidingqing/KeepJobs/common/config"
	mgo "gopkg.in/mgo.v2"
)

var session *mgo.Session

//DatabaseName ..
var DatabaseName = config.Database.DatabaseName

//AccountCollectionName account table name
var AccountCollectionName = "accounts"

func fatalError(err error) {
	log.Printf("mongodb error")
	os.Exit(1)
}
