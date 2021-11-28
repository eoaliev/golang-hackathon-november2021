package main

import (
    "errors"
    "fmt"
    "github.com/eoaliev/golang-hackathon-november2021/datajson"
    "github.com/eoaliev/golang-hackathon-november2021/reports"
    "github.com/eoaliev/golang-hackathon-november2021/utils"
)

func main() {
    utils.ActualizeTimer("main")

    report := reports.TransactionsToUsersExpensesReport(datajson.GetTransactions())

    utils.PrintDuration("main", "Generate report for")

    var level int
    fmt.Println("Enter level (1, 2, 3, 404):")
    fmt.Scanf("%d", &level)

    err := outputReportByLevel(level, report)
    if (err != nil) {
        fmt.Printf("Report failure generated: %+v\n", err)
        return
    }

    utils.StopTimer("main")
}

func outputReportByLevel(level int, report []reports.UserExpense) (error) {
    switch level {
    case 1:
        return errors.New("Not implemented yet!")

    case 2:
        utils.ActualizeTimer("main")
        err := reports.WriteReportToJsonFile(report, reports.GetUsersExpensesReportFileName())

        utils.PrintDuration("main", "Write report to json file for")
        return err

    case 3:
        return errors.New("Not implemented yet!")

    case 404:
        return errors.New("Not implemented yet!")
    }

    return errors.New("Incorrect level use (1, 2, 3, 404)")
}
