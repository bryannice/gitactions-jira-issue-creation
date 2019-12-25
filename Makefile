# -------------
# VARIABLES
# -------------

# -------------
# FUNCTIONS
# -------------


# -------------
# TASKS
# -------------
.PHONY: fmt
fmt:
	@gofmt -w -s -d configuration
	@gofmt -w -s -d main.go

.PHONY: build
build: fmt
	@go mod download && \
     go get ./... && \
     go install ./...