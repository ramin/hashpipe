package main

import (
	"bufio"
	"crypto/md5"
	"encoding/csv"
	"flag"
	"fmt"
	"encoding/hex"
	"io"
	"os"
	"strings"
	"strconv"
)

var (
	input *string
	output *string
	encryptedFields *string
	encryptionScheme *string
	fieldMap []int
)

func DataReader(input *string) *bufio.Reader {
	if *input == "stdin" {
		return bufio.NewReader(os.Stdin)
	} else {
		source, err := os.Open(*input)
		defer source.Close()

		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		return bufio.NewReader(source)
	}
}

func DataWriter(output *string) *csv.Writer {
	return csv.NewWriter(os.Stdout)
}

func BuildOffsetMap(fields string) []int {
	splitFields := strings.Split(*encryptedFields, ",")

	finalMap := []int{}

	for _, field := range splitFields {
		if field == "" {
			continue
		}

		num, err := strconv.Atoi(field)

		if err != nil {
			fmt.Println(err)
			continue
		}

		finalMap = append(finalMap, num)
	}

	return finalMap
}

func main() {
	input = flag.String("input", "stdin", "input file")
	output = flag.String("output", "stdout", "output file")
	encryptionScheme = flag.String("scheme", "sha1", "sha1,md5...")
	encryptedFields = flag.String("fields", "", "eg: 0,3,8")
	flag.Parse()

	if *encryptedFields == "" {
		fmt.Println("please specify fields like 0,2,5,6 ")
		return
	}

	fieldMap := BuildOffsetMap(*encryptedFields)
	hashBuilder := md5.New()

	if len(fieldMap) == 0 {
		fmt.Println("no valid fields, try 0 indexed, 0,1,2")
		return
	}

	reader := DataReader(input)
	data := csv.NewReader(reader)
	writer := DataWriter(output)

	for {
		line, err := data.Read()

		for _, offset := range fieldMap {
			value := []byte(line[offset])
			hash := hashBuilder.Sum(value)
			line[offset] = hex.EncodeToString(hash)
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			break
		}

		writer.Write(line)
		writer.Flush()
	}
}
