golang-driver
=============

Golang wrapper of the DataStax/Cassandra [C/C++ driver](https://github.com/datastax/cpp-driver)

Basic support for prepared statements and ad hoc queries. Lacking support for collections, but that will be remedied shortly.

### Build

1. Build and install the DataStax [C/C++ driver](https://github.com/datastax/cpp-driver)
1. Install `go get github.com/mstump/golang-driver/cassandra`
1. Run the example `go run $GOPATH/src/github.com/mstump/golang-driver/examples/basic.go`

### Example Usage

```go
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
```
