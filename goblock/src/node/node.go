package node

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type Account struct {
	AccId      string
	AccName    string
	AccBalance uint
}

func CreateAccount(accName string) Account {
	timeNow := strconv.FormatInt(time.Now().Unix(), 10)
	shaName := sha256.Sum256([]byte(accName + timeNow))

	return Account{
		AccId:      hex.EncodeToString(shaName[:]),
		AccName:    accName,
		AccBalance: 0,
	}
}

func StoreAccount(fileLoc string, acc Account) {
	accJson, err := json.MarshalIndent(acc, "", " ")

	err = ioutil.WriteFile(fileLoc+"/"+acc.AccName, accJson, 0444)

	if err != nil {
		fmt.Println("User already exists")
	}
}

func RetrieveAccount(fileLoc string, accName string) Account {
	accDetail, err := ioutil.ReadFile(fileLoc + "/" + accName)

	acc := Account{}
	err = json.Unmarshal(accDetail, &acc)
	check(err)

	return acc
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) > 1 {
		StoreAccount("../store/node", CreateAccount(os.Args[1]))
	}
}
