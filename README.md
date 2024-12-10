# In-memory Cache

Этот пакет реализует in-memory кэш с методами для записи, чтения и удаления данных.

## Установка

```bash
go get -u github.com/your-username/myproject/cache
go 

Вот как использовать этот кэш:

go
Копировать код
package main

import (
    "fmt"
    "github.com/твоя_учётка/cache"
)

func main() {
    cache := cache.New()

    cache.Set("userId", 42)
    fmt.Println(cache.Get("userId")) // Вывод: 42

    cache.Delete("userId")
    fmt.Println(cache.Get("userId")) // Вывод: <nil>
}

Чтобы протестировать этот пакет, выполните следующую команду в терминале:

bash
Копировать код
go test ./...