1. Enter `docker compose watch` to rebuild docker image on save
2. Enter `docker compose down` to remove the stack
3. Enter `sudo docker image prune` to remove dangling images

Schema migration
1. go to root and `cd sql/schema && goose sqlite3 ../../index.db up`

Generating go counterparts with SQLC
1. go to root and `sqlc generate`