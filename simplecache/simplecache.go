package simplecache

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

type Config struct {
	StorePath string
}

var config Config

func SetConfig(conf Config) {
	config = conf
}

func GetConfig() Config {
	return config
}

func Store(key string, value []byte) error {

	filePath := getFilePathFromKey(key)
	return ioutil.WriteFile(filePath, value, 0644)
}

func Read(key string) ([]byte, error) {

	filePath := getFilePathFromKey(key)
	return ioutil.ReadFile(filePath)
}

func Remove(key string) error {

	filePath := getFilePathFromKey(key)
	return os.Remove(filePath)
}

func getFilePathFromKey(key string) string {

	hash := md5.Sum([]byte(key))
	return filepath.Join(config.StorePath, hex.EncodeToString(hash[:]))
}

func RemoveAll() error {

	files, err := filepath.Glob(filepath.Join(config.StorePath, "*"))
	if err != nil {
		return err
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func StoreUint64(key string, value uint64) error {
	return Store(key, []byte(fmt.Sprintf("%d", value)))
}

func ReadUint64(key string) (uint64, error) {

	bval, err := Read(key)
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(string(bval), 10, 64)
}
