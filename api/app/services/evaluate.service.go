package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/devmeireles/go-vue-image-evaluate/app/dto"
	"github.com/devmeireles/go-vue-image-evaluate/app/utils"
)

// CreateEvaluateRequest do a request on cloudmersive API
// Possible classifications outcoming: SafeContent_HighProbability,
// UnsafeContent_HighProbability, RacyContent, SafeContent_ModerateProbability
func CreateEvaluateRequest(imageAddress string) (*dto.GetEvaluateDTO, error) {

	url := "https://api.cloudmersive.com/image/nsfw/classify"
	method := "POST"

	payload, err := utils.ImageLoader(imageAddress)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "multipart/form-data")
	req.Header.Add("Apikey", os.Getenv("API_KEY"))

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))

	var dataRes dto.GetEvaluateDTO
	err = json.Unmarshal(body, &dataRes)

	if err != nil {
		return nil, err
	}

	return &dataRes, nil
}
