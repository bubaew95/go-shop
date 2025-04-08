mup:
	migrate -path migrations -database $${DB_SHOP} up $(count)
mdown:
	migrate -path migrations -database $${DB_SHOP} down $(count)
migrate:
	migrate create -ext sql -dir migrations $(name)

mock:
	mockgen -source=./internal/application/product/domain/product.go -destination=./internal/application/product/infra/postgresql/product.go -package=mock -mock_names=ProductRepository=MockProductRepository
