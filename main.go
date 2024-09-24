package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	FileEncryption "github.com/0x00f00bar/tiked_FileEncryption"
)

func main() {
	encryptFile := flag.String("e", "", "Encrypt the file at given path")
	decryptFile := flag.String("d", "", "Decrypt the file at given path")
	key := flag.String("k", "", "Key to decrypt an encrypted file")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "\nEncrypt/Decrypt a file securly using AES\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [-d | -e] <filepath> [-k] <key>\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	var err error

	switch {
	case *encryptFile != "":
		err = encrypt(encryptFile, flag.Args()...)
	case *decryptFile != "":
		err = decrypt(decryptFile, key)
	default:
		err = fmt.Errorf("error: invalid operation")
	}

	if err != nil {
		fmt.Println(err)
		flag.Usage()
		os.Exit(1)
	}
}

func encrypt(filePath *string, args ...string) error {
	//args contains additional arguments to the cmd
	if len(args) < 1 {
		fileAbsPath, err := checkFilePath(filePath)
		if err != nil {
			return err
		}

		key := make([]byte, 32)
		_, err = rand.Read(key)
		if err != nil {
			return err
		}
		fmt.Println("Encrypting file...", fileAbsPath)
		FileEncryption.InitializeBlock(key)
		outFile, err := FileEncryption.Encrypter(fileAbsPath)
		if err != nil {
			return err
		}
		fmt.Printf("Done. File saved as: %s\n", outFile)
		fmt.Printf("Use key %s to decrypt!\n", base64.StdEncoding.EncodeToString(key))
		return nil
	}

	return fmt.Errorf("can encrypt only one file at a time")
}

func decrypt(filepath *string, keyStr *string) error {
	if *keyStr == "" {
		return fmt.Errorf("error: no key provided")
	}

	key, err := base64.StdEncoding.DecodeString(*keyStr)
	if err != nil {
		return err
	}
	fileAbsPath, err := checkFilePath(filepath)
	if err != nil {
		return err
	}
	fmt.Println("Decrypting file...", fileAbsPath)
	FileEncryption.InitializeBlock(key)
	FileEncryption.Decrypter(fileAbsPath)
	fmt.Println("Done!")
	return nil
}

func checkFilePath(filePath *string) (string, error) {
	var path string
	if file, err := os.Stat(*filePath); err == nil {
		if file.IsDir() {
			return "", fmt.Errorf("error: path provided is a directory; file is required")
		}
		path, err = filepath.Abs(*filePath)
		if err != nil {
			return "", err
		}
	} else if errors.Is(err, os.ErrNotExist) {
		return "", err
	}
	return path, nil
}
