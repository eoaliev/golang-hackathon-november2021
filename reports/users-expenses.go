package reports

import (
    "github.com/eoaliev/golang-hackathon-november2021/datajson"
)

type UserExpenseCategory struct {
    Name string `json:"name"`
    Count int `json:"count"`
    Sum int `json:"sum"`
}

type UserCategoryIndex struct {
    UserID int
    Category string
}

type UserExpense struct {
    UserID int `json:"user_id"`
    Sum int `json:"sum"`
    Categories []UserExpenseCategory `json:"categories"`
}

func TransactionsToUsersExpensesReport(transactions []datajson.Transaction) ([]UserExpense) {
    userCategoryExpenseMap := transactionsToUserCategoryExpenseMap(transactions);

    usersExpenseCategoriesMap, usersExpenseSumMap := getUsersExpenseCategoriesAndSum(userCategoryExpenseMap);

    var userExpense UserExpense

    report := make([]UserExpense, 0, len(usersExpenseCategoriesMap))
    for userId, categories := range usersExpenseCategoriesMap {
        userExpense.UserID = userId
        userExpense.Categories = categories

        if sum, ok := usersExpenseSumMap[userId]; ok {
            userExpense.Sum = sum
        }

        report = append(report, userExpense)
    }

    return report
}

func transactionsToUserCategoryExpenseMap(transactions []datajson.Transaction) (map[UserCategoryIndex]UserExpenseCategory) {
    var userExpenseCategory UserExpenseCategory
    var userCategoryIndex UserCategoryIndex

    userCategoryExpenseMap := map[UserCategoryIndex]UserExpenseCategory{}
    for _, transaction := range transactions {
        userCategoryIndex.UserID = transaction.UserID
        userCategoryIndex.Category = transaction.Category

        userExpenseCategory.Name = transaction.Category
        userExpenseCategory.Count = 1
        userExpenseCategory.Sum = transaction.Amount

        if i, ok := userCategoryExpenseMap[userCategoryIndex]; ok {
            userExpenseCategory.Count += i.Count
            userExpenseCategory.Sum += i.Sum
        }

        userCategoryExpenseMap[userCategoryIndex] = userExpenseCategory
    }

    return userCategoryExpenseMap
}

func getUsersExpenseCategoriesAndSum(userCategoryExpenseMap map[UserCategoryIndex]UserExpenseCategory) (map[int][]UserExpenseCategory, map[int]int) {
    usersExpenseCategoriesMap := map[int][]UserExpenseCategory{}
    usersExpenseSumMap := map[int]int{}
    for userCategoryIndex, userExpenseCategory := range userCategoryExpenseMap {
        usersExpenseCategoriesMap[userCategoryIndex.UserID] = append(usersExpenseCategoriesMap[userCategoryIndex.UserID], userExpenseCategory)

        if _, ok := usersExpenseSumMap[userCategoryIndex.UserID]; !ok {
            usersExpenseSumMap[userCategoryIndex.UserID] = 0
        }

        usersExpenseSumMap[userCategoryIndex.UserID] += userExpenseCategory.Sum
    }

    return usersExpenseCategoriesMap, usersExpenseSumMap;
}

func GetUsersExpensesReportFileName() (string) {
    return "users-expenses.json";
}
