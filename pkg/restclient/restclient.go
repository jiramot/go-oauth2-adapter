package restclient

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
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

func PostJson(url string, value interface{}, response interface{}) error {
    payloadBuf := new(bytes.Buffer)
    json.NewEncoder(payloadBuf).Encode(value)
    request, err := http.NewRequest("POST", url, payloadBuf)
    request.Header.Set("Content-Type", "application/json")
    if err != nil {
        return err
    }
    res, err := Client.Do(request)
    if err != nil {
        return err
    }
    bytes, _ := ioutil.ReadAll(res.Body)
    defer res.Body.Close()
    if err := json.Unmarshal(bytes, &response); err != nil {
        return err
    }
    return nil
}
