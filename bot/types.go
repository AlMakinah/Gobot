package bot

// This files has all the structs needed to decode the requests sent
// by facebook on the webhook api

// Check the documentation for the Facebook Messenger Webhook API

type botEvent struct {
	Object  string     `json:"object"`
	Entries []botEntry `json:"entry"`
}

func (b *botEvent) String() string {
	return b.Entries[0].Messaging[0].Message.Text
}

type botEntry struct {
	Id        string         `json:"id"`
	Time      int            `json:"time"`
	Messaging []botMessaging `json:"messaging"`
}

type botMessaging struct {
	Sender    botPerson `json:"sender"`
	Recipient botPerson `json:"recipient"`

	Timestamp int        `json:"timestamp"`
	Message   botMessage `json:"message"`
}

type botPerson struct {
	Id string `json:"id"`
}

type botMessage struct {
	Mid  string `json:"mid"`
	Seq  int    `json:"seq"`
	Text string `json:"text"`
}
