package dingtalk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/url"
	"strconv"
	"time"
)

// Client dingtalk robot client
type Client struct {
	// Webhook dingtalk robot webhook
	Webhook string
	// Secret dingtalk robot secret
	Secret string
	client *resty.Client
}

// NewClient create Client
func NewClient() *Client {
	return &Client{
		client: resty.New(),
	}
}

// SendMessage send message
func (c *Client) SendMessage(message Message) (*Response, error) {
	if message == nil {
		return nil, errors.New("message missing")
	}
	var webhook = c.Webhook
	if len(c.Secret) > 0 {
		timestamp := time.Now().UnixNano() / 1e6
		sign, err := c.GenSign(c.Secret, timestamp)
		if err != nil {
			return nil, err
		}
		webhook = webhook + "&" + c.urlParamsEncode(sign, timestamp)

	}
	return c.send(webhook, message.ToMessageMap())
}

// SendMessageByUrl send message custom url
func (c *Client) SendMessageByUrl(url string, message Message) (*Response, error) {
	if message == nil {
		return nil, errors.New("message missing")
	}
	return c.send(url, message.ToMessageMap())
}

// SendMessageStr send message json string
func (c *Client) SendMessageStr(json string) (*Response, error) {
	var webhook = c.Webhook
	if len(c.Secret) > 0 {
		timestamp := time.Now().UnixNano() / 1e6
		sign, err := c.GenSign(c.Secret, timestamp)
		if err != nil {
			return nil, err
		}
		webhook = webhook + "&" + c.urlParamsEncode(sign, timestamp)

	}
	return c.send(webhook, json)
}

// SendMessageStrByUrl send message custom url and json string message
func (c *Client) SendMessageStrByUrl(url, json string) (*Response, error) {
	return c.send(url, json)
}
func (c *Client) send(url string, body interface{}) (*Response, error) {
	resp, err := c.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&Response{}).
		ForceContentType("application/json").
		Post(url)
	if err != nil {
		return nil, err
	}
	result := resp.Result().(*Response)
	return result, nil
}

// GenSign sign secret
func (c *Client) GenSign(secret string, timestamp int64) (string, error) {
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, secret)
	mac := hmac.New(sha256.New, []byte(secret))
	_, err := mac.Write([]byte(stringToSign))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(mac.Sum(nil)), nil
}

// urlParamsEncode set url params encode
func (c *Client) urlParamsEncode(sign string, timestamp int64) string {
	value := url.Values{}
	value.Set("timestamp", strconv.FormatInt(timestamp, 10))
	value.Set("sign", sign)
	return value.Encode()
}

type Response struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}

func (r Response) Success() bool {
	return r.ErrorCode == 0
}
