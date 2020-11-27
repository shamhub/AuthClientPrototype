.PHONY: install

install:
	go install github.com/shamhub/codingtest/internal/data
	go install github.com/shamhub/codingtest/internal/platform
	go install github.com/shamhub/codingtest/api
	go install github.com/shamhub/codingtest/cmd/authclient