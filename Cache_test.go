package cache_test

import (
	"testing"
	"time"

	"github.com/akyoto/cache"
)

func Test(t *testing.T) {
	c := cache.New(20 * time.Millisecond)
	defer c.Close()

	c.Set("Hello", "World", 10*time.Millisecond)
	hello, found := c.Get("Hello")

	if !found {
		t.FailNow()
	}

	if hello.(string) != "World" {
		t.FailNow()
	}

	time.Sleep(10 * time.Millisecond)

	_, found = c.Get("Hello")

	if found {
		t.FailNow()
	}

	time.Sleep(20 * time.Millisecond)

	_, found = c.Get("404")

	if found {
		t.FailNow()
	}
}
