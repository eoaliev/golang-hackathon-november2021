package utils

import (
    "fmt"
    "io/ioutil"
    "os"
)

func ReadBytesOfJsonFile(relativePath string) ([]byte, error) {
    wd, err := os.Getwd()
    if err != nil {
        return nil, err
    }

    file, err := os.Open(fmt.Sprintf("%v"+relativePath, wd))
    if err != nil {
        return nil, err
    }

    defer file.Close()

    byteFile, err := ioutil.ReadAll(file)
    if err != nil {
        return nil, err
    }

    return byteFile, nil;
}

func WriteBytesToJsonFile(relativePath string, byteFile []byte) (error) {
    wd, err := os.Getwd()
    if err != nil {
        return err
    }

    err = ioutil.WriteFile(fmt.Sprintf("%v"+relativePath, wd), byteFile, 0644)
    if err != nil {
        return err
    }

    return nil
}
