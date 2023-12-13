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

func (s *HTTPUserService) Callback(requestBody domain.LoginCallbackParams) (domain.UserPayload, error) {
	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return domain.UserPayload{}, err
	}

	resp, err := http.Post(s.userServiceURL+"/callback", "application/json", bytes.NewBuffer(requestJSON))
	if err != nil {
		return domain.UserPayload{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		return domain.UserPayload{}, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, bodyString)
	}

	var userPayload domain.UserPayload

	if err := json.NewDecoder(resp.Body).Decode(&userPayload); err != nil {
		return domain.UserPayload{}, err
	}

	return userPayload, nil
}

func (s *HTTPUserService) DataInfo(params domain.DataInfoParams) (domain.DataInfoResponse, error) {

	url := fmt.Sprintf("%s/%s/%s/%s", s.userServiceURL, params.Provider, params.DataType, params.UserID)
	resp, err := http.Get(url)
	if err != nil {
		return domain.DataInfoResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		return domain.DataInfoResponse{}, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, bodyString)
	}
	var dataInfoResponse domain.DataInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&dataInfoResponse); err != nil {
		return domain.DataInfoResponse{}, err
	}
	return dataInfoResponse, nil
}
