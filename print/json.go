package print

import (
	"encoding/json"
	"fmt"
)

func JSON(v any) {
	d, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(d))
}
