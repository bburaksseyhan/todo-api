package cassandra

import (
	"github.com/gocql/gocql"
	"github.com/sirupsen/logrus"
)

// ConnectDatabase function Connect New Clustor and return *gocql.Session
func ConnectDatabase(url string, keyspace string) *gocql.Session {

	cluster := gocql.NewCluster(url)
	cluster.Keyspace = keyspace

	session, err := cluster.CreateSession()

	if err != nil {
		logrus.Fatal(err)
	}

	return session
}
