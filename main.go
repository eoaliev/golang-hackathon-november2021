package main

import (
    "fmt"
    "github.com/eoaliev/golang-hackathon-november2021/datajson"
    "github.com/eoaliev/golang-hackathon-november2021/reports"
    "github.com/eoaliev/golang-hackathon-november2021/utils"
)

func main() {
	utils.ActualizeTimer("main")

    transactions := datajson.GetTransactions()

    utils.PrintDuration("main", "Parse json for")

    utils.ActualizeTimer("main")

    report := reports.TransactionsToUsersExpensesReport(transactions)

    utils.PrintDuration("main", "Generate report for")

    utils.ActualizeTimer("main")

    err := reports.WriteReportToJsonFile(report, reports.GetUsersExpensesReportFileName())
    if (err != nil) {
        fmt.Printf("Report failure generated: %+v\n", err)
        return
    }

    utils.PrintDuration("main", "Write report to json file for")

    utils.StopTimer("main")
}
