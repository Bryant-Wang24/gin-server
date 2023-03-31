package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"example.com/blog/model"
)

var ClientSecret = "6868107ee8817a97f1c4731c579f82f7ac97b358"
var ClientId = "Iv1.355ecd3abdf8db39"

var conf = model.Conf{
	ClientId:     ClientId,
	ClientSecret: ClientSecret,
	RedirectUrl:  "http://localhost:8080/articles",
}

// GetToken 通过url发送请求，获取access_token
func GetToken(code string) (string, error) {
	url := `https://github.com/login/oauth/access_token?client_id=` + conf.ClientId + `&client_secret=` + conf.ClientSecret + `&code=` + code
	// 发送请求
	var req *http.Request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("http.NewRequestErr", err)
	}
	// 设置请求头
	req.Header.Set("Accept", "application/json")
	// 发送请求
	var response *http.Response
	response, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("http.DefaultClient.DoErr", err)
	}
	// 关闭响应
	defer response.Body.Close()
	// 读取响应
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("ioutil.ReadAllErr", err)
	}
	// 拿到access_token
	var token map[string]interface{}
	err = json.Unmarshal(body, &token)
	if err != nil {
		fmt.Println("tokenJson.UnmarshalErr", err)
	}
	return token["access_token"].(string), nil
}

// GetUserInfo 通过access_token获取用户信息
func GetUserInfo(accessToken string) (map[string]interface{}, error) {
	url := `https://api.github.com/user`
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("http.NewRequestErr", err)
	}
	request.Header.Add("Authorization", "token "+accessToken)
	request.Header.Add("Accept", "application/json")
	// 发送请求
	var response *http.Response
	response, err = http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("http.DefaultClient.DoErr", err)
	}
	// 关闭响应
	defer response.Body.Close()
	// 读取响应
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAllErr", err)
	}
	// 拿到用户信息
	var userInfo map[string]interface{}
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		fmt.Println("userInfoJson.UnmarshalErr", err)
	}
	return userInfo, nil
}
