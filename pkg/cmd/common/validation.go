package common

import (
	"errors"
	"github.sec.samsung.com/m5-kim/hdm/pkg/cmd/config"
	"strconv"
)

func ConvertToIndexTarget(target string) (int, error) {
	index, err := strconv.Atoi(target)
	var ret int
	if err == nil {
		ret = index
		return ret, nil
	}
	conf := config.GetConfig()
	for i, di := range conf.DeployInfo {
		for _, alias := range di.Alias {
			if alias == target {
				ret = i + 1
				return ret, nil
			}
		}
	}
	return -1, errors.New("NO TARGET EXISTS")
}
