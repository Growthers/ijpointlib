package ijpoint

import (
	"strings"
)

/*
 文字列からデータセットに含まれる単語が存在するか
*/

func SearchWords(sentence string, dataset IjDataSet) (float64, []string) {
	var point float64
	hitString := make([]string, 0)

	for _, v := range dataset.Words {
		if strings.Contains(sentence, v.Word) {
			point += v.Weight
			hitString = append(hitString, v.Word)
		}
	}
	return point, hitString
}
