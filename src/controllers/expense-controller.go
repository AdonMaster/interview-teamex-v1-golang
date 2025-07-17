package controllers

import (
	"fmt"
	"interview-teamex-v1/src/db/model"
	"interview-teamex-v1/src/repo"
	"interview-teamex-v1/src/repo/opts"
	"interview-teamex-v1/src/response"
	"interview-teamex-v1/src/validator"
	"net/http"
	"strconv"
)

type ExpenseController struct {
	Repo *repo.Repo
}

func extractFilterFrom(req *http.Request) []opts.ExpenseGetFilterOpt {
	// filter
	var filter []opts.ExpenseGetFilterOpt

	// normal case
	page := req.URL.Query().Get("page")
	if page != "" {
		filter = append(filter, opts.WithPage(page))
	}
	pageSize := req.URL.Query().Get("pageSize")
	if pageSize != "" {
		filter = append(filter, opts.WithPageSize(pageSize))
	}
	category := req.URL.Query().Get("category")
	if category != "" {
		filter = append(filter, opts.WithCategory(category))
	}
	dtIni := req.URL.Query().Get("dtIni")
	if dtIni != "" {
		filter = append(filter, opts.WithDtIni(dtIni))
	}
	dtEnd := req.URL.Query().Get("dtEnd")
	if dtEnd != "" {
		filter = append(filter, opts.WithDtEnd(dtEnd))
	}

	return filter
}

func (c ExpenseController) Get(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()

	//
	filter := extractFilterFrom(req)

	// retrieving models
	models, err := c.Repo.ExpenseGet(ctx, filter)
	if err != nil {
		response.New400(err.Error()).Write(w)
		return
	}

	// summary too
	totalIncome, totalExpense, err := c.Repo.ExpenseGetSummary(ctx, filter)
	if err != nil {
		response.New400(err.Error()).Write(w)
		return
	}

	// returning - (default 200)
	data := map[string]any{
		"total_income":  totalIncome,
		"total_expense": totalExpense,
		"data":          models,
	}
	response.New(response.WithPayload(data)).Write(w)
}

func (c ExpenseController) Summary(w http.ResponseWriter, req *http.Request) {
	//
	ctx := req.Context()
	//
	filter := extractFilterFrom(req)
	filter = append(filter, opts.WithSummary()) // <-- the key here

	//
	totalIncome, totalExpense, err := c.Repo.ExpenseGetSummary(ctx, filter)
	if err != nil {
		response.New400(err.Error()).Write(w)
		return
	}

	//
	response.New(
		response.WithPayload(
			map[string]float64{
				"total_income":  totalIncome,
				"total_expense": totalExpense,
			},
		),
	).Write(w)
}

func (c ExpenseController) Create(w http.ResponseWriter, req *http.Request) {

	var data struct {
		Date     string            `json:"date" sv:"required"`
		Amount   float64           `json:"amount" v:"required"`
		Type     model.ExpenseType `json:"type" v:"required"`
		Category string            `json:"category" v:"required"`
	}
	if !validator.Validate(w, req, &data) {
		return
	}

	//
	ctx := req.Context()

	//
	expense, err := c.Repo.ExpenseCreate(ctx, data.Date, data.Amount, data.Type, data.Category)
	if err != nil {
		response.New400(err.Error()).Write(w)
		return
	}

	//
	response.New(response.WithPayload(expense)).Write(w)
}

func (c ExpenseController) Update(w http.ResponseWriter, req *http.Request) {

	//
	ctx := req.Context()

	// I'll trust this param and not the json body, since it's more REST-like...
	// But, most cases, I just use one route 'save' and check if id is non-zero...
	// not today
	idStr := req.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	//
	var data struct {
		Date     string            `json:"date" sv:"required"`
		Amount   float64           `json:"amount" v:"required"`
		Type     model.ExpenseType `json:"type" v:"required"`
		Category string            `json:"category" v:"required"`
	}
	if !validator.Validate(w, req, &data) {
		return
	}

	//
	err := c.Repo.ExpenseUpdate(ctx, data.Date, data.Amount, data.Type, data.Category, id)
	if err != nil {
		response.New400(err.Error()).Write(w)
		return
	}

	//
	response.New(response.WithMessage("Model updated successfully")).Write(w)
}

func (c ExpenseController) Destroy(w http.ResponseWriter, req *http.Request) {
	//
	ctx := req.Context()

	// id parse
	idStr := req.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	// deleting
	tag, err := c.Repo.Conn.Exec(ctx, `DELETE FROM expenses WHERE id = $1`, id)
	if err != nil {
		response.New400(err.Error()).Write(w)
		return
	}

	//
	message := fmt.Sprintf("rows affected: %d", tag.RowsAffected())
	response.New(response.WithMessage(message)).Write(w)

}
