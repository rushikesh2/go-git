package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	contentType = "application/json"
)

var (
	ErrInvalidURL = errors.New("Invalid url")
)

type HttpResponse struct {
	StatusCode   int
	ResponseBody []byte
}

func TrackTime(start time.Time, url string) {
	elapsed := time.Since(start).Seconds()
	log.Println(fmt.Sprintf("Time spent: %v (s), on action: %v", elapsed, url))
}

func MakeRequest(method, url string) (response HttpResponse, err error) {
	client := &http.Client{}
	if url == "" {
		return response, ErrInvalidURL
	}

	defer TrackTime(time.Now(), url)

	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	request.Header.Add("Accept", contentType)
	request.Header.Add("Content-Type", contentType)
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
		return response, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err.Error())
		return response, err
	}
	httpResponse := HttpResponse{
		StatusCode:   resp.StatusCode,
		ResponseBody: data,
	}
	return httpResponse, nil
}
