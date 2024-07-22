package types

type PlacesProps struct {
	Title string `json:"title"`
	ImgURL string `json:"img_url"`
	Description string `json:"description"`
}

type UserSignUpProps struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}