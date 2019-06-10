package control

import ( "log"
		 "io/ioutil"
		 "net/http"
		 "encoding/json"
		 "pedro-henrique/proj/validation"
		 "pedro-henrique/proj/models"
		 "pedro-henrique/proj/database" )


func RegisterUser(w http.ResponseWriter, r *http.Request) {

	// retrieve body from request
	body := r.Body

	// declare use entity
	var user models.User

	// parsing io.ReadCLoser to slice of bytes []byte
	bytes, _ := ioutil.ReadAll(body)

	// parses json into user struct
	err := json.Unmarshal(bytes, &user)

	// checks if any error occurs in json parsing  
	if err != nil {

		log.Printf("[WARN] problem parsing json body, because, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// checks if struct is a valid one
	if err := validation.Validator.Struct(user); err != nil {

		log.Printf("[WARN] invalid user information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	err = database.InsertUser(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message" : "User inserted successfuly."}`))

}

func RemoveUser(w http.ResponseWriter, r *http.Request) {
	
	body := r.Body
	var user models.User
	bytes, _ := ioutil.ReadAll(body)
	err := json.Unmarshal(bytes, &user)

	if err != nil {
		log.Printf("[WARN] problem parsing json body, because, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := validation.Validator.Struct(user); err != nil {
		log.Printf("[WARN] invalid user information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	err = database.RemoveUser(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`{"message": "success"}`))

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	
	body := r.Body
	var user models.User
	bytes, _ := ioutil.ReadAll(body)
	err := json.Unmarshal(bytes, &user)

	if err != nil {
		log.Printf("[WARN] problem parsing json body, because, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := validation.Validator.Struct(user); err != nil {
		log.Printf("[WARN] invalid user information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	err = database.UpdateUser(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`{"message": "success"}`))

}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	
	body := r.Body
	var ua models.UserAuth
	bytes, _ := ioutil.ReadAll(body)
	err := json.Unmarshal(bytes, &ua)

	if err != nil {
		log.Printf("[WARN] problem parsing json body, because, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := validation.Validator.Struct(ua); err != nil {
		log.Printf("[WARN] invalid userAuth information, because, %v\n", err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	userExists , err := database.VerifyUser(ua)
	if userExists != false {
		w.Write([]byte(`{"message": "access denied"}`))
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`{"message": "success"}`))

}