package strategy

import "fmt"

// Fifo 先进先出
type Fifo struct {
}

func (l *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by fifo strategy")
}
