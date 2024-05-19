package auth

import "tronglv_upload_svc/helper/util/token"

func ParseClient(t, publicKey string) (ClientData, error) {
	client, err := token.NewTokenParser(token.WithPublicKey(publicKey)).Parse(t)
	if err != nil {
		return nil, err
	}
	return &clientData{
		Id:     client.GetInt("id"),
		Uid:    client.GetString("sub"),
		Name:   client.GetString("name"),
		Scopes: client.GetSliceString("scp"),
	}, nil
}

type ClientData interface {
	GetId() int32
	GetUid() string
	GetName() string
	GetScopes() []string
}

type clientData struct {
	Id     int      `json:"id,omitempty"`
	Uid    string   `json:"uid,omitempty"`
	Name   string   `json:"name,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
}

func (c *clientData) GetId() int32 {
	return int32(c.Id)
}

func (c *clientData) GetUid() string {
	return c.Uid
}

func (c *clientData) GetName() string {
	return c.Name
}

func (c *clientData) GetScopes() []string {
	return c.Scopes
}
