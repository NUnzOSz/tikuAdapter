package search

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/itihey/tikuAdapter/pkg/errors"
	"github.com/itihey/tikuAdapter/pkg/model"
	"github.com/itihey/tikuAdapter/pkg/util"
	"unicode"
)

type question struct {
	QID      string   `json:"qid"`
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   []string `json:"answer"`
}

type response struct {
	Code  int        `json:"code"`
	Msg   string     `json:"msg"`
	QList []question `json:"qlist"`
}

// AidianClient 爱点题库
type AidianClient struct {
	Disable bool
	YToken  string
}

func (in *AidianClient) getHTTPClient() *resty.Client {
	return resty.New()
}

// SearchAnswer 搜索答案
func (in *AidianClient) SearchAnswer(req model.SearchRequest) (answer [][]string, err error) {
	answer = make([][]string, 0)

	url := "http://new.api.51aidian.com/publics/newapi/freedirect" // 免费接口 会限流
	if in.YToken != "" {
		url = "http://new.api.51aidian.com/publics/newapi/direct" // 付费接口 不会限流按次购买
	}
	client := in.getHTTPClient()

	resp, err := client.R().
		SetBody(map[string]string{
			"question": req.Question,
			"token":    in.YToken,
		}).
		Post(url)

	var res response

	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return nil, errors.ErrParserJSON
	}
	for _, q := range res.QList {
		q.Options = util.FormatOptions(q.Options)
		var as = make([]string, 0)
		for _, s := range q.Answer {
			if isAlpha(s) { // ABCD 或者A 这种
				for _, i := range s {
					index := int(i - 65)
					if len(q.Options) > index {
						as = append(as, q.Options[index])
					}
				}
			} else {
				as = append(as, s)
			}
		}
		if len(as) > 0 {
			answer = append(answer, as)
		}
	}
	return answer, nil
}

func isAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}