package civ_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ecshreve/civ-bot-go/internal/civ"
	"github.com/ecshreve/civ-bot-go/internal/constants"
)

func TestGetCivByString(t *testing.T) {
	civs := civ.GenCivs()
	civMap := civ.GenCivMap(civs)

	testcases := []struct {
		desc     string
		inp      string
		expected constants.CivKey
	}{
		{
			desc:     "exact match",
			inp:      "america",
			expected: constants.AMERICA,
		},
		{
			desc:     "exact match mixed capitalization",
			inp:      "AmerICa",
			expected: constants.AMERICA,
		},
		{
			desc:     "exact match leader",
			inp:      "washington",
			expected: constants.AMERICA,
		},
		{
			desc:     "exact match leader mixed capitalization",
			inp:      "WasHINGton",
			expected: constants.AMERICA,
		},
		{
			desc:     "civ misspelled a little bit",
			inp:      "amearica",
			expected: constants.AMERICA,
		},
		{
			desc:     "civ misspelled a lot",
			inp:      "AmMericaas",
			expected: constants.AMERICA,
		},
		{
			desc:     "leader misspelled a little bit",
			inp:      "washhingten",
			expected: constants.AMERICA,
		},
		{
			desc:     "leader misspelled a lot",
			inp:      "WashinnSSHton",
			expected: constants.AMERICA,
		},
		{
			desc:     "civ short substring",
			inp:      "meric",
			expected: constants.AMERICA,
		},
		{
			desc:     "leader short substring",
			inp:      "wash",
			expected: constants.AMERICA,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.desc, func(t *testing.T) {
			expectedCiv := civMap[testcase.expected]
			require.NotNil(t, expectedCiv)

			actualCiv := civ.GetCivByString(testcase.inp, civs)
			assert.Equal(t, expectedCiv, actualCiv)
		})
	}
}

func TestGenCivMap(t *testing.T) {
	civs := civ.GenCivs()
	civMap := civ.GenCivMap(civs)

	// Make sure that the values in our CivMap point to the same Civs as the
	// items in the slice passed in to GenCivMap.
	for _, c := range civs {
		assert.Same(t, c, civMap[c.Key])
	}
}

func TestSortCivs(t *testing.T) {
	civs := civ.GenCivs()

	testcases := []struct {
		description string
		input       []*civ.Civ
		expected    []*civ.Civ
	}{
		{
			description: "sorting nil slice returns nil",
			input:       nil,
			expected:    nil,
		},
		{
			description: "sorting slice of length 1 returns that civ",
			input:       []*civ.Civ{civs[0]},
			expected:    []*civ.Civ{civs[0]},
		},
		{
			description: "sorting slice of two out of order civs returns them in order",
			input:       []*civ.Civ{civs[1], civs[0]},
			expected:    []*civ.Civ{civs[0], civs[1]},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			civ.SortCivs(testcase.input)
			assert.Equal(t, testcase.expected, testcase.input)
		})
	}
}
