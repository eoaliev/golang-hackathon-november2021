package datajson

import (
    "encoding/json"
    "fmt"
    "github.com/eoaliev/golang-hackathon-november2021/utils"
    "os"
)

type Transaction struct {
    UserID int `json:"user_id"`
    Timestamp string `json:"timestamp"`
    Category string `json:"category"`
    Card string `json:"card"`
    Amount int `json:"amount"`
}

var transactions []Transaction

func init() {
    byteFile, err := utils.ReadBytesOfJsonFile("/datajson/transactions.json")
    if err != nil {
        er(err)
    }

    err = json.Unmarshal(byteFile, &transactions)
    if err != nil {
        er(err)
    }
}

func GetTransactions() ([]Transaction) {
    return transactions
}

func er(msg interface{}) {
    fmt.Println("Error:", msg)
    os.Exit(1)
}
