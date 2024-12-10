package cache

import "testing"

func TestCache(t *testing.T) {
	c := New()

	// Проверяем метод Set
	c.Set("testKey", "testValue")
	value := c.Get("testKey")
	if value != "testValue" {
		t.Errorf("Ожидали 'testValue', но получили %v", value)
	}

	// Проверяем метод Delete
	c.Delete("testKey")
	value = c.Get("testKey")
	if value != nil {
		t.Errorf("Ожидали nil, но получили %v", value)
	}
}
