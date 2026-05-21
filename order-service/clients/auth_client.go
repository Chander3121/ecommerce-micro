package clients

import (
	"encoding/json"
	"net/http"
  "fmt"
)

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type ValidateResponse struct {
	Valid bool `json:"valid"`
	User  User `json:"user"`
}

func ValidateToken(token string) (*ValidateResponse, error) {
	client := &http.Client{}

	req, err := http.NewRequest(
		"POST",
		"http://auth-service:3001/validate",
		nil,
	)

	if err != nil {
		return nil, err
	}

	req.Header.Set(
		"Authorization",
		token,
	)

	resp, err := client.Do(req)

  if err != nil {
    return nil, err
  }

  fmt.Println("STATUS CODE:", resp.StatusCode)

  defer resp.Body.Close()

  var validateResponse ValidateResponse

  err = json.NewDecoder(resp.Body).Decode(
    &validateResponse,
  )

  if err != nil {
    return nil, err
  }

  fmt.Printf("VALIDATE RESPONSE: %+v\n", validateResponse)

  return &validateResponse, nil
}
