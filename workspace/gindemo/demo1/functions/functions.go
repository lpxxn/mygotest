package funcs

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

// RevStr ...
func RevStr(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}

// Struct2Map ...
func Struct2Map(obj interface{}) (data map[string]interface{}) {
	data = make(map[string]interface{})
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	for i := 0; i < objT.NumField(); i++ {
		data[objT.Field(i).Name] = objV.Field(i).Interface()
	}
	return
}
func AppFilePath(workPath, PROJDir, localPath string, fileName string, check_exit bool) string {
	lastDir := SubStr(
		workPath,
		strings.LastIndex(workPath, "/")+1,
		len(workPath)-strings.LastIndex(workPath, "/"),
	)

	var configPath string

	if lastDir == PROJDir {
		configPath = JoinFilePath(workPath, localPath, fileName)
	} else if lastDir == "bin" {
		parentDir := GetParentDirectory(workPath)
		configPath = JoinFilePath(parentDir, "src", PROJDir, localPath, fileName)
		if check_exit {
			fileStatus, err := FileExists(configPath)
			if nil != err || !fileStatus {
				configPath = JoinFilePath(parentDir, localPath, fileName)
			}
		}
	}
	return configPath
}

// Slice2Map ...
func Slice2Map(obj []interface{}) (data map[int]interface{}) {
	data = make(map[int]interface{})
	for k, v := range obj {
		data[k] = v
	}
	return
}

// RandCode ...
func RandCode(ty string, l int) string {
	model := map[string]string{
		"d":   "0123456789",
		"c":   "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"s":   "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"mix": "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^",
	}

	return GetRandomString(model[ty], l)
}

// GetRandomString ...
func GetRandomString(str string, l int) string {
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// GetCurrentDirectory ...
func GetCurrentDirectory() (string, error) {
	var dir string
	var err error
	dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	dir = strings.Replace(dir, "\\", "/", -1)
	if path.Base(dir) == "exe" {
		dir, err = os.Getwd()
	}

	if err != nil {
		return "", err
	}
	return strings.Replace(dir, "\\", "/", -1), nil
}

// GetParentDirectory ...
func GetParentDirectory(directory string) string {
	return SubStr(directory, 0, strings.LastIndex(directory, "/"))
}

// SubStr ...
func SubStr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

// JoinFilePath ...
func JoinFilePath(directory ...string) string {
	var path string
	for _, v := range directory {
		path = filepath.Join(path, v)
	}
	return path
}

// FileExists ...
func FileExists(FilePath string) (bool, error) {
	if _, err := os.Stat(FilePath); os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// StrMD5 MD5
func StrMD5(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

// B2S ...
func B2S(bs []uint8) string {
	b := make([]byte, len(bs))
	for i, v := range bs {
		b[i] = byte(v)
	}
	return string(b)
}
