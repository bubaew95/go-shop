mup:
	migrate -path migrations -database $${DB_SHOP} up $(count)
mdown:
	migrate -path migrations -database $${DB_SHOP} down $(count)
migrate:
	migrate create -ext sql -dir migrations $(name)
mock:
	### product mock
	mockgen -source=./internal/application/product/domain/product.go -destination=./internal/application/product/infra/postgresql/mock/product.go -package=mock \
	&& mockgen -source=./internal/application/category/domain/category.go -destination=./internal/application/category/infra/postgresql/mock/category.go -package=mock

lint:
	goimports -local "github.com/bubaew95/go_shop" -w .