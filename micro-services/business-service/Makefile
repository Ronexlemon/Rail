# -----------------------
# VARIABLES
# -----------------------
GO := go
# Path to the Go file for Kafka topic registration
TOPIC_CMD := ./shared/kafka/topics/cmd/register.go
# Prisma client command (assuming this is the path/module used to invoke prisma commands)
PRISMA_CMD := github.com/steebchen/prisma-client-go
PRISMA_GENERATE_CMD := github.com/steebchen/prisma-client-go 
PROTO_DIR := ./proto
PROTO_OUT := ./proto

# -----------------------
# TARGETS
# -----------------------

# Default target runs all setup steps: topics, dbpush, dbgenerate
.PHONY: all
all: topics dbpush dbgenerate protos

# Run the Kafka topic creation script
.PHONY: topics
topics:
	@echo "🚀 Creating Kafka topics..."
	@$(GO) run $(TOPIC_CMD)
	@echo "✅ Kafka topics successfully created!"

# Run Prisma DB push (migrates schema to the database)
.PHONY: dbpush
dbpush:
	@echo "🚀 Running Prisma DB push..."
	@$(GO) run $(PRISMA_CMD) db push
	@echo "✅ Prisma DB push complete!"

# Run Prisma DB generate (generates the Go client based on the schema)
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

# Clean build cache (optional)
.PHONY: clean
clean:
	@echo "🧹 Cleaning Go build cache..."
	@$(GO) clean -cache -modcache
	@echo "✅ Clean complete!"