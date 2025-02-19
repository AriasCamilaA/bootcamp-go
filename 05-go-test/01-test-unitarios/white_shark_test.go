package hunt_test

import (
	hunt "testdoubles"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for the WhiteShark implementation - Hunt method
func TestWhiteSharkHunt(t *testing.T) {
	t.Run("case 1: white shark hunts successfully", func(t *testing.T) {
		// ARRANGE
		shark := hunt.NewWhiteShark(true, false, 100)
		tuna := hunt.NewTuna("Lovest tuna", 30)
		// ACT
		err := shark.Hunt(tuna)
		// ASSRT
		require.NoError(t, err)
		require.False(t, shark.Hungry)
		require.True(t, shark.Tired)
	})

	t.Run("case 2: white shark is not hungry", func(t *testing.T) {
		// ARRANGE
		shark := hunt.NewWhiteShark(false, false, 100)
		tuna := hunt.NewTuna("Lovest tuna", 30)
		ErrExpected := hunt.ErrSharkIsNotHungry
		// ACT
		err := shark.Hunt(tuna)
		// ASSRT
		require.ErrorIs(t, err, ErrExpected)
		require.False(t, shark.Hungry)
		require.False(t, shark.Tired)
	})

	t.Run("case 3: white shark is tired", func(t *testing.T) {
		// ARRANGE
		shark := hunt.NewWhiteShark(true, true, 100)
		tuna := hunt.NewTuna("Lovest tuna", 30)
		ErrExpected := hunt.ErrSharkIsTired
		// ACT
		err := shark.Hunt(tuna)
		// ASSRT
		require.ErrorIs(t, err, ErrExpected)
		require.True(t, shark.Hungry)
		require.True(t, shark.Tired)
	})

	t.Run("case 4: white shark is slower than the tuna", func(t *testing.T) {
		// ARRANGE
		shark := hunt.NewWhiteShark(true, false, 100)
		tuna := hunt.NewTuna("Fastest tuna", 110)
		ErrExpected := hunt.ErrSharkIsSlower
		// ACT
		err := shark.Hunt(tuna)
		// ASSRT
		require.ErrorIs(t, err, ErrExpected)
		require.True(t, shark.Hungry)
		require.False(t, shark.Tired)
	})

	t.Run("case 5: tuna is nil", func(t *testing.T) {
		// ARRANGE
		shark := hunt.NewWhiteShark(true, false, 100)
		ErrExpected := hunt.ErrTunnaIsEmpty
		// ACT
		err := shark.Hunt(nil)
		// ASSRT
		require.ErrorIs(t, err, ErrExpected)
		require.True(t, shark.Hungry)
		require.False(t, shark.Tired)
	})
}
