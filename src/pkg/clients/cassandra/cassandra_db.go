package cassandra

import "github.com/gocql/gocql"

func ConnectDatabase(url string, keyspace string) *gocql.Session {

	cluster := gocql.NewCluster(url)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.Quorum

	session, _ := cluster.CreateSession()

	return session
}
