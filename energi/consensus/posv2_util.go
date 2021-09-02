package consensus

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type PoSDiffV2TestCase struct {
	Time       uint64
	Parent     int64
	Drift      int64
	Integral   int64
	Derivative int64
	Result     int64
}

func readJsonPoSV2TestCases() (result []PoSDiffV2TestCase) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile(dir + "/intervalgen/PoSV2_test_cases.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		panic(err)
	}

	return
}
