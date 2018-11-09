package models

type ActivityWrapper struct {
	TimestampMs string     `json:"timestampMs"`
	Activities  []Activity `json:"activity"`
}

type Activity struct {
	Type       string `json:"type"`
	Confidence int    `json:"confidence"`
}
