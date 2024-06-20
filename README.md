News API
=

Этот репозиторий содержит API на Go для управления новостными статьями и связанными с ними категориями.

Возможности:
-
1)Создание, обновление и удаление новостных статей
2)Присвоение категорий новостным статьям
3)Поддержка пагинации для списка новостей
4)Использование фреймворка Fiber для обработки HTTP-запросов
5)Fiber для работы с базой данных PostgreSQL

Требования
-
1)Go 1.22.2 или новее
2)PostgreSQL или MySQL
3)Fiber, Reform


Установка
-
1)Клонируйте репозиторий:
git clone https://github.com/Olegsuus/news-api.git


2)Установите зависимости:
Настройте подключение к базе данных в файлах config.yaml.

3)Запустите миграции базы данных:

goose -dir db/migrations postgres "user=yourusername dbname=yourdbname sslmode=disable" up

4)Соберите и запустите сервер:

go build -o server cmd/server/main.go
./server
API будет доступно по адресу http://localhost:8080.



API Эндпоинты
-
1)GET /list: Получить список новостей с пагинацией


2)POST /edit/:id: Редактировать новость


Пример:

curl -X POST http://localhost:8080/edit/1 -d '{"title": "Новое описание", "content": "Новое содержание статьи", "categories": [1, 2]}'

