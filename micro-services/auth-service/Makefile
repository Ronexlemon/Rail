# -----------------------
# VARIABLES
# -----------------------
GO := go
TOPIC_CMD := ./shared/kafka/topics/cmd/register.go
PRISMA_CMD := github.com/steebchen/prisma-client-go
PRISMA_GENERATE_CMD := github.com/steebchen/prisma-client-go

PROTO_DIR := ./proto
PROTO_OUT := ./proto

# -----------------------
# TARGETS
# -----------------------

.PHONY: all
all: topics dbpush dbgenerate protos

.PHONY: topics
topics:
	@echo "🚀 Creating Kafka topics..."
	@$(GO) run $(TOPIC_CMD)
	@echo "✅ Kafka topics successfully created!"

.PHONY: dbpush
dbpush:
	@echo "🚀 Running Prisma DB push..."
	@$(GO) run $(PRISMA_CMD) db push
	@echo "✅ Prisma DB push complete!"

.PHONY: dbgenerate
dbgenerate:
	@echo "🚀 Running Prisma DB generate..."
	@$(GO) run $(PRISMA_GENERATE_CMD) generate
	@echo "✅ Prisma DB generate complete!"

.PHONY: protos
protos:
	@echo "🚀 Compiling gRPC proto files..."
	protoc --go_out=$(PROTO_OUT) --go-grpc_out=$(PROTO_OUT) $(PROTO_DIR)/*.proto
	@echo "✅ Proto compilation complete!"

.PHONY: clean
clean:
	@echo "🧹 Cleaning Go build cache..."
	@$(GO) clean -cache -modcache
	@echo "✅ Clean complete!"
