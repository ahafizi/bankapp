package stats

import (
	"fmt"
	"github.com/ahafizi/bank/pkg/bank/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID: 1,
			Amount: 100_00,
			Category: "shirt",
		},
		{
			ID: 1,
			Amount: 150_00,
			Category: "jacket",
		},
		{
			ID: 1,
			Amount: 500_00,
			Category: "dress",
		},
	}

	fmt.Println(Avg(payments))

	// Output:
	// 25000
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID: 1,
			Amount: 100_00,
			Category: "shirt",
		},
		{
			ID: 2,
			Amount: 150_00,
			Category: "jacket",
		},
		{
			ID: 3,
			Amount: 500_00,
			Category: "dress",
		},
		{
			ID: 4,
			Amount: 500_00,
			Category: "dress",
		},
	}

	fmt.Println(TotalInCategory(payments, "dresss"))

	// Output:
	// 250_00
}
