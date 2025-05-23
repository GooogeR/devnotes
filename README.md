# DevNotes

DevNotes — это простое приложение для создания, регистрации и управления заметками, используя фреймворк Gin и Golang. Приложение позволяет пользователям регистрироваться, входить в систему и создавать заметки, которые сохраняются в памяти.

## Описание проекта

Это приложение состоит из следующих компонентов:
- Регистрация пользователей
- Аутентификация
- Создание и хранение заметок

Приложение использует фреймворк Gin для обработки HTTP-запросов и реализации маршрутов.

## Установка

Для запуска проекта выполните следующие шаги:

1. Клонируйте репозиторий:

    ```bash
    git clone https://github.com/your-username/devnotes.git
    cd devnotes
    ```

2. Убедитесь, что у вас установлен Go (версия 1.18 или выше).

3. Запустите проект:

    ```bash
    go run main.go
    ```

4. Приложение будет доступно по адресу [http://localhost:8080](http://localhost:8080).

## Маршруты

### 1. Регистрация пользователя

- **Маршрут**: `POST /register`
- **Описание**: Регистрация нового пользователя.
- **Тело запроса**:
    ```json
    {
        "username": "your-username",
        "password": "your-password"
    }
    ```

- **Ответ**:
    ```json
    {
        "id": "user-id",
        "username": "your-username"
    }
    ```

### 2. Вход пользователя

- **Маршрут**: `POST /login`
- **Описание**: Вход в систему.
- **Тело запроса**:
    ```json
    {
        "username": "your-username",
        "password": "your-password"
    }
    ```

- **Ответ**:
    ```json
    {
        "message": "Login successful",
        "user_id": "user-id"
    }
    ```

### 3. Создание заметки

- **Маршрут**: `POST /notes`
- **Описание**: Создание новой заметки.
- **Тело запроса**:
    ```json
    {
        "title": "Note Title",
        "content": "Note Content"
    }
    ```

- **Ответ**:
    ```json
    {
        "id": "note-id",
        "title": "Note Title",
        "content": "Note Content"
    }
    ```

### 4. Получение заметок

- **Маршрут**: `GET /notes`
- **Описание**: Получение списка всех заметок пользователя.
- **Ответ**:
    ```json
    [
        {
            "id": "note-id",
            "title": "Note Title",
            "content": "Note Content"
        }
    ]
    ```

## Технологии

- **Golang** — для реализации серверной логики.
- **Gin** — для маршрутизации и обработки HTTP-запросов.
- **Memory Store** — для хранения данных в памяти.
