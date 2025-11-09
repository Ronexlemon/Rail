# -----------------------
# VARIABLES
# -----------------------
GO := go
TOPIC_CMD := ./shared/kafka/topics/cmd/main.go


# -----------------------
# TARGETS
# -----------------------

# Default target
.PHONY: all
all: topics

# Run the topic creation script
.PHONY: topics
topics:
	@echo "🚀 Creating Kafka topics "
	@$(GO) run $(TOPIC_CMD)
	@echo "✅ Kafka topics successfully created!"

# Clean build cache (optional)
.PHONY: clean
clean:
	@echo "🧹 Cleaning Go build cache..."
	@$(GO) clean -cache -modcache
	@echo "✅ Clean complete!"
