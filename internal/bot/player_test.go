package bot

import (
	"fmt"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/samsarahq/go/snapshotter"
	"github.com/stretchr/testify/assert"
)

func TestGetPlayerIDToPlayerMap(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	var testPlayers []*Player
	testUserIDs := []string{"testPlayer1", "testPlayer2"}
	for _, id := range testUserIDs {
		testUser := &discordgo.User{ID: id}
		testPlayer := NewPlayer(testUser)
		testPlayers = append(testPlayers, testPlayer)
	}

	playerIDToPlayerMap := GetPlayerIDToPlayerMap(testPlayers)
	snap.Snapshot("playerIdToPlayerMap", playerIDToPlayerMap)
}

func TestBanCiv(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	b, _ := MockBot(t)
	testUser := &discordgo.User{ID: "testPlayer"}
	testPlayer := NewPlayer(testUser)
	b.CivState.Players = append(b.CivState.Players, testPlayer)
	b.CivState.PlayerMap[testPlayer.PlayerID] = testPlayer

	testcases := []struct {
		description string
		civToBan    string
		expectError bool
	}{
		{
			description: "empty toBan arg",
			civToBan:    "",
			expectError: true,
		},
		{
			description: "toBan arg doesn't match any civs",
			civToBan:    "34523",
			expectError: true,
		},
		{
			description: "valid toBan arg",
			civToBan:    "america",
			expectError: false,
		},
		{
			description: "toBan arg is already banned",
			civToBan:    "america",
			expectError: true,
		},
		{
			description: "different valid toBan arg",
			civToBan:    "korea",
			expectError: false,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			civ, err := testPlayer.BanCiv(b, testcase.civToBan)

			if testcase.expectError {
				assert.Error(t, err)
				assert.Nil(t, civ)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, civ)
			}

			snap.Snapshot(fmt.Sprintf("%s - error", testcase.description), fmt.Sprintf("%v", err))
			snap.Snapshot(fmt.Sprintf("%s - bans", testcase.description), b.CivState.Bans)
		})
	}
}
