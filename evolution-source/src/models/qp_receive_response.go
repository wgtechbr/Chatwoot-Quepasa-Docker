package models

import (
	whatsapp "github.com/sufficit/sufficit-quepasa/whatsapp"
)

type QpReceiveResponse struct {
	QpResponse
	Total    uint64                     `json:"total"`
	Messages []whatsapp.WhatsappMessage `json:"messages,omitempty"`
	Server   *QpServer                  `json:"server,omitempty"`
}
