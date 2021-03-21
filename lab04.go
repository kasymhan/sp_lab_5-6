package main

import (
	"fmt"
	"regexp"
	"math"
	"strings"
	"unicode"
	"sync"
	"time"
)
// SplitOnNonLetters splits a string on non-letter runes
func SplitOnNonLetters(s string) []string {
	notALetter := func(char rune) bool { return !unicode.IsLetter(char)	}
	return strings.FieldsFunc(s, notALetter)
}
type Bigram struct {
	bigrams map[string]uint32
	mutex sync.Mutex
}

var str = "Мир - парк поезд. А куда ты! Мир парк"
func main() {
	fmt.Println(str)
	c := Bigram{bigrams: make(map[string]uint32) }
	words := regexp.MustCompile("[*.?!]{1} ").Split(str, -1)
	for _, word := range words {
		str = strings.ToLower(word)
		parts := SplitOnNonLetters(str)
		go  c.ngrams(parts, 2)
	}
	time.Sleep(time.Second)
	fmt.Println(c.Answer("answer"))
}
func (c *Bigram) Answer(key string) map[string]uint32 {
  c.mutex.Lock()
  defer c.mutex.Unlock()
  return c.bigrams
}

func (c *Bigram) ngrams(words []string, size int) {
	
	c.mutex.Lock()
	offset := int(math.Floor(float64(size / 2)))
	max := len(words)
	for i := range words {
		if i < offset || i+size-offset > max {
			continue
		}
		gram := strings.Join(words[i-offset:i+size-offset], " ")
		c.bigrams[gram]++
	}	
	c.mutex.Unlock()
}