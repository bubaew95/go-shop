mup:
	migrate -path migrations -database $${DB_SHOP} up $(count)
mdown:
	migrate -path migrations -database $${DB_SHOP} down $(count)
migrate:
	migrate create -ext sql -dir migrations $(name)