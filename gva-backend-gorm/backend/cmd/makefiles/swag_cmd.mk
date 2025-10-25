swag:
	make swag.admin
	make swag.bot
	make swag.web

swag.admin:
	swag init -g main.go -d ./app/admin --parseDependency --parseInternal --output ./app/admin/docs --instanceName admin

swag.bot:
	swag init -g main.go -d ./app/bot --parseDependency --parseInternal --output ./app/bot/docs --instanceName bot

swag.web:
	swag init -g main.go -d ./app/web --parseDependency --parseInternal --output ./app/web/docs --instanceName web
