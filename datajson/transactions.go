package datajson

import (
    "encoding/json"
    "fmt"
    "github.com/eoaliev/golang-hackathon-november2021/utils"
    "os"
)

type Transaction struct {
    UserID int `json:"user_id"`
    // Закоментировал потому что в тестовых файлах на 1M и на 10M тип int а в предыдущих string
    // Timestamp string `json:"timestamp"`
    Category string `json:"category"`
    Card string `json:"card"`
    Amount int `json:"amount"`
}

var transactions []Transaction

func init() {
    utils.ActualizeTimer("datajson")

    byteFile, err := utils.ReadBytesOfJsonFile("/staticfiles/transactions.json")
    if err != nil {
        er(err)
    }

    utils.PrintDuration("datajson", "Read json for")

    utils.ActualizeTimer("datajson")

    err = json.Unmarshal(byteFile, &transactions)
    if err != nil {
        er(err)
    }

    utils.PrintDuration("datajson", "Parse json for")

    utils.StopTimer("datajson")
}

func er(msg interface{}) {
    fmt.Println("Error:", msg)
    os.Exit(1)
}

func GetTransactions() ([]Transaction) {
    return transactions
}
