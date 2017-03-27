package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
)

func main() {
	/* Get Input string from command line
	Decode, then encode to base64
	*/
	// Convert a hex to base 64
	// input: 49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d
	// should ouput:
	// SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
	s1c1in := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	o := tobase64(decodehex(s1c1in))
	fmt.Println("*****Challenge 1 answer:")
	fmt.Println(string(o))
	fmt.Println("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t")

	s1c21 := "1c0111001f010100061a024b53535009181c"
	s1c22 := "686974207468652062756c6c277320657965"

	//var o1, o2 []byte
	o1 := decodehex(s1c21)
	fmt.Println("O1: ", string(o1))
	o2 := decodehex(s1c22)
	dst := fastXORBytes(o1, o2)
	fmt.Println("***Challenge 2")
	fmt.Println("746865206b696420646f6e277420706c6179")
	n := hex.EncodeToString(dst)
	fmt.Println(n)
	fmt.Println(string(dst))

	fmt.Println("****Challenge 3")
	s1c3 := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	loopthroughalpha(s1c3)
}

func decodehex(s string) []byte {
	decoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	return (decoded)
}

func tobase64(in []byte) []byte {

	var encoded bytes.Buffer
	wo := bufio.NewWriter(&encoded)
	encoder := base64.NewEncoder(base64.StdEncoding, wo)
	encoder.Write(in)
	encoder.Close()
	wo.Flush()

	//	fmt.Println("endcoded len", encoded.Len())
	dst := encoded.Next(encoded.Len())
	//	fmt.Println("DST Len", len(dst))
	return (dst)
}

func fastXORBytes(a, b []byte) []byte {
	n := len(a)

	if len(b) < n {
		n = len(b)
	}

	dst := make([]byte, n)

	for i := 0; i < n; i++ {
		dst[i] = a[i] ^ b[i]
	}

	return dst
}

func loopthroughalpha(in string) {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	n := len(in)
	//const alpha = "abcdefghijklmnopqrstuvwxyz"
	inbyte := decodehex(in)
	var tmp []byte
	for i := 0; i < len(letterRunes); i++ {
		s := strings.Repeat(string(letterRunes[i]), n)
		hi := []byte(s)
		tmp = hex.EncodeToString(hi)
		fastXORBytes(tmp, inbyte)
		fmt.Println()

	}
}
