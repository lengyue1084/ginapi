package tool

import (
	"os"
	"time"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Date(f string) string {
	if f == "Y-m-d H:i:s" || f == "" {
		return time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	}
	if f == "Y-m-d" {
		return time.Unix(time.Now().Unix(), 0).Format("2006-01-02")
	}
	return ""
}
