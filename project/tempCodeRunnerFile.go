package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	//"encoding/hex"
)

func main() {
	fmt.Println("BEGAN ...\n")
	num := 0
	for true {
		str := strconv.Itoa(num)
		h := sha256.New()
		h.Write([]byte(str))
		sh := hex.EncodeToString(h.Sum(nil))
		if sh[0] == '0' && sh[1] == '0' && sh[2] == '0' && sh[3] == '0' {
			fmt.Printf("NUMBER =%d", num)
			fmt.Printf("\n\n%x", h.Sum(nil))
			break
		}
		num += 1
	}
	fmt.Println("\n\nFINISHED ...")