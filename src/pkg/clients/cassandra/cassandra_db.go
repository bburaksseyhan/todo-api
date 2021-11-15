package cassandra

import "github.com/gocql/gocql"

/*
  ConnectDatabase(url string, keyspace string) take 2 parameters
  Connect New Clustor and return *gocql.Session
*/
func ConnectDatabase(url string, keyspace string) *gocql.Session {

	cluster := gocql.NewCluster(url)
	cluster.Keyspace = keyspace

	session, _ := cluster.CreateSession()

	return session
}
