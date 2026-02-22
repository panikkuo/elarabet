Есть 2 микросервиса
1. Отвечает за работу с телеграммом. То есть запрос строить для осовного итд итп
2. Просто работает с бд. Начнем с него. Ему похуй на все он знает только про себя и про себя отвечает.


Posgres:
- users
Redis:
- notes
- refresh-tokens

back-core:
POST api/v1/signup
POST api/v1/login
GET api/v1/users/{id}
GET api/v1/notes
POST api/v1/notes

