package main

import (
	"testing"
	"time"
)

func TestTTLCache(t *testing.T) {
	t.Run("set and get", func(t *testing.T) {
		cache := NewTTLCache()
		cache.Set("key1", "value1", time.Hour)

		value, exists := cache.Get("key1")
		if !exists {
			t.Error("expected key to exist")
		}
		if value != "value1" {
			t.Errorf("expected value1, got %v", value)
		}
	})

	t.Run("get non-existent key", func(t *testing.T) {
		cache := NewTTLCache()
		_, exists := cache.Get("nonexistent")
		if exists {
			t.Error("expected key to not exist")
		}
	})

	t.Run("expiration", func(t *testing.T) {
		cache := NewTTLCache()
		cache.Set("key1", "value1", time.Millisecond*500)

		// Wait for expiration
		time.Sleep(time.Millisecond * 1000)

		_, exists := cache.Get("key1")
		if exists {
			t.Error("expected key to be expired")
		}
	})

	t.Run("update existing key", func(t *testing.T) {
		cache := NewTTLCache()
		cache.Set("key1", "value1", time.Hour)
		cache.Set("key1", "value2", time.Hour)

		value, exists := cache.Get("key1")
		if !exists {
			t.Error("expected key to exist")
		}
		if value != "value2" {
			t.Errorf("expected value2, got %v", value)
		}
	})

	t.Run("delete", func(t *testing.T) {
		cache := NewTTLCache()
		cache.Set("key1", "value1", time.Hour)
		cache.Delete("key1")

		_, exists := cache.Get("key1")
		if exists {
			t.Error("expected key to not exist after deletion")
		}
	})

	t.Run("concurrent access", func(t *testing.T) {
		cache := NewTTLCache()
		done := make(chan bool)

		go func() {
			for i := 0; i < 100; i++ {
				cache.Set("key", i, time.Hour)
			}
			done <- true
		}()

		go func() {
			for i := 0; i < 100; i++ {
				cache.Get("key")
			}
			done <- true
		}()

		<-done
		<-done
	})
}
