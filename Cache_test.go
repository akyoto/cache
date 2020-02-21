package cache_test

import (
	"testing"
	"time"

	"github.com/akyoto/cache"
)

func TestGetSet(t *testing.T) {
	cycle := 100 * time.Millisecond
	c := cache.New(cycle)
	defer c.Close()

	c.Set("sticky", "forever", 0)
	c.Set("hello", "Hello", cycle/2)
	hello, found := c.Get("hello")

	if !found {
		t.FailNow()
	}

	if hello.(string) != "Hello" {
		t.FailNow()
	}

	time.Sleep(cycle / 2)

	_, found = c.Get("hello")

	if found {
		t.FailNow()
	}

	time.Sleep(cycle)

	_, found = c.Get("404")

	if found {
		t.FailNow()
	}
	
	_, found = c.Get("sticky")
	if !found {
		t.FailNow()
	}
}

func TestDelete(t *testing.T) {
	c := cache.New(time.Minute)
	c.Set("hello", "Hello", time.Hour)
	_, found := c.Get("hello")

	if !found {
		t.FailNow()
	}

	c.Delete("hello")

	_, found = c.Get("hello")

	if found {
		t.FailNow()
	}
}

func TestRange(t *testing.T) {
	c := cache.New(time.Minute)
	c.Set("hello", "Hello", time.Hour)
	c.Set("world", "World", time.Hour)
	count := 0

	c.Range(func(key, value interface{}) bool {
		count++
		return true
	})

	if count != 2 {
		t.FailNow()
	}
}

func TestRangeTimer(t *testing.T) {
	c := cache.New(time.Minute)
	c.Set("message", "Hello", time.Nanosecond)
	c.Set("world", "World", time.Nanosecond)
	time.Sleep(time.Microsecond)

	c.Range(func(key, value interface{}) bool {
		t.FailNow()
		return true
	})
}

func BenchmarkNew(b *testing.B) {
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cache.New(5 * time.Second).Close()
		}
	})
}

func BenchmarkGet(b *testing.B) {
	c := cache.New(5 * time.Second)
	defer c.Close()
	c.Set("Hello", "World", 0)

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Get("Hello")
		}
	})
}

func BenchmarkSet(b *testing.B) {
	c := cache.New(5 * time.Second)
	defer c.Close()

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Set("Hello", "World", 0)
		}
	})
}

func BenchmarkDelete(b *testing.B) {
	c := cache.New(5 * time.Second)
	defer c.Close()

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Delete("Hello")
		}
	})
}
