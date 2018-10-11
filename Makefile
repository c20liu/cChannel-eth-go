bind:
	@echo Generating Go bindings...
	abigen --sol cChannel-eth/contracts/GenericChannel.sol --pkg channel --out channel/generic_channel.go
	abigen --sol cChannel-eth/contracts/HTLRegistry.sol --pkg htlregistry --out htlregistry/registry.go
	abigen --sol cChannel-eth/contracts/VirtContractResolver.sol --pkg resolver --out resolver/resolver.go
	abigen --sol cChannel-eth/contracts/DepositPool.sol --pkg dpool --out dpool/deposit_pool.go
