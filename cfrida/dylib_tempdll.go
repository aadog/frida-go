// +build windows linux
// +build tempdll

package cfrida

import (
	"bytes"
	"fmt"
	"github.com/a97077088/libfridabinres"
	"hash/crc32"
	"io"
	"io/ioutil"
	"os"
)


func BytesToFile(destFileName string, input []byte) error {
	fi, err := os.Create(destFileName)
	if err != nil {
		return err
	}
	defer fi.Close()
	_, err = io.Copy(fi, bytes.NewReader(input))
	if err != nil {
		return nil
	}
	return nil
}

var (
	spStr = string(os.PathSeparator)
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func checkAndReleaseDLL() (bool, string) {
	crc:=crc32.ChecksumIEEE(libfridabinres.FridaBinRes)
	tempDLLDir := fmt.Sprintf("%s/libfrida/%x", os.TempDir(), crc)
	// create liblcl: $tempdir/libfrida/{crc32}/libfrida.{ext}
	if !fileExists(tempDLLDir) {
		if err := os.MkdirAll(tempDLLDir, 0775); err != nil {
			return false, ""
		}
	}
	tempDLLFileName := fmt.Sprintf("%s/%s", tempDLLDir, getDLLName())
	// test crc32
	if fileExists(tempDLLFileName) {
		bs, err := ioutil.ReadFile(tempDLLFileName)
		if err == nil {
			if crc32.ChecksumIEEE(bs) != crc {
				os.Remove(tempDLLFileName)
			}
		}
	}
	if !fileExists(tempDLLFileName) {
		if err := BytesToFile(tempDLLFileName, libfridabinres.FridaBinRes); err != nil {
			if os.Remove(tempDLLFileName) != nil {
				return false, ""
			}
		}
	}
	return true, tempDLLFileName
}