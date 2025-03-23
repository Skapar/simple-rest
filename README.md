# Go Simple Rest 🚀

## 📌 Описание
Этот проект написан на **Go** и использует **GORM** для работы с базой данных PostgreSQL.  
Перед запуском необходимо создать файл `.env` и загрузить переменные окружения.

---

## 🛠️ Установка и запуск

### 1️⃣ Установите зависимости
```sh
go mod tidy
```

### 2️⃣ Установите `goose`
```sh
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### 3️⃣ Создайте файл `.env`  
Пример `.env` файла:
```sh
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=yourdb
DB_PORT=5432
```

### 4️⃣ Загрузите переменные окружения
```sh
source .env
```

### 5️⃣ Выполните миграции
```sh
make migrate-up
```

### 6️⃣ Запустите проект
```sh
go run main.go
```


---

## 📦 Используемые технологии
- **Go** (Golang)
- **Gin** (HTTP-фреймворк)
- **GORM** (ORM для работы с PostgreSQL)
- **PostgreSQL** (Реляционная база данных)
- **Goose** (Миграции базы данных)

---

## 🎯 TODO
- [ ] Добавить...

---
 

📌 **Пример коммита:**  
```sh
feat: для добавления новых фич  
fix: для исправления багов  
refactor: для изменения структуры кода без изменения функционала  
perf: для оптимизации производительности  
docs: для изменений в документации  
test: для добавления или обновления тестов  
ci: для изменений в CI/CD процессах  
chore: для обновлений зависимостей и вспомогательных задач  
build: для изменений, связанных со сборкой проекта или зависимостями  
init: для первоначальной инициализации проекта  
```

---

