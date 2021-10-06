package cassandra

import (
	"github.com/gocql/gocql"
)

var cluster *gocql.ClusterConfig

func init() {
	cluster = gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
}

func GetSession() (session *gocql.Session, err error) {
	session, err = cluster.CreateSession()
	return
}
