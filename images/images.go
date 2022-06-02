package images

import (
	"os"
)

func SaveImage(name string, data []byte) bool {
	err := os.WriteFile("storage/"+name, data, 0644)
	if err != nil {
		return false
	}
	return true
}
