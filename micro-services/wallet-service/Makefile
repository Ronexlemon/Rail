# -----------------------
# VARIABLES. solc --abi erc20.sol -o .
#abigen --abi=IERC20.abi --pkg=token --out=erc20.go
# -----------------------
GO := go
TOPIC_CMD := ./shared/kafka/topics/cmd/register.go
PRISMA_CMD := github.com/steebchen/prisma-client-go
PRISMA_GENERATE_CMD := github.com/steebchen/prisma-client-go

PROTO_DIR := ./proto
PROTO_OUT := ./proto


SOLC := solc
ABIGEN := abigen
ABI_DIR := ./utils/helpers/tokens/abi   # folder containing the contract
SOL_CONTRACT := $(ABI_DIR)/erc20.sol
ABI_OUT := $(ABI_DIR)/IERC20.abi
GO_OUT := $(ABI_DIR)/erc20.go
PKG := token



# -----------------------
# TARGETS
# -----------------------

.PHONY: all
all: topics dbpush dbgenerate protos contract

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
	protoc \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/*.proto
	@echo "✅ Proto compilation complete!"


.PHONY: contract
contract:
	@echo "🚀 Compiling Solidity contract and generating Go bindings..."
	$(SOLC) --abi ./utils/helpers/tokens/abi/erc20.sol -o ./utils/helpers/tokens/abi
	$(ABIGEN) --abi= ./utils/helpers/tokens/abi/IERC20.abi --pkg=token --out=erc20.go
	@echo "✅ Contract Go bindings generated!"
	
.PHONY: clean
clean:
	@echo "🧹 Cleaning Go build cache..."
	@$(GO) clean -cache -modcache
	@echo "✅ Clean complete!"
