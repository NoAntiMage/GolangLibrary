package main

import (
	"fmt"
	"strings"
)

// 1. 每行代表一条记录，字段之间以逗号，分隔
// 2. 字段包含逗号时,以双引号包围该字段
// 3. 字段包含双引号时,以双引号包围该字段,双引号由一个变两个
// 结果输出，解析记录，字段间以 \t 分隔

func main() {
	a := `John,45,"足球,摄影",New York`
	b := `Carter Job, 33."""健身"",远足","河北,石家庄"`
	c := `Steve,33"大屏幕164""","DC""Home"""`
	d := `"Jul,y",33,Football,Canada`
	parseLine(a)
	parseLine(b)
	parseLine(c)
	parseLine(d)

}

func parseLine(line string) string {
	//flag for quote stack
	var stackFlag bool = false
	var quoteCount int = 0
	var word string
	var words = make([]string, 0)

	inputs := []rune(line)
	// convert double quotes to #
	for i := 0; i < len(inputs); i++ {
		if inputs[i] == '"' && quoteCount == 0 {
			x := i
		count:
			for {
				// limit x
				if x >= len(inputs) {
					break count
				}

				if inputs[x] == '"' {
					quoteCount++
				} else {
					break count
				}

				x++
			}
		}

		if quoteCount != 0 {
			if quoteCount%2 == 1 && stackFlag == false { // start stack
				for j := 1; j < quoteCount; j++ {
					inputs[i+j] = '#'
				}
				stackFlag = true
			} else if quoteCount%2 == 1 && stackFlag == true { // end stack
				for j := 0; j < quoteCount-1; j++ {
					inputs[i+j] = '#'
				}
				stackFlag = false
				i += quoteCount - 1
			} else { //double quote
				for j := 0; j < quoteCount; j++ {
					inputs[i+j] = '#'
				}
			}
			quoteCount = 0
		}
	}

	for i := 0; i < len(inputs); i++ {
		if inputs[i] == '"' {
			stackFlag = !stackFlag
			continue
		} else if inputs[i] == ',' {
			if stackFlag == false {
				words = append(words, word)
				word = ""
				continue
			}
		}

		word = word + string(inputs[i])
	}

	// the last word
	words = append(words, word)
	for i := 0; i < len(words); i++ {
		if strings.Contains(words[i], "##") {
			words[i] = strings.Replace(words[i], "##", `"`, -1)
		}
	}

	// result sep \t
	ret := strings.Join(words, "\t")
	fmt.Println(ret)
	return ret
}
