package gcloudFunctions

const SHOPPERS = 208
const LOBLAWS = 2018
const SOBEYS = 2072
const METRO = 2269

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Data []byte `json:"data"`
}
