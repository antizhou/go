package util

import (
	"regexp"
	"sync"
)

var (
	regexpCacheMu sync.Mutex
	regexpCache   = make(map[string]*regexp.Regexp)
)

func GetRegexpFromCache(exp string) (*regexp.Regexp, error) {
	regexpCacheMu.Lock()
	defer regexpCacheMu.Unlock()

	r, e := regexpCache[exp]

	if e {
		return r, nil
	}

	r, err := regexp.Compile(exp)

	if err != nil {
		return nil, err
	}

	regexpCache[exp] = r

	return r, nil
}
