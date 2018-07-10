// init mongodb session and provide common functions

package models

import (
	"fmt"
	"time"

	conf "github.com/irisnet/irishub-server/configs"
	"github.com/irisnet/irishub-server/modules/logger"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	session *mgo.Session
)

func init() {
	InitWithAuth()
}

func InitWithAuth() {
	addr := fmt.Sprintf("%s:%s", conf.ConfMongodb.Host, conf.ConfMongodb.Port)
	addrs := []string{addr}

	dialInfo := &mgo.DialInfo{
		Addrs:     addrs,
		Database:  conf.ConfMongodb.DbName,
		Username:  conf.ConfMongodb.User,
		Password:  conf.ConfMongodb.Password,
		Direct:    false,
		Timeout:   time.Second * 10,
		PoolLimit: 4096, // Session.SetPoolLimit
	}

	var err error
	session, err = mgo.DialWithInfo(dialInfo)
	session.SetMode(mgo.Monotonic, true)
	if err != nil {
		logger.Error.Panicln(err)
	}
}

func getSession() *mgo.Session {
	// max session num is 4096
	return session.Clone()
}

// get collection object
func ExecCollection(collection string, s func(*mgo.Collection) error) error {
	session := getSession()
	defer session.Close()
	c := session.DB(conf.ConfMongodb.DbName).C(collection)
	return s(c)
}

func Find(collection string, query interface{}) *mgo.Query {
	session := getSession()
	defer session.Close()
	c := session.DB(conf.ConfMongodb.DbName).C(collection)
	return c.Find(query)
}

func Save(doc Document) error {
	save := func(c *mgo.Collection) error {
		n, _ := c.Find(doc.PkKvPair()).Count()
		if n >= 1 {
			logger.Info.Println("db: record existed while save data")
			return nil
		}
		logger.Info.Printf("insert %s  %+v\n", doc.Name(), doc)
		return c.Insert(doc)
	}

	return ExecCollection(doc.Name(), save)
}

func SaveOrUpdate(h Document) error {
	save := func(c *mgo.Collection) error {
		n, err := c.Find(h.PkKvPair()).Count()
		logger.Info.Printf("Count:%d err:%+v\n", n, err)
		if n >= 1 {
			return Update(h)
		}
		logger.Info.Printf("insert %s  %+v\n", h.Name(), h)
		return c.Insert(h)
	}

	return ExecCollection(h.Name(), save)
}

func Update(h Document) error {
	update := func(c *mgo.Collection) error {
		key := h.PkKvPair()
		logger.Info.Printf("update %s set %+v where %+v\n", h.Name(), h, key)
		return c.Update(h.PkKvPair(), h)
	}
	return ExecCollection(h.Name(), update)
}

func Query(collectionName string, query bson.M, fields bson.M, skip int, limit int, sorts ...string) (results []interface{}, err error) {
	exop := func(c *mgo.Collection) error {
		return c.Find(query).Select(fields).Sort(sorts...).Skip(skip).Limit(limit).All(&results)
	}
	return results, ExecCollection(collectionName, exop)
}
