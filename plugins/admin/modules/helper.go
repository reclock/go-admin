package modules

import (
	"bytes"
	"encoding/gob"
	"github.com/NebulousLabs/fastrand"
	"strconv"
)

func InArray(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func InArrayWithoutEmpty(arr []string, str string) bool {
	if len(arr) == 0 {
		return true
	}
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func RemoveBlankFromArray(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func Uuid(length int64) string {
	ele := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "v", "k",
		"l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G",
		"H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	ele, _ = Random(ele)
	uuid := ""
	var i int64
	for i = 0; i < length; i++ {
		uuid += ele[fastrand.Intn(59)]
	}
	return uuid
}

func Random(strings []string) ([]string, error) {
	for i := len(strings) - 1; i > 0; i-- {
		num := fastrand.Intn(i + 1)
		strings[i], strings[num] = strings[num], strings[i]
	}

	str := make([]string, 0)
	for i := 0; i < len(strings); i++ {
		str = append(str, strings[i])
	}
	return str, nil
}

func SetDefault(source, def string) string {
	if source == "" {
		return def
	}
	return source
}

func GetPage(page string) (pageInt int) {
	if page == "" {
		pageInt = 1
	} else {
		pageInt, _ = strconv.Atoi(page)
	}
	return
}

func CopyMap(m map[string]string) map[string]string {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)
	err := enc.Encode(m)
	if err != nil {
		panic(err)
	}
	var cm map[string]string
	err = dec.Decode(&cm)
	if err != nil {
		panic(err)
	}
	return cm
}
