package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/wjessop/passthroughdigest"
)

func main() {
	input := bytes.NewReader([]byte("Digest this!"))

	output, err := os.Create("/tmp/outputfile")
	if err != nil {
		panic(err)
	}

	defer output.Close()

	digestWriter := passthroughdigest.NewPassthroughDigest(output)
	bytesCopied, err := io.Copy(digestWriter, input)

	hashString := hex.EncodeToString(digestWriter.Digest())

	fmt.Printf("wrote %d bytes, digest of data was %s\n", bytesCopied, hashString)

	// Confirm the file with a local MD5 command:
	// $ md5 /tmp/outputfile
	// MD5 (/tmp/outputfile) = f7d9dd511e4d7c60b72cfb897e9e17c3
}
