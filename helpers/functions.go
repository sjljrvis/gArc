package helpers

import (
	"bytes"
	"errors"
	"net"
	"os"
)

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

func CheckDir(dirname string) error {
	println(dirname)
	_, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return nil
	} else {
		return errors.New(dirname + " Already Exists")
	}
}
