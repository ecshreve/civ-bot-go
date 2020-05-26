package discord

import (
	"testing"

	"github.com/ecshreve/civ-bot-go/constants"
	"github.com/ecshreve/civ-bot-go/util"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetCivByString(t *testing.T) {
	cs := NewCivSession()

	testcases := []struct {
		desc     string
		inp      string
		expected util.CivKey
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
			var expectedCiv *Civ
			for _, c := range cs.Civs {
				if c.Key == testcase.expected {
					expectedCiv = c
				}
			}
			require.NotNil(t, expectedCiv)

			actualCiv := cs.getCivByString(testcase.inp)
			assert.Equal(t, expectedCiv, actualCiv)
		})
	}
}
