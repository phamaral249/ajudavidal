package models

type Movie struct{
 
	ID          int       `json:"id"       validate:"required"`
	Title       string    `json:"title"    validate:"required"`
	Genre       string    `json:"genre"    validate:"required"`
	Description string    `json:"description"`
	Budget      int       `json:"budget"`
	

}
