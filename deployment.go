package lolo

import (
	"fmt"
)

type Deployment struct {
	Id       string `json:"id"`
	Version  int64 `json:"version"`
	Replicas int16 `json:"replicas"`
}

func (client *Client) GetDeployment(appId, runtimeId string) (*Deployment, error) {
	var deployment Deployment
	url := fmt.Sprintf("/apps/%s/deploy?runtimeId=%s", appId, runtimeId)

	if err := client.sendRequest("GET", url, nil, &deployment); err != nil {
		return nil, err
	}

	return &deployment, nil
}
