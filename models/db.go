// init mongodb session and provide common functions

package models

import (
	"fmt"
	"time"

	conf "github.com/irisnet/iris-api-server/configs"
	"github.com/irisnet/iris-api-server/modules/logger"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	session *mgo.Session
)

func Init() {
	if session == nil {
		url := fmt.Sprintf("mongodb://%s:%s", conf.ConfMongodb.Host, conf.ConfMongodb.Port)

		logger.Info.Printf("Mgo start on %s\n", url)

		var err error
		session, err = mgo.Dial(url)
		if err != nil {
			logger.Error.Fatalln(err)
		}
		session.SetMode(mgo.Monotonic, true)
	}
}

func InitWithAuth(addrs []string, username, password string) {
	dialInfo := &mgo.DialInfo{
		Addrs:     addrs, //[]string{"192.168.6.122"}
		Direct:    false,
		Timeout:   time.Second * 1,
		Database:  conf.ConfMongodb.DbName,
		Username:  username,
		Password:  password,
		PoolLimit: 4096, // Session.SetPoolLimit
	}

	session, err := mgo.DialWithInfo(dialInfo)
	session.SetMode(mgo.Monotonic, true)
	if nil != err {
		panic(err)
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
		//先按照关键字查询，如果存在，直接返回
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
		//先按照关键字查询，如果存在，直接返回
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

/**
 * 执行查询，此方法可拆分做为公共方法
 * [SearchPerson description]
 * @param {[type]} collectionName string [description]
 * @param {[type]} query          bson.M [description]
 * @param {[type]} sort           bson.M [description]
 * @param {[type]} fields         bson.M [description]
 * @param {[type]} skip           int    [description]
 * @param {[type]} limit          int)   (results      []interface{}, err error [description]
 */
func Query(collectionName string, query bson.M, sort string, fields bson.M, skip int, limit int) (results []interface{}, err error) {
	exop := func(c *mgo.Collection) error {
		return c.Find(query).Sort(sort).Select(fields).Skip(skip).Limit(limit).All(&results)
	}
	return results, ExecCollection(collectionName, exop)
}
