APP_NAME := ai_insights
ENTRY := ./cmd/server/main.go
RUN_CMD := ./cmd/server/main.go


.PHONY: all run 

all: run

run:
	@echo "Running Go application build"
	go run $(RUN_CMD)
