// Package interest implements functions to deal with your balance and its interest rate.
package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	result := 3.213
	if balance >= 0 && balance < 1000 {
		result = 0.5
	} else if balance >= 1000 && balance < 5000 {
		result = 1.621
	} else if balance >= 5000 {
		result = 2.475
	}
	return float32(result)
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * (float64(InterestRate(balance)) / 100.0)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	years := 0
	for index := balance; index < targetBalance; index = AnnualBalanceUpdate(index) {
		years += 1
	}
	return years
}
