package main

// package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main(){
	fmt.Println("go file operator")
}

func CreateDir(src string) error {
	return os.MkdirAll(src, os.ModePerm)
}

func LoadFile(src string) ([]byte, error) {
	fd, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	c, err := ioutil.ReadAll(fd)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func WriteFile(src string, content []byte) error {
	return ioutil.WriteFile(src, content, os.ModeAppend)
}

func WriteFileAppend(src string, content []byte) error {
	if IsFile(src) == false {
		return WriteFile(src, content)
	}
	c, err := LoadFile(src)
	if err != nil {
		return err
	}
	var s [][]byte
	s = [][]byte{c, content}
	sep := []byte("")
	var newContent []byte = bytes.Join(s, sep)
	return WriteFile(src, newContent)
}

func CutFile(src string, newSrc string) error {
	return os.Rename(src, newSrc)
}

func CopyFile(src string, dest string) (bool, error) {
	srcF, err := os.Open(src)
	if err != nil {
		return false, err
	}
	defer srcF.Close()
	desfF, err := os.Create(dest)
	if err != nil {
		return false, err
	}
	defer destF.Close()
	_, err = io.Copy(destF, srcF)
	if err != nil {
		return false, err
	}
	return true, err
}

func DeleteFile(src string) error {
	return os.RemoveAll(src)
}

func IsExist(src string) bool {
	_, err := os.Stat(src)
	return err == nil || os.IsExist(err)
}

func IsFile(src string) bool {
	info, err := os.Stat(src)
	return err == nil && !info.IsDir()
}

func IsFolder(src string) bool {
	info, err := os.Stat(src)
	return err == nil && info.IsDir()
}

func GetFileList(src string, filter string, isSrc bool) ([]string, error) {
	var fs []string
	dir, err := ioutil.ReadDir(src)
	if err != nil {
		return nil, err
	}
	var filters []string
	if filter != "" {
		filters = strings.Split(filter, "|")
	}
	for _, v := range dir {
		var appendSrc string
		if isSrc == true {
			appendSrc = src + GetPathSep() + v.Name()
		} else {
			appendSrc = v.Name()
		}
		if v.IsDir() == true {
			fs = append(fs, appendSrc)
			continue
		}
		if filter == "" {
			fs = append(fs, appendSrc)
			continue
		}
		names := strings.Split(v.Name(), ".")
		if len(names) == 1 {
			fs = append(fs, appendSrc)
			continue
		}
		t := names[len(names)-1]
		for _, filterValue := range filters {
			if t != filterValue {
				continue
			}
			fs = append(fs, appendSrc)
		}
	}

	return fs, nil
}

func getFileListCount(src string) (int, error) {
	dir, err := ioutil.ReadDir(src)
	if err != nil {
		return 0, err
	}
	var res int
	for range dir {
		res += 1
	}
	return res, nil
}

func GetPathSep() string {
	return string(os.PathSeparator)
}

func GetFileSize(src string) int64 {
	info, err := os.Stat(src)
	if err != nil {
		return 0
	}
	return info.Size()
}

func GetFileNames(src string) (map[string]string, error) {
	info, err := GetFileInfo(src)
	if err != nil {
		return nil, err
	}
	res := map[string]string{
		"name":     info.Name(),
		"type":     "",
		"onlyName": info.Name(),
	}
	if res["name"] == "" {
		return res, nil
	}
	names := strings.Split(res["name"], ".")
	if len(names) < 2 {
		return res, nil
	}
	for i := range names {
		if i != 0 && i < len(names)-1 {
			res["onlyName"] = res["onlyName"] + "." + names[i]
		}
	}
	return res, nil
}

func GetFileInfo(src string) (os.FileInfo, error) {
	return os.Stat(src)
}

func GetFileSha1(src string) (string, error) {
	content, err := LoadFile(src)
	if err != nil {
		return "", err
	}
	if content != nil {
		sha := sha1.New()
		sha.Write(content)
		res := sha.Sum(nil)
		return hex.EncodeToString(res), nil
	}
	return "", nil
}

//eg : Return and create the path ,"[src]/201611/"
//eg : Return and create the path ,"[src]/201611/2016110102-03[appendFileType]"
func GetTimeDirSrc(src string, appendFileType string) (string, error) {
	t := time.Now()
	sep := GetPathSep()
	newSrc := src + sep + t.Format("200601")
	err = CreateDir(newSrc)
	newSrc = newSrc + sep
	if appendFileType != "" {
		newSrc = newSrc + t.Format("20060102-03") + appendFileType
	}
	return newSrc, err
}
