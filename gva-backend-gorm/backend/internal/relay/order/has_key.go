package order

import "backend/internal/relay/utils"

func HasKey(input any, key string) (bool, error) {
	filter, err := utils.ConvertToMapString(input)
	if err != nil {
		return false, err
	}

	for k := range filter {
		if k == key {
			return true, nil
		}
	}

	return false, nil
}
