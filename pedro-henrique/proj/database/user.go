package database

import ( "go.mongodb.org/mongo-driver/bson"
		 "pedro-henrique/proj/models"
         "context"
         "log" )

func InsertUser(user models.User) (err error) {

	// select users collection
	c := Db.Collection("users")

	// insert user into users collection
	_, err = c.InsertOne(context.TODO(), user)

	if err != nil {
		log.Printf("[ERROR] Could not insert user: %v\n", err)
		return
	}
	return
}

// Deletes a user from the DB.
func RemoveUser(user models.User) (err error) {
	c := Db.Collection("users")
	var errors interface{} 
	err = c.FindOneAndDelete(context.TODO(),&user).Decode(&errors)
	if err != nil {
		log.Printf("[ERROR] Could not delete user. Error: %v", err)
		return
	}

	return
}


// Verifies the existance of a username and its password in the DB.
func VerifyUser(ua models.UserAuth) (userExists bool, err error) {
	c := Db.Collection("users")
	var user models.User

	// Searches for a user that matches the key in question.
	err = c.FindOne(context.TODO(), bson.D{{"email", ua.Email}, {"password", ua.Password}}).Decode(&user)
	if err != nil {
		log.Printf("[ERROR] Could not find user. Error: %v", err)
		return false, err
	}

	return true, nil

}

func UpdateUser(user models.User) (err error) {

	c := Db.Collection("user")
	_, err = c.UpdateOne(context.TODO(), user.Email, &user)
	if err != nil {
		log.Printf("[ERROR] Could not update user: %v %v", err)
		return
	}
	return
	
}
