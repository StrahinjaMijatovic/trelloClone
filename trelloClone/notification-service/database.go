package main

import (
	"log"

	"github.com/gocql/gocql"
)

var cassandraSession *gocql.Session

func ConnectDatabase() {
	cluster := gocql.NewCluster("cassandra")
	cluster.Keyspace = "notifications"
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Failed to connect to Cassandra: %v", err)
	}

	cassandraSession = session
	log.Println("Connected to Cassandra")
}

func CloseDatabase() {
	if cassandraSession != nil {
		cassandraSession.Close()
	}
}
