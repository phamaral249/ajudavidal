package service

import ( "log"
		 "time"
		 "net/http"
		 "pedro-henrique/proj/configs"
		 "pedro-henrique/proj/validation" )

func StartServer() {

	// creates a handler (router or multiplexer)
	h := createHandler()

	// creates a HTTP server with default parameters
	s := createServer()

	// creates a CORS struct to deal with cross origin
	c := createCORS()

	// set cors treatment on router
	handler := c.Handler(h)

	// associate handler to server
	s.Handler = handler

	// create global validator
	validation.CreateValidator()

	// instanciates a HTTP server wrapped in a log fatal 
	log.Fatal(s.ListenAndServe())
}

func StopServer() {
	
}

func createServer() (server *http.Server){


	// create a http server instance
	server = &http.Server {

		Addr: configs.SERVER_ADDR,  
		IdleTimeout:  100 * time.Millisecond,
		ReadTimeout:  100 * time.Millisecond,
		WriteTimeout: 100 * time.Millisecond,

	}

	return
}