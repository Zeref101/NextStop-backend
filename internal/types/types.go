package types

type PlacesProps struct {
	Title string `json:"title"`
	ImgURL string `json:"img_url"`
	Description string `json:"description"`
}

type User struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type Sign_in_props struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func (User) TableName() string{

	return "User"
}