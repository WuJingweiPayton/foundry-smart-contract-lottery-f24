package label

// nolint
const (
	MaxInFlightTransactionsWarning          = `WARNING: If this happens a lot, you may need to increase EVM.Transactions.MaxInFlight to boost your node's transaction throughput, however you do this at your own risk. You MUST first ensure your ethereum node is configured not to ever evict local transactions that exceed this number otherwise the node can get permanently stuck. See the performance guide for more details: https://docs.chain.link/docs/evm-performance-configuration/`
	MaxQueuedTransactionsWarning            = `WARNING: Hitting EVM.Transactions.MaxQueued is a sanity limit and should never happen under normal operation. Unless you are operating with very high throughput, this error is unlikely to be a problem with your Chainlink node configuration, and instead more likely to be caused by a problem with your eth node's connectivity. Check your eth node: it may not be broadcasting transactions to the network, or it might be overloaded and evicting Chainlink's transactions from its mempool. It is recommended to run Chainlink with multiple primary and sendonly nodes for redundancy and to ensure fast and reliable transaction propagation. Increasing EVM.Transactions.MaxQueued will allow Chainlink to buffer more unsent transactions, but you should only do this if you need very high burst transmission rates. If you don't need very high burst throughput, increasing this limit is not the correct action to take here and will probably make things worse. See the performance guide for more details: https://docs.chain.link/docs/evm-performance-configuration/`
	NodeConnectivityProblemWarning          = `WARNING: If this happens a lot, it may be a sign that your eth node has a connectivity problem, and your transactions are not making it to any miners. It is recommended to run Chainlink with multiple primary and sendonly nodes for redundancy and to ensure fast and reliable transaction propagation. See the performance guide for more details: https://docs.chain.link/docs/evm-performance-configuration/`
	RPCTxFeeCapConfiguredIncorrectlyWarning = `WARNING: Gas price was rejected by the eth node for being too high. By default, go-ethereum (and clones) have a built-in upper limit for gas price. It is preferable to disable this and rely Chainlink's internal gas limits instead. Your RPC node's RPCTxFeeCap needs to be disabled or increased (recommended configuration: --rpc.gascap=0 --rpc.txfeecap=0). If you want to limit Chainlink's max gas price, you may do so by setting EVM.GasEstimator.PriceMax on the Chainlink node. Chainlink will never send a transaction with a total cost higher than EVM.GasEstimator.PriceMax. See the performance guide for more details: https://docs.chain.link/docs/evm-performance-configuration/`
)
