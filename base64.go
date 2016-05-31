package main

import (
	"io"
	"os"
)

var base64 = map[byte]byte{0: 65, 1: 66, 2: 67, 3: 68, 4: 69, 5: 70, 6: 71, 7: 72,
	8: 73, 9: 74, 10: 75, 11: 76, 12: 77, 13: 78, 14: 79, 15: 80, 16: 81, 17: 82, 18: 83,
	19: 84, 20: 85, 21: 86, 22: 87, 23: 88, 24: 89, 25: 90, 26: 97, 27: 98, 28: 99,
	29: 100, 30: 101, 31: 102, 32: 103, 33: 104, 34: 105, 35: 106, 36: 107, 37: 108,
	38: 109, 39: 110, 40: 111, 41: 112, 42: 113, 43: 114, 44: 115, 45: 116, 46: 117,
	47: 118, 48: 119, 49: 120, 50: 121, 51: 122, 52: 48, 53: 49, 54: 50, 55: 51, 56: 52,
	57: 53, 58: 54, 59: 55, 60: 56, 61: 57, 62: 43, 63: 47}

func encodetriplet(buf []byte) []byte {
	retbuf := make([]byte, 4)
	retbuf[0] = base64[buf[0]>>2]
	switch len(buf) {
	case 3:
		retbuf[1] = base64[buf[0]<<6>>2+buf[1]>>4]
		retbuf[2] = base64[buf[1]<<4>>2+buf[2]>>6]
		retbuf[3] = base64[buf[2]<<2>>2]
	case 2:
		retbuf[1] = base64[buf[0]<<6>>2+buf[1]>>4]
		retbuf[2] = base64[buf[1]<<4>>2]
		retbuf[3] = '='
	case 1:
		retbuf[1] = base64[(buf[0]<<6)>>2]
		retbuf[2] = '='
		retbuf[3] = '='
	}
	return retbuf

}

func main() {
	buf := make([]byte, 1024*3*10)

	for n, err := os.Stdin.Read(buf); err != io.EOF; n, err = os.Stdin.Read(buf) {
		for i := 0; i < n; i += 3 {
			if i%19 == 0 && i != 0 {
				os.Stdout.WriteString("\n")
			}
			j := i + 3
			if n < j {
				j = n
			}
			os.Stdout.Write(encodetriplet(buf[i:j]))
		}
	}
	os.Stdout.WriteString("\n")
}
