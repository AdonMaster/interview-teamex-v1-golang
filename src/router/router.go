package router

import (
	"interview-teamex-v1/src/controllers"
	"interview-teamex-v1/src/repo"
	"net/http"
)

func Init(repo *repo.Repo) *http.ServeMux {
	//
	router := http.NewServeMux()

	// auth
	expenseController := controllers.ExpenseController{Repo: repo}
	router.HandleFunc("GET /expenses", expenseController.Get)
	router.HandleFunc("GET /expenses/summary", expenseController.Summary)

	router.HandleFunc("POST /expenses", expenseController.Create)
	router.HandleFunc("PUT /expenses/{id}", expenseController.Update)
	router.HandleFunc("DELETE /expenses/{id}", expenseController.Destroy)

	//
	return router
}
