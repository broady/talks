package main

import (
	"fmt"
	"runtime"
	"sort"
)

func main() {
	// START OMIT
	versions := []string{"go1", runtime.Version(), "go1.1", "go1.5", "go1.3", "go1.7"}

	sort.Slice(versions, func(i, j int) bool {
		return versions[i] > versions[j]
	})

	fmt.Println(versions[0])
	// END OMIT
}
