package infrastructure

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"encoding/json"
	"fmt"

	"github.com/chunnior/api-gateway/internal/domain"
)

type HTTPUserService struct {
	userServiceURL string
	httpClient     *http.Client
}

func NewHTTPUserService(userServiceURL string, httpClient *http.Client) *HTTPUserService {
	return &HTTPUserService{
		userServiceURL: userServiceURL,
		httpClient:     httpClient,
	}
}

func (s *HTTPUserService) Login(requestBody domain.LoginUserServiceRequest) (domain.LoginUserServiceResponse, error) {
	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return domain.LoginUserServiceResponse{}, err
	}
	resp, err := http.Post(s.userServiceURL+"/login", "application/json", bytes.NewBuffer(requestJSON))
	if err != nil {
		return domain.LoginUserServiceResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		return domain.LoginUserServiceResponse{}, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, bodyString)
	}

	var loginResponse domain.LoginUserServiceResponse

	if err := json.NewDecoder(resp.Body).Decode(&loginResponse); err != nil {
		return domain.LoginUserServiceResponse{}, err
	}

	return loginResponse, nil
}

func (s *HTTPUserService) Callback(requestBody domain.LoginCallbackParams) (domain.CallbackResponse, error) {
	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return domain.CallbackResponse{}, err
	}

	resp, err := http.Post(s.userServiceURL+"/callback", "application/json", bytes.NewBuffer(requestJSON))
	if err != nil {
		return domain.CallbackResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		return domain.CallbackResponse{}, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, bodyString)
	}

	var callbackResponse domain.CallbackResponse

	if err := json.NewDecoder(resp.Body).Decode(&callbackResponse); err != nil {
		return domain.CallbackResponse{}, err
	}

	return callbackResponse, nil
}
