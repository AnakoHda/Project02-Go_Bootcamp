package main

import "testing"

func Test_over_capacity(t *testing.T) {
	cache, err := NewCache[int, string](5)
	if err != nil {
		t.Fatalf("ошибка создания кэша: %s", err)
	}
	cache.Set(1, "one")
	cache.Set(2, "two")
	cache.Set(3, "tree")
	cache.Set(4, "four")
	cache.Set(5, "five")
	cache.Set(6, "six")
	if data, ok := cache.Get(1); ok {
		t.Fatalf("данные не удалились %s", data)
	}
}

func Test_unused_data(t *testing.T) {
	cache, err := NewCache[int, string](5)
	if err != nil {
		t.Fatalf("ошибка создания кэша: %s", err)
	}
	cache.Set(1, "one")
	cache.Set(2, "two")
	cache.Set(3, "tree")
	cache.Set(4, "four")
	cache.Set(5, "five")
	cache.Set(6, "six")
	if data, ok := cache.Get(1); ok {
		t.Fatalf("данные не удалились %s", data)
	}
	cache.Get(6)
	cache.Get(2)
	cache.Get(5)
	cache.Set(7, "seven")
	if data, ok := cache.Get(3); ok {
		t.Fatalf("данные не удалились %s", data)
	}
}
