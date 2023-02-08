package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once
var single *Single

func getInstance() *Single {
	if single == nil {
		once.Do(func() {
			single = &Single{}
			fmt.Println("Single has been created")
		})
	} else {
		fmt.Println("Single is already been created")
	}

	return single
}
