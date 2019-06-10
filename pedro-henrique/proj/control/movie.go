package control

import ( "log"
		 "io/ioutil"
		 "net/http"
		 "encoding/json"
		 "pedro-henrique/proj/validation"
		 "pedro-henrique/proj/models"
		 "pedro-henrique/proj/database"
		 "github.com/gorilla/mux" )


func AllMovies(w http.ResponseWriter, r *http.Request) {
	
	movies, err := database.FindAllMovies()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resmovies, err := json.Marshal(movies)
	if err != nil {
		log.Printf("[ERROR] Problem parsing movies to JSON. Error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

    w.WriteHeader(http.StatusOK)
    w.Write(resmovies)
}

func FindMovie(w http.ResponseWriter, r *http.Request) {
	
	params := mux.Vars(r)
    movie, err := database.FindMovieById(params["id"])

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    movieResponse, err := json.Marshal(movie)
    if err != nil {
        log.Printf("[ERROR] Problem parsing movie to JSON. Error: %v", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(movieResponse)

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {

	body := r.Body
	var movie models.Movie
	bytes, _ := ioutil.ReadAll(body)
	err := json.Unmarshal(bytes, &movie)

	if err != nil {
		log.Printf("[WARN] problem parsing json body, because, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := validation.Validator.Struct(movie); err != nil {
		log.Printf("[WARN] invalid movie information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	err = database.InsertMovie(movie)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`{"message": "success"}`))
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	var movie models.Movie
	bytes, _ := ioutil.ReadAll(body)
	err := json.Unmarshal(bytes, &movie)

	if err != nil {
		log.Printf("[WARN] problem parsing json body, because, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := validation.Validator.Struct(movie); err != nil {
		log.Printf("[WARN] invalid movie information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	err = database.UpdateMovie(movie)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`{"message": "success"}`))
}


func DeleteMovie(w http.ResponseWriter, r *http.Request) {

	body := r.Body
	var movie models.Movie
	bytes, _ := ioutil.ReadAll(body)
	err := json.Unmarshal(bytes, &movie)

	if err != nil {
		log.Printf("[WARN] problem parsing json body, because, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := validation.Validator.Struct(movie); err != nil {
		log.Printf("[WARN] invalid movie information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	err = database.DeleteMovie(movie)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`{"message": "success"}`))
}
