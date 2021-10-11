package cassandra

import (
	"github.com/gocql/gocql"
)

var session *gocql.Session

func init() {
	var err error
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

// Program is crashing over here
func GetSession() *gocql.Session {
	return session
}
