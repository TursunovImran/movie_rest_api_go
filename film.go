package movierestapigo

type Film struct{
	ID 				int 	`json:"id"`
	Title 			string 	`json:"title"`
	Description 	string 	`json:"description"`
	Release_date 	string 	`json:"release_date" format:"2000-01-01"`
	Rating      	float32 `json:"rating" validate:"min=0,max=10"`
	Actors      	[]Actor `json:"actors"`
}