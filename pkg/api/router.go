package api

import (
	"mvc/pkg/controller"
	"mvc/pkg/jwt_middleware_handler"
	"mvc/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	r := mux.NewRouter()
	r.Use(jwt_middleware_handler.VerifyTokenMiddleware)

	r.HandleFunc("/", controller.UserLogin).Methods("GET")
	r.HandleFunc("/client/profilePage", controller.ProfilePage).Methods("GET")
	r.HandleFunc("/client/userIssue", controller.UserIssueRequests).Methods("GET")
	r.HandleFunc("/client/userReturn", controller.UserReturnRequest).Methods("GET")
	r.HandleFunc("/client/userBrowse", controller.BrowseBooks).Methods("GET")
	r.HandleFunc("/client/serverError", controller.ClientServerErrorPage).Methods("GET")
	r.HandleFunc("/signUp", controller.AddUserP).Methods("POST")
	r.HandleFunc("/login", controller.LoginUserP).Methods("POST")
	r.HandleFunc("/client/requestIssue", controller.RequestIssue).Methods("POST")
	r.HandleFunc("/client/withdrawIssueRequest", controller.WithdrawIssueRequest).Methods("POST")
	r.HandleFunc("/client/withdrawReturnRequest", controller.WithdrawReturnRequest).Methods("POST")
	r.HandleFunc("/client/requestReturn", controller.RequestReturn).Methods("POST")
	r.HandleFunc("/userLogout", controller.LogoutUser).Methods("POST")

	r.HandleFunc("/adminHome", controller.AdminHome).Methods("GET")
	r.HandleFunc("/admin/addBook", controller.AddPage).Methods("GET")
	r.HandleFunc("/admin/booksInventory", controller.List).Methods("GET")
	r.HandleFunc("/admin/addAdmin", controller.AddAdminPage).Methods("GET")
	r.HandleFunc("/admin/issueRequests", controller.ListIssueRequest).Methods("GET")
	r.HandleFunc("/admin/returnRequests", controller.ListReturnRequest).Methods("GET")
	r.HandleFunc("/admin/viewAdmins", controller.ViewAdmins).Methods("GET")
	r.HandleFunc("/admin/serverError", controller.AdminServerErrorPage).Methods("GET")
	r.HandleFunc("/admin/addBook", controller.AddBook).Methods("POST")
	r.HandleFunc("/adminHome", controller.LoginAdmin).Methods("POST")
	r.HandleFunc("/admin/quantityIncrease", controller.IncreaseQuantity).Methods("POST")
	r.HandleFunc("/admin/quantityDecrease", controller.DecreaseQuantity).Methods("POST")
	r.HandleFunc("/admin/removeBook", controller.RemoveBook).Methods("POST")
	r.HandleFunc("/admin/addAdmin", controller.AddAdmin).Methods("POST")
	r.HandleFunc("/admin/acceptIssueRequest", controller.AcceptIssue).Methods("POST")
	r.HandleFunc("/admin/rejectIssueRequest", controller.RejectIssue).Methods("POST")
	r.HandleFunc("/admin/acceptReturnRequest", controller.AcceptReturn).Methods("POST")
	r.HandleFunc("/admin/rejectReturnRequest", controller.RejectReturn).Methods("POST")
	r.HandleFunc("/admin/removeAdmin", controller.RemoveAdmin).Methods("POST")
	r.HandleFunc("/adminLogout", controller.LogoutAdmin).Methods("POST")

	pathCSS, err := utils.GetCurrentDirPath()
	if err == nil {
		s := http.StripPrefix("/static/", http.FileServer(http.Dir(pathCSS+"/templates/static/")))
		r.PathPrefix("/static/").Handler(s)
	}

	http.ListenAndServe(":8000", r)
}
