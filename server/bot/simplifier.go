package bot

import (
	"context"
	"time"

	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types/events"
)

// sendText sends a text message to the chat from which the given message was received.
// It uses the provided events.Message to identify the chat and constructs a new waE2E.Message
// with the given text for sending. Returns an error if the message fails to send.
func sendText(m *events.Message, markRead bool, text string) error {
	if markRead {
		client.MarkRead([]string{m.Info.ID}, time.Now(), m.Info.Chat, m.Info.Sender)
	}
	_, err := client.SendMessage(context.Background(), m.Info.Chat, &waE2E.Message{
		Conversation: &text,
	})
	return err
}
