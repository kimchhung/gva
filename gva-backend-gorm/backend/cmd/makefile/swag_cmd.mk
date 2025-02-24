swag:
	make swag.admin
	make swag.bot

swag.admin:
	swag init -g main.go -d ./api/admin --parseDependency --parseInternal --output ./api/admin/docs --instanceName admin

swag.bot:
	swag init -g main.go -d ./api/bot --parseDependency --parseInternal --output ./api/bot/docs --instanceName bot
