Varint Eocoding
---

### Conception

---

1. The highest bit of each byte is the reserved bit. If it is 1, it indicates whether the next byte belongs to the current data. If it is 0, it is the last byte of the current data Look at the following code. Since the highest bit of a byte is the reserved bit, only the following 7 bits can save data in this byte.
2. Therefore, if x > 127, the data needs to be saved larger than one byte, so the highest bit of the current byte is 1. See buf [n] = 0x80 |
    0x80 indicates that the highest position of this byte is 1, and the following X & 0x7F is to obtain the lower 7-bit data of X, so the overall meaning of 0x80 | uint8 (X & 0x7F) is the highest bit of this byte is 1, which means this is not the last byte, and the last 7 is the official data! Note that x > > = 7 should be set before the next byte is operated
3. If x < = 127, then x can be represented by 7bits now. Then the highest bit does not need to be 1, and it's ok if it's directly 0! So the last one is buf [n] = uint8 (x)
    //
    //If the data is larger than one byte (127 is the maximum data of one byte), then continue, i.e. add 1 to the highest bit



### Example

---

```
	123456 int
					
    111 1000100 1000000 bin
      \      |      /
        \    |    /
          \  |  /
            \|/
            /|\
          /  |  \
        /    |    \
      /      |      \ 
11000000 110000100 00000111 varint
```



### Reference 

---

1. /go/src/encoding/binary/varint.go
2. https://developpaper.com/explain-the-principle-of-varint-coding-in-detail/
