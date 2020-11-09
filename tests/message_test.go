package tests

import (
	"os"
	"testing"

	"github.com/872409/go-netease-im"
)

type MSI = map[string]interface{}

func TestSendTextMessage(t *testing.T) {
	ms := MSI{"aa": "aa"}
	t.Log(ms)

	msg := &netease.TextMessage{Message: "欢迎使用KT\n\n    使用KT，您将获得安全，快速，免费的加密消息传递和加密的语音/视频通话。\n    服务器将永远不会访问您的任何通信信息，也永远不会存储您的任何数据。 \n    您可以同时使用此帐户登录其他Mac，Windows设备，并且消息将实时同步。\n\n温馨提示：\n     KT平台没有任何第三方团队进行任何形式的推广、营销等活动，请用户在使用过程中，注意核实身份，谨防诈骗。 如在使用过程中遭遇刷单，网赚，网赌、电信等诈骗，因此遭受个人财产损失，平台不会承担任何责任！\n    对于某些用户和群的恶意诈骗行为，可在用户详情和群详情中进行举报，我们会及时进行处理。请大家共同监督，营造一个干净，良好的沟通平台。\n\nKT 团队"}
	err := client.SendTextMessage("100", "335014782736284399", msg, nil)
	err = client.SendTextMessage("100", "335012783647111919", msg, nil)
	// err := client.SendTextMessage("335014782736284399", "335014782736284399", msg, nil)
	// err := client.SendTextMessage("346464272269720303", "335012783647111919", msg, nil)
	// err := client.SendTextMessage("100", "346464272269720303", msg, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestSendTipMessage(t *testing.T) {
	msg := &netease.TextMessage{Message: "message test 1"}
	// err := client.SendTextMessage("100", "346464272269720303", msg, nil)
	err := client.SendTipMessage("100", "335014782736284399", msg, nil)
	// err := client.SendTextMessage("100", "335012783647111919", msg, nil)
	// err := client.SendTipMessage("100", "346464272269720303", msg, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestSendBatchTextMessage(t *testing.T) {
	msg := &netease.TextMessage{Message: "message test"}
	str, err := client.SendBatchTextMessage("1", []string{"169143"}, msg, nil)
	t.Log(str)
	if err != nil {
		t.Error(err)
	}
}

func TestSendBatchAttachMessage(t *testing.T) {
	err := client.SendBatchAttachMsg("1", "{'msg':'test'}", []string{"2", "3"}, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestBroadcastMsg(t *testing.T) {
	os.Setenv("GOCACHE", "off")
	t.Log(client.BroadcastMsg("好久不见了呢，我在这里等你哦", "", nil, nil))
}

func TestRecallMsg(t *testing.T) {
	err := client.RecallMessage("280384449779", "1559633306342", "test1", "test2", 7)
	if err != nil {
		t.Error(err)
	}
}
