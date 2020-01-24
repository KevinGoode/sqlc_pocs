-- name: GetAllAssets :many
SELECT * FROM assets;
-- name: GetAllHosts :many
SELECT * FROM hosts;
-- name: GetAllAssetHosts :many
SELECT * FROM asset_hosts;
-- name: GetAsset :one
SELECT * FROM assets WHERE id = $1 LIMIT 1;
-- name: GetHost :one
SELECT * FROM hosts WHERE id = $1 LIMIT 1;
-- name: GetHostsForAsset :many
SELECT name, address FROM hosts INNER JOIN asset_hosts ON hosts.id = asset_hosts.host_id WHERE asset_hosts.asset_id=$1;
-- name: CreateHost :exec
INSERT INTO hosts (id,name,atlas_id,last_updated) VALUES ($1, $2, $3, $4);
-- name: CreateAsset :exec
INSERT INTO assets (id,name,last_updated) VALUES ($1, $2, $3);
-- name: CreateAssetHost :exec
INSERT INTO asset_hosts (id,host_id, asset_id) VALUES ($1, $2, $3);
-- name: UpdateHostAddress :exec
UPDATE hosts SET address = $1 WHERE id = $2;
-- name: DeleteHost :exec
DELETE FROM hosts WHERE id = $1;
-- name: DeleteAsset :exec
DELETE FROM assets WHERE id = $1;
-- name: DeleteAssetHost :exec
DELETE FROM asset_hosts WHERE id = $1;


