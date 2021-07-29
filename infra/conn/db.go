package conn

import (
	"TODO_LIST/app/common"
	"TODO_LIST/app/utils/consts"
	"TODO_LIST/infra/logger"

	mgo "gopkg.in/mgo.v2"
)

func ConnectDb() *mgo.Database {
	sess, err := mgo.Dial(consts.HostName)
	common.CheckErr(err)
	sess.SetMode(mgo.Monotonic, true)
	db := sess.DB(consts.DbName)
	logger.Info("Databases connection successful...")
	return db
}
