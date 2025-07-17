package opts

import (
	"strconv"
)

type ExpenseGetFilter struct {
	IsSummary bool
	Page      *int
	PageSize  *int
	Category  *string
	DtIni     *string
	DtEnd     *string
}
type ExpenseGetFilterOpt func(*ExpenseGetFilter)

func WithSummary() ExpenseGetFilterOpt {
	return func(filter *ExpenseGetFilter) {
		filter.IsSummary = true
	}
}
func WithPage(page string) ExpenseGetFilterOpt {
	i, _ := strconv.Atoi(page)
	return func(filter *ExpenseGetFilter) {
		filter.Page = &i
	}
}
func WithPageSize(pageSize string) ExpenseGetFilterOpt {
	i, _ := strconv.Atoi(pageSize)
	return func(filter *ExpenseGetFilter) {
		filter.PageSize = &i
	}
}
func WithCategory(category string) ExpenseGetFilterOpt {
	return func(filter *ExpenseGetFilter) {
		filter.Category = &category
	}
}
func WithDtIni(dtIni string) ExpenseGetFilterOpt {
	return func(filter *ExpenseGetFilter) {
		filter.DtIni = &dtIni
	}
}
func WithDtEnd(dtEnd string) ExpenseGetFilterOpt {
	return func(filter *ExpenseGetFilter) {
		filter.DtEnd = &dtEnd
	}
}

func ExpenseGetFilterOptFrom(options []ExpenseGetFilterOpt) *ExpenseGetFilter {
	var filter ExpenseGetFilter
	for _, opt := range options {
		opt(&filter)
	}
	return &filter
}
