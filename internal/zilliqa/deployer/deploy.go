// This golang program attempts to deploy a simple contract using the most priviledged of the 3 golang sdks for Zilliqa in hopes that someday someone might be able to interact with their blockchain outside of Curl...

package deployer

import (
	"io/ioutil"

	"github.com/Zilliqa/gozilliqa-sdk/account"
	"github.com/Zilliqa/gozilliqa-sdk/keytools"
	provider2 "github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/Zilliqa/gozilliqa-sdk/util"
)

const host = "https://dev-api.zilliqa.com"
const chainID = 333
const msgVersion = 1
const privateKey = "c4e0d95cc91d4cd72c68c9cf58eec49eb9e5c2cbc63fec61e7bef5e7555de0f0"

func deploy(contractURI string, init string, privateKey string) {
	publickKey := keytools.GetPublicKeyFromPrivateKey(util.DecodeHex(privateKey), true)
	address := keytools.GetAddressFromPublic(publickKey)
	pubkey := util.EncodeHex(publickKey)
	provider := provider2.NewProvider(host)

	wallet := account.NewWallet()
	wallet.AddByPrivateKey(privateKey)

	code, _ := ioutil.ReadFile(contractURI)

}

func main() {

	deploy("./HelloWorld.scilla")
}
