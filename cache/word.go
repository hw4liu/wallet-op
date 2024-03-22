package cache

import (
	"bufio"
	"os"
)

type Caches map[string]struct{}

func (c Caches) List() (list []string) {
	for k := range c {
		list = append(list, k)
	}
	return
}

var word Caches

func init() {
	word = make(Caches)
}

func InitWordCache(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word[scanner.Text()] = struct{}{}
	}
	return nil
}

func ListWord() []string {
	return word.List()
}
