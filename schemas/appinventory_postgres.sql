CREATE TABLE hosts(id VARCHAR PRIMARY KEY,name VARCHAR, atlas_id VARCHAR, address VARCHAR, last_updated BIGINT);
CREATE TABLE assets(id VARCHAR PRIMARY KEY,name VARCHAR ,last_updated BIGINT);
CREATE TABLE asset_hosts(id VARCHAR PRIMARY KEY,host_id VARCHAR,asset_id VARCHAR,FOREIGN KEY(host_id) REFERENCES hosts(id),FOREIGN KEY(asset_id) REFERENCES asset(id));
