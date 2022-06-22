package ijpoint

import (
	"encoding/json"
	"io/ioutil"
)

/*

方針:
- データセット
	- 単語を保存しておく
	- 重みは0~1(float64)
	- 重みの初期値は0.50

*/

type IjDataSet struct {
	Words []WordData `json:"words"`
}

type WordData struct {
	Word   string  `json:"word"`
	Weight float64 `json:"weight"`
}

func ReadIjDataSet(name string) (IjDataSet, error) {
	file, err := ioutil.ReadFile(name)
	if err != nil {
		return IjDataSet{}, err
	}

	dataSet := IjDataSet{}
	err = json.Unmarshal(file, &dataSet)
	if err != nil {
		return IjDataSet{}, err
	}
	return dataSet, nil
}

func UpdateIjDataSet(name string, dataset IjDataSet) error {
	datajson, err := json.Marshal(dataset)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(name, datajson, 0600)
	if err != nil {
		return err
	}

	return nil
}
