postgres:
	docker run -p 5432:5432 --name hospital-management -e POSTGRES_USER=user1 -e POSTGRES_PASSWORD=test -e POSTGRES_DB=hospital -d postgres
migrateup:
	migrate -path db/migration -database "postgresql://user1:test@localhost/hospital?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://user1:test@localhost:5432/hospital?sslmode=disable" -verbose down

migrationversion:
	migrate -path db/migration -database "postgresql://user1:test@localhost:5432/hospital?sslmode=disable" version

cleanDatabase: 
	migrate -path db/migration -database "postgresql://user1:test@localhost:5432/hospital?sslmode=disable" drop


	
.PHONY:	postgres	migrateup	migratedown	migrationversion	cleanDatabase