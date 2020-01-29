package main

import (
	"context"
	"database/sql"
	"fmt"

	"./db"
	_ "github.com/mattn/go-sqlite3"
	migrate "github.com/rubenv/sql-migrate"
)

func createDatabase(fileName string) *sql.DB {
	database, _ := sql.Open("sqlite3", "./appinventory.db")
	// Read migrations from a folder:
	migrations := &migrate.FileMigrationSource{Dir: "./schemas"}
	n, err := migrate.Exec(database, "sqlite3", migrations, migrate.Up)
	if err != nil {
		// Handle errors!
	}
	fmt.Printf("Applied %d migrations!\n", n)
	return database
}
func createDatabaseFromBinaryFile(fileName string) *sql.DB {
	database, _ := sql.Open("sqlite3", "./appinventory.db")
	// Read migrations from binary data: bindata.go
	migrations := &migrate.AssetMigrationSource{
		Asset:    db.Asset,
		AssetDir: db.AssetDir,
		Dir:      "schemas",
	}
	n, err := migrate.Exec(database, "sqlite3", migrations, migrate.Up)
	if err != nil {
		// Handle errors!
	}
	fmt.Printf("Applied %d migrations!\n", n)
	return database
}
func main() {
	ctx := context.Background()
	//database := createDatabase("./schemas/appinventory_schema.sql")
	database := createDatabaseFromBinaryFile("./schemas/appinventory_schema.sql")
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
