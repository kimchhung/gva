swag:
	make swag.admin
	make swag.bot
	make swag.web

swag.admin:
	swag init -g main.go -d ./api/admin --parseDependency --parseInternal --output ./api/admin/docs --instanceName admin

swag.bot:
	swag init -g main.go -d ./api/bot --parseDependency --parseInternal --output ./api/bot/docs --instanceName bot

swag.web:
	swag init -g main.go -d ./api/web --parseDependency --parseInternal --output ./api/web/docs --instanceName web
