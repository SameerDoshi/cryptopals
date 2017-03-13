package main
import (
	"encoding/hex";
	"encoding/base64";
	"bufio";
	"bytes";
	"os";
	"fmt";
	"log"
)
func main() {
	/* Get Input string from command line
		Decode, then encode to base64	
	*/
	var s string = os.Args[1]
	fmt.Println("Input hex:", s)
	var decoded []byte
	var err error 
	decoded, err= hex.DecodeString(s) 
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Decoded:",decoded)
	var encoded bytes.Buffer 
	wo:=bufio.NewWriter(&encoded)
	encoder := base64.NewEncoder(base64.StdEncoding, wo)
	encoder.Write(decoded)
	encoder.Close()
	wo.Flush()
	fmt.Println("Encoded:",encoded.String())

}
