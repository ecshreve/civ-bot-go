package bot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddHandlers(t *testing.T) {
	b, mock := MockBot(t)

	t.Run("add handlers to intialized bot", func(t *testing.T) {
		mock.Expect(b.DS.AddHandler, b.MessageHandler)
		mock.Expect(b.DS.AddHandler, b.ReactionHandler)
		b.AddHandlers()
		Check(mock.Check(), true, t)
	})

	t.Run("add handlers to bot with nil discordgo Session, expect error", func(t *testing.T) {
		b.DS = nil
		err := b.AddHandlers()
		assert.Error(t, err)
	})
}

func TestStartSession(t *testing.T) {
	b, mock := MockBot(t)

	mock.Expect(b.DS.Open)
	b.StartSession()
	Check(mock.Check(), true, t)
}

func TestEndSession(t *testing.T) {
	b, mock := MockBot(t)

	mock.Expect(b.DS.Close)
	b.EndSession()
	Check(mock.Check(), true, t)
}
