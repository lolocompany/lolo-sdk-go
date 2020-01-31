package lolo

import (
	"net/url"
	"strconv"
)

type Node struct {
	Id string `json:"id"`
}

type Edge struct {
	Id string `json:"id"`
}

type App struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	//Nodes       []Node `json:"nodes,omitempty"`
	//Edges       []Edge `json:"edges,omitempty"`
}

type AppList struct {
	Apps []App `json:"apps"`
}

func (client *Client) CreateApp(app *App) error {
	if err := client.sendRequest("POST", "/apps", *app, app); err != nil {
		return err
	}
	return nil
}

func (client *Client) UpdateApp(app *App) error {
	id := (*app).Id
	if err := client.sendRequest("PATCH", "/apps/"+id, *app, app); err != nil {
		return err
	}
	return nil
}

func (client *Client) GetApps(limit int) (*AppList, error) {
	var appList AppList

	base, _ := url.Parse("/apps")
	params := url.Values{}
	params.Add("limit", strconv.Itoa(limit))
	base.RawQuery = params.Encode()

	if err := client.sendRequest("GET", base.String(), nil, &appList); err != nil {
		return nil, err
	}

	return &appList, nil
}

func (client *Client) GetApp(id string) (*App, error) {
	var app App
	if err := client.sendRequest("GET", "/apps/"+id, nil, &app); err != nil {
		return nil, err
	}

	return &app, nil
}

func (client *Client) DeleteApp(id string) error {
	if err := client.sendRequest("DELETE", "/apps/"+id, nil, nil); err != nil {
		return err
	}

	return nil
}
