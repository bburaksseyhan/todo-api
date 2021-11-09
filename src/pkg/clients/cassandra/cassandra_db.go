package cassandra

import "github.com/gocql/gocql"

func ConnectDatabase(url string, keyspace string) *gocql.Session {

	cluster := gocql.NewCluster(url)
	cluster.Keyspace = keyspace

	session, _ := cluster.CreateSession()

	return session
}
