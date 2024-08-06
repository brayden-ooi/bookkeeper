This project has been archived

1. Enter `docker compose watch` to rebuild docker image on save
2. Enter `docker compose down` to remove the stack
3. Enter `sudo docker image prune` to remove dangling images

Schema migration
1. go to root and `cd sql/schema && goose sqlite3 ../../index.db up`

Generating go counterparts with SQLC
1. go to root and `sqlc generate`


TODO
1. /handler/transaction.go
  - transactions currently assume users will only submit transactions under the latest financial year and the latest transaction_counter
  - transactions currently assume users will only submit 2 entries
2. create users settings table one-to-one
   1. default currency
   2. default year
   3. voucher format string
   4. Ability to reset counter
3. ability to edit submitted transaction is currently disabled
