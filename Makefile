migrate:
	@echo "start migrate..."
	@migrate -source file://backend/infrastracture/database/migration/  -database 'mysql://root:@tcp(localhost:3306)/go_tutorial' up

show-migrations:
	 mysqldef -uroot go_tutorial --export > schema.sql  

mysqldef-dry:
	 mysqldef -uroot go_tutorial --dry-run < schema.sql

mysqldef:
	 mysqldef -uroot go_tutorial < schema.sql

gql-generate:
	@echo "GraphQL models generate..."
	@cd backend/interface/graphql/ && gqlgen -v

web:
	make web-install
	make web-start

web-install:
	yarn

web-start:
	@cd frontend/ && yarn start
