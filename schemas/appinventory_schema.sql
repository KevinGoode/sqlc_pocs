CREATE TABLE hosts(id TEXT PRIMARY KEY,name TEXT, atlas_id TEXT, address TEXT, last_updated BIGINT);
CREATE TABLE assets(id TEXT PRIMARY KEY,name TEXT ,last_updated BIGINT, some_real REAL, some_double DOUBLE PRECISION, some_var_char_255 VARCHAR(255) );
CREATE TABLE asset_hosts(id VARCHAR PRIMARY KEY,host_id VARCHAR,asset_id VARCHAR,FOREIGN KEY(host_id) REFERENCES hosts(id),FOREIGN KEY(asset_id) REFERENCES asset(id));
