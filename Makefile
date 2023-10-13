qtc:
	@qtc -dir=views -ext=html

watch:
	@go run cmd/fileWatcher/fileWatcher.go

sqlc:
	@sqlc generate

build:
	@go build -ldflags "-w -s" main.go

rename:
	@find . -type f -exec sed -i -e 's/github.com\/Jiang-Gianni\/gthc/github.com\/Jiang-Gianni\/$(name)/g' {} \;