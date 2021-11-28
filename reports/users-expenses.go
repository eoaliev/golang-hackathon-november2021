package reports

import (
    "github.com/eoaliev/golang-hackathon-november2021/datajson"
)

type UserExpenseCategory struct {
    Name string `json:"name"`
    Count int `json:"count"`
    Sum int `json:"sum"`
}

type UserExpense struct {
    UserID int `json:"user_id"`
    Sum int `json:"sum"`
    Categories []UserExpenseCategory `json:"categories"`
}

var usersExpensesMap map[int]map[string]UserExpenseCategory

func GenerateUsersExpensesReport(transactions []datajson.Transaction) ([]UserExpense) {
    InitUsersExpenseMap(transactions);

    report := make([]UserExpense, 0, len(usersExpensesMap))
    for userId, userExpensesMap := range usersExpensesMap {
        var userExpense UserExpense
        userExpense.UserID = userId
        userExpense.Sum = 0
        userExpense.Categories = make([]UserExpenseCategory, 0, len(userExpensesMap))

        for _, userExpenseCategory := range userExpensesMap {
            userExpense.Categories = append(userExpense.Categories, userExpenseCategory)
            userExpense.Sum += userExpenseCategory.Sum
        }

        report = append(report, userExpense)
    }

    return report
}

func InitUsersExpenseMap(transactions []datajson.Transaction) {
    usersExpensesMap = map[int]map[string]UserExpenseCategory{}
    for _, transaction := range transactions {
        if i, ok := usersExpensesMap[transaction.UserID][transaction.Category]; ok {
            i.Sum += transaction.Amount
            i.Count++
            usersExpensesMap[transaction.UserID][transaction.Category] = i
            continue
        }

        var i UserExpenseCategory
        i.Name = transaction.Category
        i.Count = 1
        i.Sum = transaction.Amount

        if _, ok := usersExpensesMap[transaction.UserID]; !ok {
            usersExpensesMap[transaction.UserID] = map[string]UserExpenseCategory{}
            continue
        }

        usersExpensesMap[transaction.UserID][transaction.Category] = i
    }
}

func GetUsersExpensesReportFileName() (string) {
    return "users-expenses.json";
}
