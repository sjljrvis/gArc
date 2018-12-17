package helpers

import (
	"bytes"
	"errors"
	"net"
	"os"
	"path/filepath"
)

// GetMacAddress of current node
func GetMacAddress() (addr string) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				// Don't use random as we have a real address
				addr = i.HardwareAddr.String()
				break
			}
		}
	}
	return
}

// CheckDir  *for checking directory
func CheckDir(dirname string) error {
	_, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return nil
	}
	return errors.New(dirname + " Already Exists")
}

// ListFiles  *for listing  files in directory
func ListFiles(dirName string) ([]string, error) {
	files, err := filepath.Glob(os.Getenv("HOME") + "/gArch/*")
	if err != nil {
		return nil, err
	}
	return files, nil
}
