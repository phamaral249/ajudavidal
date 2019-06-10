package database

import ( "go.mongodb.org/mongo-driver/bson"
		 "pedro-henrique/proj/models"
         "context"
         "log" )

// Find list of movies
func FindAllMovies() (ms []models.Movie, err error) {

	c := Db.Collection("movies")
	res, err := c.Find(context.TODO(),bson.D{})

	if err != nil {
		log.Printf("[ERROR] Could not find movies: %v %v", err, ms)
		return
	}
	err = res.All(context.TODO(), &ms)
	if err != nil {
		log.Printf("[ERROR] Could not extract result query to movies array: %v", err)
		return
	}

	return
}

// Find a movie by its id
func FindMovieById(id string) (m models.Movie, err error) {

	c := Db.Collection("movies")
	err = c.FindOne(context.TODO(), bson.D{{"ID", id}}).Decode(&m)

	if err != nil {
		log.Printf("[ERROR] Could not search movie: %v %v", err)
		return
	}
	return
}

// Insert a movie into database
func InsertMovie(movie models.Movie) (err error) {

	c := Db.Collection("movies")
	_, err = c.InsertOne(context.TODO(), movie)
	if err != nil {
		log.Printf("[ERROR] Could not insert movie: %v %v", err)
		return
	}
	return
}

// Delete an existing movie
func DeleteMovie(movie models.Movie) (err error) {

	c := Db.Collection("movies")
	var errors interface{} 
	err = c.FindOneAndDelete(context.TODO(),&movie).Decode(&errors)
	if err != nil {
		log.Printf("[ERROR] Could not remove movie: %v %v", err)
		return
	}
	return
}

// Update an existing movie
func UpdateMovie(movie models.Movie) (err error) {

	c := Db.Collection("movies")
	_, err = c.UpdateOne(context.TODO(), movie.ID, &movie)
	if err != nil {
		log.Printf("[ERROR] Could not update movie: %v %v", err)
		return
	}
	return
	
}

