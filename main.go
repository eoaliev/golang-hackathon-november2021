package main

import (
    "fmt"
    "github.com/eoaliev/golang-hackathon-november2021/datajson"
    "github.com/eoaliev/golang-hackathon-november2021/reports"
    "github.com/eoaliev/golang-hackathon-november2021/utils"
)

func main() {
	utils.ActualizeTimer()

    transactions := datajson.GetTransactions()

    utils.PrintDuration("Parse json for")

    utils.ActualizeTimer()

    report := reports.GenerateUsersExpensesReport(transactions)

    utils.PrintDuration("Generate report for")

    utils.ActualizeTimer()

    err := reports.WriteReportToJsonFile(report, reports.GetUsersExpensesReportFileName())
    if (err != nil) {
        fmt.Printf("Report failure generated: %+v\n", err)
        return
    }

    utils.PrintDuration("Write report to json file for")
}
