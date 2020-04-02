package netease

import (
	"errors"

	jsoniter "github.com/json-iterator/go"
)

func (c *ImClient) AddFriend(accid string, faccid string) error {
	param := map[string]string{"accid": accid, "faccid": faccid, "type": "1"}
	_, err := c.post(friendAdd, param, "info")
	return err
}

func (c *ImClient) SetFriend(accid string, faccid string, alias string, ex string, serverex string) error {
	param := map[string]string{"accid": accid, "faccid": faccid, "alias": alias, "ex": ex, "serverex": serverex}
	_, err := c.post(friendUpdate, param, "info")
	return err
}

func (c *ImClient) GetFriends(accid string, updatetime string) (*[]Friend, error) {
	if len(accid) == 0 {
		return nil, errors.New("必须指定网易云通信ID")
	}

	param := map[string]string{"accid": accid, "updatetime": updatetime}
	// client := c.client.R()
	// c.setCommonHead(client)
	// client.SetFormData(param)

	infoJSON, err := c.post(friendGet, param, "friends")

	// infoJSON, err := handleResp(client.Post(refreshTokenPoint))

	if err != nil || infoJSON == nil {
		return nil, err
	}

	var friends = &[]Friend{}
	err = jsoniter.Unmarshal(*infoJSON, friends)
	if err != nil {
		return nil, err
	}

	return friends, nil
}
