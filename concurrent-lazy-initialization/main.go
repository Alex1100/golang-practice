package main

import (
	"fmt"
	"sync"
)

/*
  Each call to Do(loadIcons) locks
  the mutex and checks the boolean
  variable. In the first call, in which
  the variable is false, Do calls
  loadIcons and sets the variable to
  true. Subsequent calls do nothing, but
  the mutex synchronization ensures
  that the effects of loadIcons
  on memory (specifically, icons)
  become visible to all goroutines.
  Using sync.Once in this way, we can
  avoid sharing variables with other
  goroutines until they have been
  properly constrcuted.
*/

var loadIconsOnce sync.Once
var icons map[string]string
var mu sync.Mutex

func loadIcon(name string) string {
	mu.Lock()
	defer mu.Unlock()
	return name
}

func loadIcons() {
	icons = make(map[string]string)
	icons["spade.png"] = loadIcon("spade.png")
	icons["hearts.png"] = loadIcon("hearts.png")
	icons["diamonds.png"] = loadIcon("diamond.png")
	icons["clubs.png"] = loadIcon("clubs.png")
}

func Icon(name string) string {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

func main() {
	imageList := make(chan string)
	images := [4]string{"spade.png", "hearts.png", "diamonds.png", "clubs.png"}

	for _, img := range images {
		go func(img string) {
			imageList <- Icon(img)
		}(img)
		fmt.Println("CURRENT IMAGE LOADING IS: ", <-imageList)
	}
}
