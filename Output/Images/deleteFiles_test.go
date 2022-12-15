package Images

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestDelete(t *testing.T) {
	items, _ := ioutil.ReadDir(".")
	for _, item := range items {
		if !item.IsDir() {
			if strings.HasSuffix(item.Name(), ".png") {
				os.RemoveAll(item.Name())
			}
		}
	}
}
