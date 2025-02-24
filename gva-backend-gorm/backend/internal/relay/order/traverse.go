package order

import "backend/internal/relay/utils"

func traverse(table string, tables *map[string]string, orderMap map[string]string, reverse bool) (orders []string) {
	orders = make([]string, len(orderMap))
	i := int64(0)
	for key, direction := range orderMap {
		idx, field := utils.SplitOrderKey(key)
		if idx < 0 {
			idx = i
		}
		if reverse {
			direction = utils.ReverseDirection(direction)
		}

		prefix := ""

		if table != "" {
			prefix = table + "."
		}

		if tables != nil {
			for k, v := range *tables {
				if k == field {
					prefix = v + "."
					break
				}
			}
		}

		orders[idx] = prefix + field + " " + direction
		i++
	}
	return
}
