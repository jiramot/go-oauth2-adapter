package restclient

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/url"
    "strings"
)

type HTTPClient interface {
    Do(req *http.Request) (*http.Response, error)
}

var (
    Client HTTPClient
)

func init() {
    Client = &http.Client{}
}

func Get(url string) (*http.Response, error) {
    request, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }
    return Client.Do(request)
}

func PostForm(url string, data url.Values) (*http.Response, error) {
    request, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
    request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    if err != nil {
        return nil, err
    }
    return Client.Do(request)
}

func PostJson(url string, value interface{}) (*http.Response, error) {
    payloadBuf := new(bytes.Buffer)
    json.NewEncoder(payloadBuf).Encode(value)
    request, err := http.NewRequest("POST", url, payloadBuf)
    request.Header.Set("Content-Type", "application/json")
    if err != nil {
        return nil, err
    }
    return Client.Do(request)
}
