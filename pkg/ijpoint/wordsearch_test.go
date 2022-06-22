package ijpoint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SearchWords(t *testing.T) {
	t.Parallel()

	d := IjDataSet{
		[]WordData{
			{
				Word:   "あいうおえ",
				Weight: float64(0.5),
			},
		},
	}
	req1, _ := SearchWords("あいうおえおえあおえ", d)
	assert.Equal(t, float64(0.5), req1)
	req2, _ := SearchWords("い", d)
	assert.NotEqual(t, float64(0.5), req2)
}
