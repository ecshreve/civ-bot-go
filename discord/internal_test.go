package discord

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetCivByString(t *testing.T) {
	cs := NewCivSession()

	testcases := []struct {
		desc     string
		inp      string
		expected CivKey
	}{
		{
			desc:     "exact match",
			inp:      "america",
			expected: AMERICA,
		},
		{
			desc:     "exact match mixed capitalization",
			inp:      "AmerICa",
			expected: AMERICA,
		},
		{
			desc:     "exact match leader",
			inp:      "washington",
			expected: AMERICA,
		},
		{
			desc:     "exact match leader mixed capitalization",
			inp:      "WasHINGton",
			expected: AMERICA,
		},
		{
			desc:     "civ misspelled a little bit",
			inp:      "amearica",
			expected: AMERICA,
		},
		{
			desc:     "civ misspelled a lot",
			inp:      "AmMericaas",
			expected: AMERICA,
		},
		{
			desc:     "leader misspelled a little bit",
			inp:      "washhingten",
			expected: AMERICA,
		},
		{
			desc:     "leader misspelled a lot",
			inp:      "WashinnSSHton",
			expected: AMERICA,
		},
		{
			desc:     "civ short substring",
			inp:      "meric",
			expected: AMERICA,
		},
		{
			desc:     "leader short substring",
			inp:      "wash",
			expected: AMERICA,
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
