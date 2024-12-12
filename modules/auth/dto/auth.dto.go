package authdto

type LoginBody struct {
	UserId   string `json:"userId" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Details struct {
	Id      int64  `json:"id"`
	UserId  string `json:"userId"`
	Name    string `json:"name"`
	Images  string `json:"images"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type GoogleAuthRequest struct {
	RedirectURL string `form:"redirect_url" binding:"required"`
}

type StateRequest struct {
	Prefix      string `json:"prefix"`
	RedirectURL string `form:"redirect_url"`
}

type GoogleUserResponse struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
}

type RegisterBody struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Images   string `form:"images" binding:"required"`
	Password string `form:"password" binding:"required"`
	Address  string `form:"address"`
	Phone    int64  `form:"phone"`
}
