package domain

type LoginRequest struct {
	Provider string `json:"provider"`
}

type LoginUserServiceRequest struct {
	Provider    string `json:"provider"`
	CallbackURL string `json:"callback_url"`
}

type LoginUserServiceResponse struct {
	Url      string `json:"url"`
	Provider string `json:"provider"`
}

type LoginCallbackParams struct {
	Code     string `json:"code"`
	State    string `json:"state"`
	Provider string `json:"provider"`
}

type CallbackResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserService interface {
	Login(requestBody LoginUserServiceRequest) (LoginUserServiceResponse, error)
	Callback(params LoginCallbackParams) (CallbackResponse, error)
}
