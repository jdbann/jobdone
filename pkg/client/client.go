package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"go.uber.org/zap"
)

type Client struct {
	url *url.URL

	logger *zap.Logger
}

type Params struct {
	URL *url.URL

	Logger *zap.Logger
}

func New(params Params) Client {
	if params.URL == nil {
		var err error
		params.URL, err = url.Parse("http://localhost:3000")
		if err != nil {
			panic(err)
		}
	}

	if params.Logger == nil {
		params.Logger = zap.NewNop()
	}

	return Client{
		url: params.URL,

		logger: params.Logger.Named("Client"),
	}
}

func (c Client) Healthcheck() (*http.Response, error) {
	return c.get("/health", nil)
}

type registerPersonBody struct {
	ClientID string `json:"client_id"`
}

type registerPersonResponse struct {
	ID string `json:"id"`
}

func (c Client) RegisterPerson(localID string) (string, error) {
	var response registerPersonResponse
	_, err := c.post("/person", registerPersonBody{ClientID: localID}, &response)
	if err != nil {
		c.logger.Error("error attempt to register person", zap.Error(err))
		return "", err
	}

	if response.ID == "" {
		return "", errors.New("ID from server is blank")
	}

	return response.ID, nil
}

func (c Client) get(path string, res interface{}) (*http.Response, error) {
	url, err := c.url.Parse(path)
	if err != nil {
		return nil, err
	}

	response, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}

	if res == nil {
		return response, nil
	}

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c Client) post(path string, req interface{}, res interface{}) (*http.Response, error) {
	url, err := c.url.Parse(path)
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(url.String(), "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return nil, err
	}

	return response, nil
}
