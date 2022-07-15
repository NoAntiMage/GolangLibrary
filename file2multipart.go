package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Conversion: singlepart from/to multipart. split&merge")

	num := fileToChunk()
	chunkToFile(num)
}

func fileToChunk() (num int) {
	targetFile := "./source.zip"

	f, err := os.OpenFile(targetFile, os.O_RDONLY, os.ModePerm)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	fileInfo, err := f.Stat()
	if err != nil {
		fmt.Println(err)
	}

	var fileSize int64 = fileInfo.Size()

	const chunkSize = 1 * (1 << 20)
	totalPartsNum := int(math.Ceil(float64(fileSize) / float64(chunkSize)))

	fmt.Printf("Split to %d pieces.\n", totalPartsNum)

	for i := int(0); i < totalPartsNum; i++ {
		partSize := int(math.Min(chunkSize, float64(fileSize-int64(i*chunkSize))))
		partBuffer := make([]byte, partSize)

		f.Read(partBuffer)

		fileName := "obj_" + strconv.Itoa((i + 1))

		chunk, err := os.Create(fileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		chunk.Write(partBuffer)
		chunk.Close()

		fmt.Println("split to :", fileName)
	}
	return totalPartsNum
}

func chunkToFile(num int) {
	fout, err := os.OpenFile("restore.zip", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	defer fout.Close()
	if err != nil {
		fmt.Println(err)
	}

	for i := int(0); i < num; i++ {
		fileName := "obj_" + strconv.Itoa((i + 1))
		fmt.Println(fileName)
		f, _ := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
		b, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println(err)
			return
		}
		fout.Write(b)
		f.Close()
	}
}
