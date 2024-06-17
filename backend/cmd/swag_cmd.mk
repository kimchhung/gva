swag:
	make swag.web
	make swag.admin
	make swag.crawler

swag.format:
	swag fmt

swag.web:
	swag init -g main.go -d ./api/web --parseDependency --output ./api/web/docs/ --instanceName web

swag.admin:
	swag init -g main.go -d ./api/admin --parseDependency --output ./api/admin/docs --instanceName admin

swag.crawler:
	swag init -g main.go -d ./api/crawler --parseDependency --output ./api/crawler/docs --instanceName crawler