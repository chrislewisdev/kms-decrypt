package main

import (
	"flag"
	"fmt"
	"encoding/base64"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

func main() {
	//Set up our command-line arguments
	region := flag.String("region", "ap-southeast-2", "The AWS region containing the key that the value was encrypted with")
	encodedEncryptedValue := flag.String("value", "", "The value to decrypt")
	flag.Parse()

	if *encodedEncryptedValue == "" {
		fmt.Println("Error: Argument -value is required")
		return
	}

	sessionConfig := aws.Config{ Region: aws.String(*region) }
	apiSession, _ := session.NewSession(&sessionConfig)
	kmsService := kms.New(apiSession)

	//Encrypted KMS strings are base-64 encoded, so we need to decode it before decrypting it
	decodedEncryptedValue, _ := base64.StdEncoding.DecodeString(*encodedEncryptedValue)

	result, err := kmsService.Decrypt(&kms.DecryptInput{ CiphertextBlob: decodedEncryptedValue })
	if err != nil {
		panic(err)
	}

	fmt.Println(string(result.Plaintext))
}
