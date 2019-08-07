package internal

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

type BaleHook struct {
	baseUrl string
	token   string
}

func NewBale(baseUrl string, token string) *BaleHook {
	return &BaleHook{
		baseUrl: baseUrl,
		token:   token,
	}
}

func (b *BaleHook) Send(text string) error {
	url := b.baseUrl + "/v1/webhooks/" + b.token
	reader := strings.NewReader(fmt.Sprintf(`{"text":"%s"}`, text))
	resp, err := http.Post(url, "application/json", reader)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	return nil
}