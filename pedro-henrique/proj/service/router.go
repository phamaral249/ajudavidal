package service


import ( "net/http"
	     "pedro-henrique/proj/configs"
	     "pedro-henrique/proj/control"
		 "github.com/gorilla/mux"
		 "github.com/rs/cors"  )

func createHandler() (handler *mux.Router) {

	// creates router
	handler = mux.NewRouter()

	handler.HandleFunc(configs.USER_PATH, control.RegisterUser).Methods(http.MethodPost).Headers("content-type","application/json")
	handler.HandleFunc(configs.USER_PATH, control.RemoveUser).Methods(http.MethodDelete).Headers("content-type","application/json")
	handler.HandleFunc(configs.USER_PATH, control.UpdateUser).Methods(http.MethodPut).Headers("content-type","application/json")
	handler.HandleFunc(configs.LOGIN_PATH, control.LoginUser).Methods(http.MethodPost).Headers("content-type","application/json")

	
	handler.HandleFunc(configs.MOVIE_PATH, control.AllMovies).Methods(http.MethodGet).Headers("content-type","application/json")
	handler.HandleFunc(configs.MOVIE_PATH, control.CreateMovie).Methods(http.MethodPost).Headers("content-type","application/json")
	handler.HandleFunc(configs.MOVIE_PATH, control.UpdateMovie).Methods(http.MethodPut).Headers("content-type","application/json")
	handler.HandleFunc(configs.MOVIE_PATH, control.DeleteMovie).Methods(http.MethodDelete).Headers("content-type","application/json")
	handler.HandleFunc(configs.MOVIE_PATH+"{id}", control.FindMovie).Methods(http.MethodGet).Headers("content-type","application/json").Queries("id","{value}")

	// returns handle
	return
}

func createCORS() *cors.Cors {

	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Accept", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token"},
	})
}
