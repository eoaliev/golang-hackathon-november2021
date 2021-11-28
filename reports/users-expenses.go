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

func TransactionsToUsersExpensesReport(transactions []datajson.Transaction) ([]UserExpense) {
    usersExpensesMap := transactionsToUsersExpenseMap(transactions);

    report := make([]UserExpense, 0, len(usersExpensesMap))
    for userId, userExpensesMap := range usersExpensesMap {
        userExpense := getUserExpenseByUserIdAndMap(userId, userExpensesMap)

        report = append(report, userExpense)
    }

    return report
}

func transactionsToUsersExpenseMap(transactions []datajson.Transaction) (map[int]map[string]UserExpenseCategory) {
    var userExpenseCategory UserExpenseCategory

    usersExpensesMap := map[int]map[string]UserExpenseCategory{}
    for _, transaction := range transactions {
        if _, ok := usersExpensesMap[transaction.UserID]; !ok {
            usersExpensesMap[transaction.UserID] = map[string]UserExpenseCategory{}
        }

        userExpenseCategory.Name = transaction.Category
        userExpenseCategory.Count = 1
        userExpenseCategory.Sum = transaction.Amount

        if i, ok := usersExpensesMap[transaction.UserID][transaction.Category]; ok {
            userExpenseCategory.Count += i.Count
            userExpenseCategory.Sum += i.Sum
        }

        usersExpensesMap[transaction.UserID][transaction.Category] = userExpenseCategory
    }

    return usersExpensesMap;
}

func getUserExpenseByUserIdAndMap(userId int, userExpensesMap map[string]UserExpenseCategory) (UserExpense) {
    var userExpense UserExpense
    userExpense.UserID = userId
    userExpense.Categories, userExpense.Sum = usersExpensesMapToCategoriesAndSum(userExpensesMap)

    return userExpense
}

func usersExpensesMapToCategoriesAndSum(userExpensesMap map[string]UserExpenseCategory) ([]UserExpenseCategory, int) {
    categories := make([]UserExpenseCategory, 0, len(userExpensesMap))
    sum := 0

    for _, userExpenseCategory := range userExpensesMap {
        categories = append(categories, userExpenseCategory)
        sum += userExpenseCategory.Sum
    }

    return categories, sum
}

func GetUsersExpensesReportFileName() (string) {
    return "users-expenses.json";
}
