package pokecache

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	interval := 10 * time.Millisecond
	cache := NewCache(interval)

	cases := []struct {
		key string
		value []byte
	}{
		{
			key: "url-one",
			value: []byte("val-one"),
		},
		{
			key: "url-two",
			value: []byte("val-two"),
		},
		{
			key: "url-three",
			value: []byte("value-three"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v - Key: %s", i, c.key), func(t *testing.T) {
			cache.Add(c.key, c.value)
			entry, exists := cache.Get(c.key)

			if !exists {
				t.Errorf("data not found")
				return
			}

			if !bytes.Equal(c.value, entry) {
				t.Errorf("data mismatch")
			}
		})
	}
}

func TestAdd_UpdateExisting(t *testing.T) {
	interval := 10 * time.Millisecond
	cache := NewCache(interval)

	cache.Add("test-url", []byte("test-data"))

	cache.Add("test-url", []byte("updated-test-data"))

	entry, exist := cache.Get("test-url")

	if bytes.Equal([]byte("test-data"), entry) {
		t.Errorf("data not updated")
	}

	if !exist {
		t.Errorf("updated data not found")
		return
	} 

	if !bytes.Equal([]byte("updated-test-data"), entry) {
		t.Errorf("updated data not found")
	}
}

func TestGet(t *testing.T) {
	interval := 10 * time.Millisecond
	cache := NewCache(interval)

	entry, exists := cache.Get("test-url")

	if exists {
		t.Error("no data should exist")
		return
	}

	if entry != nil {
		t.Errorf("no data should exists")
	}
}

func TestReapLoop(t *testing.T) {
	interval := 10 * time.Millisecond
	cache := NewCache(interval)

	cache.Add("test-url", []byte("test-data"))
	time.Sleep(15 * time.Millisecond)
	_, exists := cache.Get("test-url")

	if exists {
		t.Errorf("not removed")
	}
}

func TestReapLoop_NotReaped(t *testing.T) {
	interval := 10 * time.Millisecond
	cache := NewCache(interval)

	cache.Add("test-url", []byte("test-data"))

	time.Sleep(5 * time.Millisecond)

	_, exists := cache.Get("test-url")
	
	if !exists {
		t.Error("entry prematurely removed")
	}
}