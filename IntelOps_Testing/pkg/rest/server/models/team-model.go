package models

type Team struct {
	Id int64 `json:"id,omitempty"`

	Bandwidth `json:"bandwidth,omitempty"`

	Length int8 `json:"length,omitempty"`

	Width int8 `json:"width,omitempty"`
}
