qtc:
	@qtc -dir=views -ext=html

watch:
	@go run cmd/fileWatcher/fileWatcher.go

sqlc:
	@sqlc generate

build:
	@go build -ldflags="-w -s -X main.environment=${ENV} -X main.gitCommit=${GIT_COMMIT}" main.go

rename:
	@find . -type f -exec sed -i -e 's/github.com\/Jiang-Gianni\/htmx-go/github.com\/Jiang-Gianni\/$(name)/g' {} \;

css:
	@sass --style=compressed style/main.scss assets/style.css