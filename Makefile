migrate: 
	migrate -source file://postgres/migrations  \
		-database postgres://crypto:@127.0.0.1:5432/crypto?sslmode=disable up

rollback: 
	migrate -source file://postgres/migrations  \
		-database postgres://crypto:@127.0.0.1:5432/crypto?sslmode=disable down
drop: 
	migrate -source file://postgres/migrations  \
		-database postgres://crypto:@127.0.0.1:5432/crypto?sslmode=disable drop

migration:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir postgres/migrations $$name

###

sqlc: 
	sqlc generate

rsync:
	rsync -avz  "/Users/alex/go/src/cr-tick/cr-tick" "alex@3dexp.com:www/cr-tick" 
#	rsync -avz  "/Users/alex/go/src/cr-tick/.env" "alex@3dexp.com:www/cr-tick" 


	