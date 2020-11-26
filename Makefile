.PHONY: install

install:
	go install github.com/shamhub/codingtest/internal/data
	go install github.com/shamhub/codingtest/internal/platform
	go install github.com/shamhub/codingtest/cmd
	go install github.com/shamhub/codingtest