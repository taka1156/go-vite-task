dev:
	go run main.go

mig-help:
	sql-migrate --help

mig-new:
	sql-migrate new ${file}

mig-up:
	sql-migrate up

mig-down:
	sql-migrate down

gqlgen-init:
	go run github.com/99designs/gqlgen init

gqlgen:
	go run github.com/99designs/gqlgen generate

xo:
	xo schema -S xo.go -o entity/model/xo mysql://root:root_password@localhost/sample

sqlc:
	sqlc generate

skeema-init:
	skeema init -h localhost -P 3306 -u root -proot_password -d schemas --schema sample

skeema-pull:
	skeema pull

skeema-push:
	go run skeema push --allow-unsafe

skeema-diff:
	go run skeema diff --allow-unsafe
