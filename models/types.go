// interface for a document

package models

import (
	"gopkg.in/mgo.v2"
)

type Document interface {
	// collection name
	Name() string
	// primary key pair(used to find a unique record)
	PkKvPair() map[string]interface{}
}
