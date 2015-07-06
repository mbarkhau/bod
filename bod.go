package bod

import "github.com/hashicorp/golang-lru"

type Manager struct {
	pickles *lru.Cache
}

func onSliceEvicted(k interface{}, v interface{}) {
	//p := v.(*pickle)
}
