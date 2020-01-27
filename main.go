package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func createDatabase(fileName string) *sql.DB {
	database, _ := sql.Open("sqlite3", "./appinventory.db")
	file, e := os.Open(fileName)
	r := bufio.NewReader(file)
	s, _, e := r.ReadLine()
	for e == nil {
		create_command := string(s)
		statement, _ := database.Prepare(create_command)
		statement.Exec()
		s, _, e = r.ReadLine()
		statement.Close()
	}
	file.Close()
	return database
}
func main() {
	ctx := context.Background()
	database := createDatabase("./schemas/appinventory_schema.sql")
	db_api := New(database)
	//Create 2 hosts
	hostParams := CreateHostParams{ID: "host1", Name: sql.NullString{"winhost1", true}, AtlasID: sql.NullString{"atlas1", true}, LastUpdated: sql.NullInt64{1, true}}
	db_api.CreateHost(ctx, hostParams)
	hostParams2 := CreateHostParams{ID: "host2", Name: sql.NullString{"winhost2", true}, AtlasID: sql.NullString{"atlas2", true}, LastUpdated: sql.NullInt64{1, true}}
	db_api.CreateHost(ctx, hostParams2)
	//Create 1 asset
	assetParams := CreateAssetParams{ID: "sql1", Name: sql.NullString{"Win sql 1", true}, LastUpdated: sql.NullInt64{1, true}}
	db_api.CreateAsset(ctx, assetParams)
	//Asset is on 2 hosts
	asset := sql.NullString{"sql1", true}
	db_api.CreateAssetHost(ctx, CreateAssetHostParams{ID: "1", HostID: sql.NullString{"host1", true}, AssetID: asset})
	db_api.CreateAssetHost(ctx, CreateAssetHostParams{ID: "2", HostID: sql.NullString{"host2", true}, AssetID: asset})
	hosts, err := db_api.GetHostsForAsset(ctx, asset)
	if err == nil {
		for _, element := range hosts {
			fmt.Println("Got host:", element.Name)
		}
	}
	database.Close()
}
