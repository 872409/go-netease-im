package tests

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/872409/go-netease-im"
)

var client = netease.CreateImClient("d36cd5f37a0657f49b46f00cae7505f7", "570d1c015865", "")

func init() {
	os.Setenv("GOCACHE", "off")
}

func TestToken(t *testing.T) {
	user := &netease.ImUser{Accid: "test2", Name: "test3", Gender: 1}
	tk, err := client.CreateImUser(user)
	if err != nil {
		t.Error(err)
	}
	t.Log(tk)
}
func TestUpdateToken(t *testing.T) {
	err := client.UpdateToken("319602531829285237", "9508f86d3266c8900662f0ac3f521ec1", "")
	t.Log(err)
}

func TestRefreshToken(t *testing.T) {
	tk, err := client.RefreshToken("319602531829285237")
	if err != nil {
		t.Error(err)
	}
	b, err := json.Marshal(tk)
	t.Log(string(b), err)
}

func Benchmark_SyncMap(b *testing.B) {
	netease.CreateImClient("", "", "")
}
