package main

import (
	"log"
	"sync"
	"testfastprint/db"
	"testfastprint/src"
	"testfastprint/util"
)

func main() {

	env, err := util.LoadConfig(".env")
	if err != nil {
		log.Fatal("cannot load env")
	}

	pg := db.Database(env.DBDriver, env.DBSource)
	db.RunDBMigration(env.MigrationURL, env.DBSource)
	defer pg.Close()

	server := src.InitServer(pg, env)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {

		defer wg.Done()
		server.Start()
	}()

	wg.Wait()
}
