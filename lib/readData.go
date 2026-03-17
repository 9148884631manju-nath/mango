package lib

import (
	"os"
)

func ReadData(file string ) (string) {
	data, err := os.ReadFile(file)
	if err!=nil {
		return "File Not Found"
	}
	return string(data)
}