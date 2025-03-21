package search

import (
	"encoding/json"
	"strings"
	"tikuAdapter/pkg/errors"
	"tikuAdapter/pkg/model"
	"time"

	"github.com/go-resty/resty/v2"
)

// 柠檬题库官网 https://www.lemtk.xyz

// LemonClient 柠檬题库
type LemonClient struct {
	Enable bool
	Token  string
}

// getHTTPClient 获取http客户端
func (in *LemonClient) getHTTPClient() *resty.Client {
	return resty.New().SetTimeout(5 * time.Second)
}

// SearchAnswer 搜索答案
func (in *LemonClient) SearchAnswer(req model.SearchRequest) (answer [][]string, err error) {
	answer = make([][]string, 0)
	if !in.Enable || in.Token == "" {
		return answer, nil
	}

	post, err := in.getHTTPClient().R().
		SetHeader("Authorization", "Bearer "+in.Token).
		SetHeader("content-type", "application/json").
		SetBody(map[string]string{
			"v":        "1.0",
			"question": req.Question,
			"uid":      "703382225",
		}).
		Post("https://api.lemtk.xyz/api/v1/mcx")
	if err != nil || post.StatusCode() != 200 {
		return nil, errors.ErrTargetServerError
	}

	var response lemonResp
	err = json.Unmarshal(post.Body(), &response)
	if err != nil {
		return nil, errors.ErrTargetServerError
	}
	if response.Code == 1000 {
		ans := strings.Split(response.Data.Answer, "#")
		return [][]string{ans}, nil
	}
	return nil, errors.ErrTargetNoAnswer
}

type lemonResp struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data lemonData `json:"data"`
}

type lemonData struct {
	Answer string `json:"answer"`
}
