package loop

import "fmt"

func Init() {
	loadTags()
}
func loadTags() {
	var unloadedTags []string = []string{
		"dup",
		"dup",
		"fine",
		"two",
	}

	var loadedTags []string
	fmt.Println("Initiating tags")
	tagLoop: for _, unloadedTag := range unloadedTags {
		for _, loadedTag := range loadedTags {
			if loadedTag == unloadedTag {
				fmt.Println("A duplicated tag is found: " + loadedTag)
				continue tagLoop
				
			}
		}

		fmt.Println("Loaded tag: " + unloadedTag)
		loadedTags = append(loadedTags, unloadedTag)
	}
}