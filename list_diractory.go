package main
import (
	"fmt"
	"io/ioutil"
	"strings"
)

const WORKDIR = "/Users/wujimaster/GoProjects/hello_go/static"
const TARGETDIR = WORKDIR + "/2020-11_archive"

func main() {
	files, err := ioutil.ReadDir(WORKDIR)
	if err != nil {
		fmt.Println(err)
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		if name := f.Name(); strings.Contains(name, "wms") {
			fmt.Println(name)
		}
	}
}