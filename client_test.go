package dingtalk

import (
	"os"
	"testing"
)

func TestClient_SendMessageStr(t *testing.T) {
	textMsg := `{"at":{"atMobiles":["180xxxxxx"],"atUserIds":["user123"],"isAtAll":false},"text":{"content":"我就是我, @XXX 是不一样的烟火"},"msgtype":"text"}`
	client := NewClient()
	client.Webhook = os.Getenv("webhook")
	client.Secret = os.Getenv("secret")
	res, err := client.SendMessageStr(textMsg)
	if err != nil {
		t.Error("send message error", err)
	}
	if res.Success() {
		t.Log("send message success")
	} else {
		t.Errorf("send message error: code: %d  msg: %s", res.ErrorCode, res.ErrorMessage)
	}
}
