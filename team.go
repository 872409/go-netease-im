package netease

import (
	"errors"
	"strconv"

	jsoniter "github.com/json-iterator/go"
)

//
func (c *ImClient) TeamCreate(tname string, owner string, members []string, announcement string) (string, error) {
	b, _ := jsoniter.Marshal(members)
	membersStr := string(b)
	param := map[string]string{"tname": tname, "owner": owner, "members": membersStr, "announcement": announcement, "magree": "1", "joinmode": "0", "msg": ""}
	infoJSON, err := c.post(teamCreate, param, "tid")
	var result string
	err = jsoniter.Unmarshal(*infoJSON, result)

	return result, err
}

//
func (c *ImClient) TeamQuery(tids []string, isWidthMembers bool, ignoreInvalid bool) ([]TeamDetail, error) {
	b, _ := jsoniter.Marshal(tids)
	tidsStr := string(b)

	var ope string
	if isWidthMembers {
		ope = "1"
	} else {
		ope = "0"
	}

	param := map[string]string{"tids": tidsStr, "ope": ope, "ignoreInvalid": strconv.FormatBool(ignoreInvalid)}
	infoJSON, err := c.post(teamQuery, param, "tinfos")
	if err != nil || infoJSON == nil {
		return nil, err
	}

	var result = &[]TeamDetail{}
	err = jsoniter.Unmarshal(*infoJSON, result)
	if err != nil {
		return nil, err
	}

	return *result, nil
}

//
// func (c *ImClient) SetFriend(accid string, faccid string, alias string, ex string, serverex string) error {
// 	param := map[string]string{"accid": accid, "faccid": faccid, "alias": alias, "ex": ex, "serverex": serverex}
// 	_, err := c.post(friendUpdate, param, "info")
// 	return err
// }

func (c *ImClient) GetJoinTeams(accid string) (*[]Team, error) {
	if len(accid) == 0 {
		return nil, errors.New("必须指定网易云通信ID")
	}

	param := map[string]string{"accid": accid}

	infoJSON, err := c.post(teamJoinGet, param, "infos")

	if err != nil || infoJSON == nil {
		return nil, err
	}

	var result = &[]Team{}
	err = jsoniter.Unmarshal(*infoJSON, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
