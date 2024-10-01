# Include .env file if it exists
ifneq (,$(wildcard .env))
    include .env
    export
endif

# Define the target
deploy-fuji:
	@echo "Running deploy script for RealEstateToken and Issuer to deploy them on Fuji"
	forge script script/DeployContracts.s.sol --rpc-url $(FUJI_RPC_URL) --private-key $(PRIVATE_KEY) --broadcast --verifier-url $(SNOWTRACE_VERIFIER_URL) --etherscan-api-key $(SNOWTRACE_API_KEY)
deploy-sepolia:
	@echo "Running deploy script for RealEstateToken to deploy it on Sepolia"
	forge script script/DeployContracts.s.sol --rpc-url $(SEPOLIA_RPC_URL) --private-key $(PRIVATE_KEY) --broadcast --verify --etherscan-api-key $(ETHERSCAN_API_KEY)

# Optional: Add a help target to describe how to use the Makefile
help:
	@echo "Usage:"
	@echo "  make deploy-fuji - Run the deploy script with the specified environment variables to deploy on Fuji"
	@echo "  make deploy-sepolia - Run the deploy script with the specified environment variables to deploy on Sepolia"
