go fmt $(go list ./...)
go vet $(go list ./...)
go test $(go list ./...) -v -test.short
golangci-lint run ./...
