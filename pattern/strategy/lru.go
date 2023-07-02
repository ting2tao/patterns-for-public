package strategy

import "fmt"

// Lru lru 最近最少使用
type Lru struct {
}

func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting by lru strategy")
}
