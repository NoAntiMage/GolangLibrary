package main

import ("fmt")

// type byte = uint8
// type rune = int32

func main() {
	s := []byte("赵钱孙李周吴郑王")
	text := string(s)
	textRune := []rune(text)
	textLen := len(textRune)
	for i := 0; i < textLen-1; i++ {
		fmt.Printf("next_char: %s\n", string(textRune[i:i+1]))
	}	
}