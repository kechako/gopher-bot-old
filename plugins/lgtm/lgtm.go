package lgtm

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	bot "github.com/kechako/gopher-bot"
	"github.com/kechako/gopher-bot/utils"
)

const (
	keyword    = "lgtm"
	lgtmGenUrl = "http://www.lgtm.in/g"
)

type plugin struct {
	client *http.Client
}

func NewPlugin() bot.Plugin {
	return &plugin{
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			CheckRedirect: checkRedirect,
		},
	}
}

func (p *plugin) DoAction(event bot.EventInfo) bool {
	if !utils.HasKeywords(event.Text(), keyword) {
		return false
	}

	go p.GetAndPostLGTM(event)

	return true
}

func (p *plugin) GetAndPostLGTM(event bot.EventInfo) {
	data, err := getLGTM(p.client)
	if err != nil {
		event.PostMessage(fmt.Sprintf("LGTM ERROR : %v", err))
		return
	}

	event.PostMessage(data.ImageUrl)
}

type lgtmData struct {
	DataUrl  string `json:"dataUrl"`
	ImageUrl string `json:"imageUrl"`
}

func getLGTM(client *http.Client) (*lgtmData, error) {
	req, err := http.NewRequest(http.MethodGet, lgtmGenUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data := &lgtmData{}
	err = json.NewDecoder(res.Body).Decode(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func checkRedirect(req *http.Request, via []*http.Request) error {
	if len(via) >= 10 {
		return errors.New("stopped after 10 redirects")
	}

	if len(via) == 0 {
		return nil
	}

	// Copy header
	for attr, val := range via[0].Header {
		if _, ok := req.Header[attr]; !ok {
			req.Header[attr] = val
		}
	}

	return nil
}

func (p *plugin) Help() string {
	return `lgtm:
	lgtm に反応して LGTM を貼り付けます。
`
}
