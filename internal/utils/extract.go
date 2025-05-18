package utils

func ExtractFrom[ItemType any, T any](s []ItemType, f func(ItemType) (bool, T)) []T {
	var res []T
	var ok bool
	var i T

	for _, item := range s {
		ok, i = f(item)
		if ok {
			res = append(res, i)
		}
	}

	return res
}
