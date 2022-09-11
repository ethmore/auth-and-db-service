package services

import (
	"bytes"
	"e-comm/authService/dotEnv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type SearchProduct struct {
	Id          string
	Title       string
	Price       string
	Description string
	Image       string
	Stock       string
}

type Message struct {
	Message string
}

func AddProductToSearchService(p SearchProduct) error {
	body, _ := json.Marshal(p)
	jsonBody := []byte(body)
	bodyReader := bytes.NewReader(jsonBody)
	requestURL := dotEnv.GoDotEnvVariable("ADD_PRODUCT_SEARCH_SERVICE")

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, clientErr := client.Do(req)
	if clientErr != nil {
		return err
	}

	//Read response
	b, readErr := io.ReadAll(res.Body)
	if err != nil {
		return readErr
	}
	defer res.Body.Close()

	var response Message
	if unmarshalErr := json.Unmarshal([]byte(b), &response); err != nil {
		return unmarshalErr
	}
	fmt.Println(response)
	return nil
}
