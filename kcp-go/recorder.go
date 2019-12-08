package kcp

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
)

var nOut int32 = 0

func getRecordDir(name string) string {
	ex, err := os.Executable()

	if err != nil {
		panic(err)
	}

	recordPath := ex + "_record/" + name;

	if _, err := os.Stat(recordPath); os.IsNotExist(err) {
		_ = os.MkdirAll(recordPath, os.ModePerm)
	}

	return recordPath
}

func recordOut(p []byte, addr net.Addr) {
	i := atomic.AddInt32(&nOut, 1)
	filename := fmt.Sprintf("%05d_%v", i, addr.String())

	filename = strings.ReplaceAll(filename, ":", "-")
	recordFile := filepath.Join(getRecordDir("out"), filename)

	err := ioutil.WriteFile(recordFile, p, os.ModePerm)

	fmt.Println(recordFile)

	if err != nil {
		fmt.Println(err.Error())
	}
}