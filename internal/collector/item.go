package collector

import "strings"

type Item map[string]string

func NewEmptyItem() Item {
	return make(Item)
}

func NewItem(str string) Item {
	it := NewEmptyItem()
	rows := strings.Split(str, "\n")
	for _, row := range rows {
		row = strings.Trim(row, " ")
		if len(row) == 0 {
			continue
		}

		it.ParseLine(row)
	}

	return it
}

func (it Item) ParseLine(str string) {
	keyValue := strings.SplitN(str, "=", 2)
	key := keyValue[0]
	var value string
	if len(keyValue) == 1 {
		value = ""
	} else {
		value = keyValue[1]
	}

	it[key] = value
}

func (it Item) Has(key string) bool {
	_, found := it[key]

	return found
}

func (it Item) Get(key string) string {
	return it[key]
}

/**
0: emerg.
1: alert.
2: crit.
3: err.
4: warning.
5: notice.
6: info.
7: debug.
*/
