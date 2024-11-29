include .env

db-migration:
	echo "Running migration into go-todo-app-pg postgresql container" 
	cat ./db/migration.sql | docker exec -i go-todo-app-pg psql -U ${POSTGRES_USER} -d ${POSTGRES_DB}
	
run:
	go run cmd/go-todo-app/main.go
