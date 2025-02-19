package tickets_test

import (
	"testing"

	"github.com/AriasCamilaMeli/go-bases/internal/tickets"
	"github.com/stretchr/testify/require"
)

func TestGetTotalTickets(t *testing.T) {
	t.Run("Total ticket to Colombia", func(t *testing.T) {
		total, err := tickets.GetTotalTickets("Colombia")
		require.NoError(t, err)
		require.Equal(t, 18, total)
	})
	t.Run("Total ticket to Russia", func(t *testing.T) {
		total, err := tickets.GetTotalTickets("Russia")
		require.NoError(t, err)
		require.Equal(t, 41, total)
	})

}

func TestGetMornings(t *testing.T) {
	t.Run("Total ticket in the morning", func(t *testing.T) {

		total, err := tickets.GetMornings("mañana")
		require.NoError(t, err)
		require.Equal(t, 223, total)
	})

	t.Run("Total ticket in the afternoon", func(t *testing.T) {

		total, err := tickets.GetMornings("tarde")
		require.NoError(t, err)
		require.Equal(t, 249, total)
	})
}
func TestAverageDestination(t *testing.T) {
	t.Run("Average ticket to Colombia", func(t *testing.T) {

		average, err := tickets.AverageDestination("Colombia", 10)
		require.NoError(t, err)
		require.Equal(t, 0, average)
	})

	t.Run("Average ticket to Russia", func(t *testing.T) {

		average, err := tickets.AverageDestination("Russia", 10)
		require.NoError(t, err)
		require.Equal(t, 0, average)
	})
}
