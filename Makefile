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

gzip:
	find ./assets -type f -exec gzip -v "{}" \; -exec mv "{}.gz" "{}" \;

gentest:
	find . -name '*.go' -not -path './cmd/*' -not -path './views/*' -not -name '*_test.go' -not -name 'main.go' -exec gotests -all -w -i -parallel {} +

cover:
	go test -coverprofile=cover.out ./... && go tool cover -html=cover.out -o=cover.html && brave-browser cover.html

lint:
	golangci-lint --enable-all -v run --disable depguard && go vet

doc:
	brave-browser http://localhost:7000/github.com/Jiang-Gianni/htmx-go && pkgsite -http=:7000