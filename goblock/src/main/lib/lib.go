package lib

import(
	"encoding/hex"
	"crypto/sha256"
	"time"
	"strconv"
	"os"
	"io/ioutil"
)

func GetShaString(name string) string {
	hash := sha256.Sum256([]byte(name))
	return hex.EncodeToString(hash[:])
}

func GetTime() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func MakeDir(fileLoc string, name string) error {
	return os.Mkdir(fileLoc+"/"+name, 0777)
}

func RmFile(fileLoc string, id string, name string) error {
	return os.Remove(fileLoc + "/" + id + "/" + name)
}

func ReadFile(fileLoc string, id string, name string) ([]byte, error) {
	return ioutil.ReadFile(fileLoc + "/" + id + "/" + name)
}

func WriteFile(fileLoc string, id string, name string, data []byte) error {
	return ioutil.WriteFile(fileLoc + "/" + id + "/" + name, data, 0444)
}

func IsFileExist (fileLoc string, fileName string) bool {
	_, err := os.Stat(fileLoc + "/" + fileName)
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}