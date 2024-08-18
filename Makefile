dev.api:
	cd ./backend && air

dev.admin:
	cd ./gva-admin && pnpm dev

format.admin:
	cd ./gva-admin && pnpm format

swag:
	cd ./backend && make swag

build.admin:
	cd ./gva-admin/app && pnpm run build:dist
