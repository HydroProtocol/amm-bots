package utils

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"math/big"
	"strings"
)

func (erc *ERC20) init(web3Url string) {
	resp, err := Web3Call(web3Url, erc.Address, "0x313ce567")
	if err != nil {
		logrus.Panic(fmt.Sprintf("get %s decimal failed", erc.Symbol), err)
	}
	var dataContainer IJsonRpcResString
	json.Unmarshal([]byte(resp), &dataContainer)
	erc.Decimal = int(ParseHexToBigint(dataContainer.Result[2:]).Int64())
	erc.Initialized = true
}

func (erc *ERC20) GetBalance(web3url string, address string) (balance *decimal.Decimal, rawBalance *decimal.Decimal, err error) {
	if !erc.Initialized {
		erc.init(web3url)
	}
	data := "0x70a08231" + ExtendAddressTo256bit(address)
	resp, err := Web3Call(web3url, erc.Address, data)
	if err != nil {
		return &decimal.Zero, &decimal.Zero, errors.New("get erc20 balance failed")
	}
	var dataContainer IJsonRpcResString
	json.Unmarshal([]byte(resp), &dataContainer)
	rawValue := ParseHexToDecimal(dataContainer.Result[2:], 0)
	value := ParseHexToDecimal(dataContainer.Result[2:], int32(erc.Decimal*-1))
	return value, rawValue, nil
}

func Web3Call(url string, contractAddress string, data string) (string, error) {
	callParams := struct {
		To   string `json:"to"`
		Data string `json:"data"`
	}{contractAddress, data}
	callParamsBytes, _ := json.Marshal(callParams)
	dataString := `{"jsonrpc":"2.0","method":"eth_call","params": [` + string(callParamsBytes) + `, "latest"],"id":1}`
	return Post(
		url,
		dataString,
		EmptyKeyPairList,
		[]KeyPair{{"Content-Type", "application/json"}},
	)
}

func Stripe0xPrefix(str string) string {
	if len(str) < 2 {
		return str
	} else {
		if str[0:2] == "0x" {
			return str[2:]
		} else {
			return str
		}
	}
}

func PrivateKeyToAddress(privateKey string) string {
	pk, err := crypto.HexToECDSA(Stripe0xPrefix(privateKey))
	if err != nil {
		panic(err)
	}
	address := crypto.PubkeyToAddress(pk.PublicKey).Hex()
	return strings.ToLower(address)
}

func SignString(privateKey string, msg string) string {
	bytesMsg := []byte(msg)
	msgHash := crypto.Keccak256([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(bytesMsg), bytesMsg)))
	pk, err := crypto.HexToECDSA(Stripe0xPrefix(privateKey))
	if err != nil {
		panic(err)
	}
	sig, _ := crypto.Sign(msgHash, pk)

	return fmt.Sprintf("0x%x", sig)
}

func toOrderSignature(sign []byte) []byte {
	var res [96]byte
	copy(res[:], []byte{sign[64] + 27})
	copy(res[32:], sign[:64])
	return res[:]
}

func SignOrderId(privateKey string, orderId string) string {
	byteOrderId, _ := hex.DecodeString(Stripe0xPrefix(orderId))
	byteOrderId = crypto.Keccak256([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(byteOrderId), byteOrderId)))
	pk, err := crypto.HexToECDSA(Stripe0xPrefix(privateKey))
	if err != nil {
		panic(err)
	}
	sig, _ := crypto.Sign(byteOrderId, pk)
	orderSign := toOrderSignature(sig)
	return fmt.Sprintf("0x%x", orderSign)
}

func ParseHexToBigint(hex string) *big.Int {
	if hex == "0x" {
		return big.NewInt(0)
	}
	i := new(big.Int)
	i.SetString(Stripe0xPrefix(hex), 16)
	return i
}

func ParseHexToDecimal(hex string, exp int32) *decimal.Decimal {
	num := decimal.NewFromBigInt(ParseHexToBigint(hex), exp)
	return &num
}

func ExtendAddressTo256bit(address string) string {
	switch len(address) {
	case 40:
		return "000000000000000000000000" + address
	case 42:
		return "000000000000000000000000" + address[2:]
	default:
		panic(fmt.Sprintf("address %s not valid", address))
	}
}
