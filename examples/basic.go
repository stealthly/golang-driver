package main

import (
	"fmt"
	"golang-driver/cassandra"
)

func main() {
	cluster := cassandra.NewCluster()
	cluster.SetContactPoints("127.0.0.1")

	session := cassandra.NewSession()
	future := cluster.SessionConnect(session)
	future.Wait()

	statement := cassandra.NewStatement("select cluster_name from system.local;", 0)
	future := session.Execute(statement)
	future.Wait()

	result := future.Result()

	fmt.Printf("Clusters:\r\n")
	for result.Next() {
		var clusterName string
		result.Scan(&clusterName)
		fmt.Printf("%s\n", clusterName)
	}

	fmt.Printf("DONE.\r\n")
}
