package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Webhook struct {
	URL string
}

func (w *Webhook) Post(m Message) error {
	data, err := json.Marshal(&m)
	if err != nil {
		return err
	}

	resp, err := http.Post(w.URL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}

	return nil
}

type Message struct {
	Text        string       `json:"text,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

type Attachment struct {
	Text     string  `json:"text,omitempty"`
	Fallback string  `json:"fallback,omitempty"`
	Color    string  `json:"color,omitempty"`
	Pretext  string  `json:"pretext,omitempty"`
	Title    string  `json:"title,omitempty"`
	Fields   []Field `json:"fields,omitempty"`
}

type Field struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	Short bool   `json:"short"`
}
