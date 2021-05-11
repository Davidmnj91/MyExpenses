package expense

type Repository interface {
	AddExpense(*Expense) (*Expense, error)
}
