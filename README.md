# Практическое задание № 3 Борисов Денис Александрович ЭФМО-01-25
Цели:
1.	Освоить базовую работу со стандартной библиотекой net/http без сторонних фреймворков.
2.	Научиться поднимать HTTP-сервер, настраивать маршрутизацию через http.ServeMux.
3.	Научиться обрабатывать параметры запроса (query, path), тело запроса (JSON/form-data) и формировать корректные ответы (код статуса, заголовки, JSON).
4.	Научиться базовому логированию запросов и обработке ошибок.


Выполнение работы:

1. Создаём скелет проекта
В ходе работы над практической работой были созданы следующий скелет проекта:
<img width="422" height="413" alt="image" src="https://github.com/user-attachments/assets/03aba111-78a5-447b-9aff-7aeb9e56b344" />


2. Первый сервер и маршрут GET /health

Посел был написан сервер и маршрут GET /health
Код для main.go:

<img width="426" height="853" alt="image" src="https://github.com/user-attachments/assets/d51ab5b7-2c20-4737-a46a-a04d2793097a" />

Проверка:

<img width="858" height="139" alt="image" src="https://github.com/user-attachments/assets/5acd1cf3-db1e-4e0d-95b5-950a271d24dc" />


3.  Память в приложении и модель сущности

Была создана простая сущность Task и in-memory хранилище.

<img width="286" height="709" alt="image" src="https://github.com/user-attachments/assets/b3acd298-c202-4bc3-bc5e-5fa7c4faa3fb" />

4.  Вспомогательные ответы и ошибки

Был написан код для вывода ответов и ошибок
<img width="530" height="485" alt="image" src="https://github.com/user-attachments/assets/898d4f67-da70-4542-84dc-10bb8a3a7762" />

5.  Обработчики: GET /tasks и POST /tasks

Были написан два обработчик GET /tasks и POST /tasks

Код для GET /tasks:

<img width="544" height="405" alt="image" src="https://github.com/user-attachments/assets/77ef4121-2d20-4acb-8de3-45e8bef92ee3" />

Код для POST /tasks:

<img width="746" height="367" alt="image" src="https://github.com/user-attachments/assets/9cc98ea4-eec3-4e1b-a3fe-fa5287cab43b" />

6.  Простое логирование (middleware-обёртка)

Было написано просто логирование в middleware
<img width="666" height="694" alt="image" src="https://github.com/user-attachments/assets/ab0eb9a5-276d-40b7-be32-6ec119a75b07" />

7.  Тестирование API

GET /health

<img width="565" height="451" alt="image" src="https://github.com/user-attachments/assets/c9aca20d-afd8-4aa2-ac01-ec2e87bc3e1c" />

POST /tasks (создание)

<img width="560" height="479" alt="image" src="https://github.com/user-attachments/assets/7d802982-2d47-4f69-96a5-48657c0e0e2c" />

GET /tasks (список)

<img width="571" height="503" alt="image" src="https://github.com/user-attachments/assets/1897a77c-a141-4564-adf8-4015a9f08e5b" />

GET /tasks/{id}

<img width="561" height="485" alt="image" src="https://github.com/user-attachments/assets/298d099b-5030-41fc-a685-eb4be21b73e6" />

Ошибки:
-	POST /tasks

<img width="565" height="474" alt="image" src="https://github.com/user-attachments/assets/e79f781e-169f-401c-bcff-9f5de95c61d6" />

-	GET /tasks/abc

<img width="563" height="460" alt="image" src="https://github.com/user-attachments/assets/35643fd5-a55b-44c3-84a1-4f446d793e74" />

-	GET /tasks/9999

<img width="572" height="460" alt="image" src="https://github.com/user-attachments/assets/b4eb73f8-c82e-4a2e-a815-3d3ababb540b" />

Дополнительно

8. CORS (минимально): добавить заголовки Access-Control-Allow-Origin: * для GET/POST

Для реализации CORS был дополнен код в middleware

<img width="620" height="264" alt="image" src="https://github.com/user-attachments/assets/ab44403e-f000-47a0-a73f-7de3035568c3" />

9. Валидация длины title (например, 1…140 символов).

Для реализации валидации была добавлена проверка в функция CreateTask

<img width="729" height="353" alt="image" src="https://github.com/user-attachments/assets/fe2c10f7-ebd9-4bbc-a139-3bdf7340fa17" />

Проверка
Успех <140

<img width="797" height="665" alt="image" src="https://github.com/user-attachments/assets/e5ec8b60-73e4-4394-a09f-2ef21394a2ab" />

Ошибка >140

<img width="776" height="627" alt="image" src="https://github.com/user-attachments/assets/96aea4bd-4d18-47c3-9053-641e2f4ff649" />

10. Метод PATCH /tasks/{id} для отметки Done=true.

В файле handlers.go был добавлена функция PatchTask

<img width="758" height="549" alt="image" src="https://github.com/user-attachments/assets/aa05f956-1337-482a-9ca4-027f794d6ece" />

Проверка

<img width="782" height="664" alt="image" src="https://github.com/user-attachments/assets/08393979-9ade-4ba6-8bb7-07d5bac43dab" />

11. Метод DELETE /tasks/{id}`. 

В файле handlers.go был добавлена функция DELETETask

<img width="507" height="385" alt="image" src="https://github.com/user-attachments/assets/ae152a73-fe97-459b-a658-8f8322dec745" />

Проверка

<img width="792" height="676" alt="image" src="https://github.com/user-attachments/assets/96e23ff4-d056-4241-b0cb-f348c5c6dd11" />

Список после удаления

<img width="792" height="759" alt="image" src="https://github.com/user-attachments/assets/abd5affc-7498-402a-ae2d-2f7e6db59799" />

12. Graceful shutdown через http.Server и контекст.

Для реализации плавного завершения сервера был обновлен код в файле main.go

<img width="603" height="514" alt="image" src="https://github.com/user-attachments/assets/9dc0c55d-0319-42ea-895b-9cad175abaf1" />

13. Юнит-тесты обработчиков с httptest.

Был написан отдельный файл test.go с тестами

<img width="540" height="742" alt="image" src="https://github.com/user-attachments/assets/154a2a80-f029-4e27-833e-d16c8bc610a1" />

Проверка

<img width="526" height="313" alt="image" src="https://github.com/user-attachments/assets/b265c6d8-a783-4384-b69c-4b5b6c54142e" />

