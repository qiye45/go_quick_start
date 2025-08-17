package main

import "bytes"

func main() {
	strSlices := []string{"h", "e", "l", "l", "o"}

	var bf bytes.Buffer
	for _, str := range strSlices {
		bf.WriteString(str)
	}
	print(bf.String())
}
