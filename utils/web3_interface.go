package utils

type NodeBlockNum struct {
	blockNum int64
	nodeUrl  string
}

type ERC20 struct {
	Symbol      string
	Address     string
	Decimal     int
	Initialized bool
}

type EthereumLog struct {
	Address          string   `json:"address"`
	BlockHash        string   `json:"blockHash"`
	BlockNumber      string   `json:"blockNumber"`
	Data             string   `json:"data"`
	LogIndex         string   `json:"logIndex"`
	Removed          bool     `json:"removed"`
	Topics           []string `json:"topics"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
}

type EthereumReceipt struct {
	BlockHash         string        `json:"blockHash"`
	BlockNumber       string        `json:"blockNumber"`
	ContractAddress   interface{}   `json:"contractAddress"`
	CumulativeGasUsed string        `json:"cumulativeGasUsed"`
	From              string        `json:"from"`
	GasUsed           string        `json:"gasUsed"`
	Logs              []EthereumLog `json:"logs"`
	LogsBloom         string        `json:"logsBloom"`
	Status            string        `json:"status"`
	To                string        `json:"to"`
	TransactionHash   string        `json:"transactionHash"`
	TransactionIndex  string        `json:"transactionIndex"`
}

type IJsonRpcResString struct {
	Result string `json:"result"`
}

type IJsonRpcResLogs struct {
	Result []EthereumLog `json:"result"`
}

type IJsonRpcResReceipt struct {
	Result EthereumReceipt `json:"result"`
}

type IJsonRpcBlockInfo struct {
	ID     int `json:"id"`
	Result struct {
		Difficulty       string   `json:"difficulty"`
		ExtraData        string   `json:"extraData"`
		GasLimit         string   `json:"gasLimit"`
		GasUsed          string   `json:"gasUsed"`
		Hash             string   `json:"hash"`
		LogsBloom        string   `json:"logsBloom"`
		Miner            string   `json:"miner"`
		MixHash          string   `json:"mixHash"`
		Nonce            string   `json:"nonce"`
		Number           string   `json:"number"`
		ParentHash       string   `json:"parentHash"`
		ReceiptsRoot     string   `json:"receiptsRoot"`
		Sha3Uncles       string   `json:"sha3Uncles"`
		Size             string   `json:"size"`
		StateRoot        string   `json:"stateRoot"`
		Timestamp        string   `json:"timestamp"`
		TotalDifficulty  string   `json:"totalDifficulty"`
		Transactions     []string `json:"transactions"`
		TransactionsRoot string   `json:"transactionsRoot"`
		Uncles           []string `json:"uncles"`
	} `json:"result"`
}
