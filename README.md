# SQLC/SQL-MIGRATE/GO-BINDATA GOLANG CODE GEN POC

This folder contains a poc of using the sqlc golang code generator to interface with a sqllite3 db. 
It also demonstrates the sql-migrate library to create and migrate versions of the database and the use   
of go-bindata to generate a go source file containing a directory of sql schema files in binary format.       
   
Sqlc code generator uses postgres parser BUT regardless, provided you use standard SQL data types and if you follow STANDARD SQL SYNTAX , there is no reason why sqlc cannot be used to access other sql db types. 
   
Suggested data types are as follows (TBC): TEXT, VARCHAR, REAL, DOUBLE PRECISION, BOOLEAN, INT, BIGINT.   
Avoid time and date types (NOT supported by sqllite). Instead, use BIGINT to store a unix time stamp    
   
This article gives a useful comparison of data types: https://www.w3resource.com/sql/data-type.php, BUT might be not wholely correct.   
   
If you dig deeper it transpires that some types are supported as alias/synonyms:   
https://www.w3resource.com/mysql/mysql-data-types.php   
https://www.tutorialspoint.com/sqlite/sqlite_data_types.htm   

## Prequisites
Install go lang   
Install sqllite3   
go get github.com/mattn/go-sqlite3   
go get -v github.com/rubenv/sql-migrate/...   
go get -u github.com/shuLhan/go-bindata/...   
Download sqlc code generator https://bin.equinox.io/c/gvM95th6ps1/sqlc-devel-linux-amd64.tgz   

## Example
The example contains 3 linked tables and generates a full suite of insert, update, select and delete commands.

## Generating code 
1. In directory sqlc_poc> /home/{USERNAME}/Downloads/sqlc-devel-linux-amd64/sqlc generate   
This produces the files: **appinventory_query.sql.go  db.go   models.go  querier.go**   
2. In directory sqlc_poc> go-bindata -pkg db -o ./db/bindata.go schemas   
This produces the file **./db/bindata.go.**   
## Running Basic Example
To test most generated apis, run the code as follows.The code tests a number of inserts and a join  on 2 tables
1. go build
2. ./sqlc_poc   

## Conclusion
The code generated for basic SQL CRUD operations can be used with a SQLLITE3 DB.   
Need to test with MySQL. Should work with postgres. Note how the database is created using sql-migrate   
in the function 'createDatabase'.As the db is extended release by release, scripts in numerical order   
are added to schema directory and the migrate library will automatically upgrade any db it finds by applying   
the scripts in order. State of db is maintained by an automatically created db table:**gorp_migrations** (See below) 


## References For SQLC and slq-migrate

https://conroy.org/introducing-sqlc   
https://github.com/kyleconroy/sqlc   
https://github.com/rubenv/sql-migrate   
https://github.com/shuLhan/go-bindata   

## NOTE
The directory of schema files need to be shipped with any executable that uses this code unless
a tool such as go-bindata is used. Because the code cannot execute without these files the preferred approach is to compile   
the go-bindata generated file in with the dbclient  and thereby remove the runtime dependency of shipping the schema files.   
The code demonstrates both approaches using 'createDatabase' and 'createDatabaseFromBinaryFile'  

## APPENDIX-  sqlite output
[goode@localhost sqlc_poc]$ **sqlite3 appinventory.db**   
SQLite version 3.26.0 2018-12-01 12:34:55   
Enter ".help" for usage hints.   
**sqlite> .tables;**   
asset_hosts  assets gorp_migrations hosts         
**sqlite> .schema hosts;**   
CREATE TABLE hosts(id TEXT PRIMARY KEY,name TEXT, atlas_id TEXT, address TEXT, last_updated BIGINT);   
**sqlite> .schema asset_hosts;**   
CREATE TABLE asset_hosts(id VARCHAR PRIMARY KEY,host_id VARCHAR,asset_id VARCHAR,FOREIGN KEY(host_id) REFERENCES hosts(id),FOREIGN KEY(asset_id) REFERENCES asset(id));   
**sqlite> .schema assets;**   
CREATE TABLE assets(id TEXT PRIMARY KEY,name TEXT ,last_updated BIGINT, some_real REAL, some_double DOUBLE PRECISION, some_var_char_255 VARCHAR(255) );   
**sqlite> .schema gorp_migrations;**   
CREATE TABLE IF NOT EXISTS "gorp_migrations" ("id" varchar(255) not null primary key, "applied_at" datetime);   
**select * from gorp_migrations;**   
01_appinventory_schema.sql|2020-01-28 15:29:19.202651433+00:00   





