## Что это?

Я попробую написать свои проблемы и решения сюда, чтобы можно было увидеть и прочитать, как я решил некоторые проблемы.

Мой уровень: C++ разработчик, довольно низкий уровень знаний для построения серверов: drogon/C++, posgreSQL

### Как работать с json в Go

#### Не через структруы

1. Создать словарь
2. json.Unmarshal()


### Как я решил хранить записки

Для начала не будем умничать и попробуем хранить записи в posgreSQL, записи будут иметь иерархическую структура что означает как-то надо будет хранить и знать кто отец


### Таблицы

Users
id
username
password
email
name

Notes
id
parent_id
user_id
text

Вообще что должно делать

Зарегистрироваться
Войти в систему
Посмотреть данные своего аккаунта
Заменить данные
Посмотреть за заметки
Войти в отца-заметку
Удалить заметки
Добавить заметку

### Endpoints

POST api/v1/signup

{
    username:
    password:
    email:
    name:
}
{
    id:
    username:
    name:
}

POST api/v1/login
{
    usename
    passowrd
}

GET api/v1/users/{id}
UPDATE api/v1/users/{id}
GET api/v1/notes
GET api/v1/notes?pid=
DELETE api/v1/{id}/notes?pid=
POST api/v1/{id}/new-notes?pid=