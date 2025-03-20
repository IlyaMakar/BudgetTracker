package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

//Transaction struct

type Transactions struct {
	ID       int
	Amount   float64
	Category string
	Date     time.Time
	Type     string
}

//badgerTracker

type BudgetTracker struct {
	transactions []Transactions
	nextID       int
}

//Interface for common behavior

type FinancialRecord interface {
	GetAmount() float64
	GetType() string
}

//Implament interface methods

func (t Transactions) GetAmount() float64 {
	return t.Amount
}

func (t Transactions) GetType() string {
	return t.Type
}

func (bt *BudgetTracker) AddTransaction(amount float64, category, tType string) {
	newTransactions := Transactions{
		ID:       bt.nextID,
		Amount:   amount,
		Category: category,
		Date:     time.Now(),
		Type:     tType,
	}

	bt.transactions = append(bt.transactions, newTransactions)
	bt.nextID++
}

// Createing DisplayTransactions
func (bt BudgetTracker) DisplayTransactions() {
	fmt.Println("ID\tAmount\tCategory\tDate\tType")
	for _, transactions := range bt.transactions {
		fmt.Printf("%d\t%.2f\t%s\t%s\t%s\n",
			transactions.ID, transactions.Amount, transactions.Category, transactions.Date.Format("2006-01-02"), transactions.Type)
	}
}

// Get total income or expensefun
func (bt BudgetTracker) CalculateTotal(tType string) float64 {
	var total float64
	for _, transactions := range bt.transactions {
		if transactions.Type == tType {
			total += transactions.Amount
		}
	}
	return total
}

// Save the transactions to a csv file
func (bt BudgetTracker) SaveToCSV(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush() //Creating a new CSV file
	//Write the CSV header
	writer.Write([]string{"ID", "Amount", "Category", "Date", "Type"})

	//Write Data
	for _, t := range bt.transactions {
		record := []string{
			strconv.Itoa(t.ID),
			fmt.Sprintf("%.2f", t.Amount),
			t.Category,
			t.Date.Format("2006-01-02"),
			t.Type,
		}
		writer.Write(record)
	}
	fmt.Println("Transactions saved to ", filename)
	return nil
}

// Main function
func main() {
	bt := BudgetTracker{}
	for {
		fmt.Println("\n--- Personal Budget Tracker ---")
		fmt.Println("1. Add Transaction")
		fmt.Println("2. Display Transactions")
		fmt.Println("3. Show Total Income")
		fmt.Println("4. Show Total Expenses")
		fmt.Println("5. save Transaction to CSV")
		fmt.Println("6. Exit ")
		fmt.Println("choose at option: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter Amount: ")
			var amount float64
			fmt.Scanln(&amount)

			fmt.Print("Enter Category: ")
			var category string
			fmt.Scanln(&category)

			fmt.Print("Enter type (Income / Expense): ")
			var tType string
			fmt.Scanln(&tType)

			bt.AddTransaction(amount, category, tType)
			fmt.Print("Transection Addd!")

		case 2:
			bt.DisplayTransactions()
		case 3:
			fmt.Printf("Total income: %.2f\n", bt.CalculateTotal("Income"))
		case 4:
			fmt.Printf("Total expenses: %.2f\n", bt.CalculateTotal("Expenses"))
		case 5:
			fmt.Print("Enter filename (e.g. transactions.csv): ")
			var filename string
			fmt.Scan(&filename)
			if err := bt.SaveToCSV(filename); err != nil {
				fmt.Println("Error saving transactions", err)
			}
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid Choice! Try Again.")
		}
	}
}
