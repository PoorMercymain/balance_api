# balance_api
<p>Сначала используем файл docker-compose.yml для запуска PostgreSQL в Docker (т.е. просто прописываем в терминале docker-compose up, находясь в директории с данным файлом)
После этого открываем проект в GoLand, и в конфигурации запуска указываем следующие переменные окружения (в случае чего, данные значения можно посмотреть в файле .env.example):
BALANCE_API_HOSTNAME=localhost;BALANCE_API_PORT=3001;BALANCE_API_USERNAME=balanceapi;BALANCE_API_PASSWORD=balanceapi;BALANCE_API_DATABASE_NAME=balanceapi
Далее, используя файл scheme.sql, создаем таблицы в БД (можно это сделать прямо из GoLand, открыв данный файлик, настроив data source и нажав на соответствующую кнопку в интерфейсе IDE)
После этого можно запустить файл main.go, сервис будет работать на localhost:8000 (конкретные роуты указаны в мэйне (как это выглядит на данный момент - см. на рисунке ниже), но я в основном тестил базовый функционал, так что что-то не из списка минимальных требований может не работать/работать некорректно (но скорее всего подавляющее большинство работает как надо))</p>
<p align=center><img src="https://user-images.githubusercontent.com/67076111/198801546-0a04cbd2-7b5b-4c6b-84c2-76646f5d3654.png"></p>
<p align=center>Роуты в мэйне</p>
<p></p>
Ниже приведены примеры из Postman, показывающие, как можно пользоваться данным сервисом.
<p align=center><img src="https://user-images.githubusercontent.com/67076111/198809445-acff72ba-4b26-47b1-855e-dcef8f43ecb5.png"></p>
<p align=center>Пример добавления пользователя</p>
<p align=center><img src="https://user-images.githubusercontent.com/67076111/198808135-de46f6e6-9e35-4c1d-8c7d-8195fce22919.png"></p>
<p align=center>Пример создания услуги</p>
<p align=center><img src="https://user-images.githubusercontent.com/67076111/198810982-57fd4682-ca5a-4645-a9a7-f949957f5129.png"></p>
<p align=center>Пример создания заказа</p>
<p>Тут уже стоит учитывать, что id пользователя должно присутствовать в таблице user (добавить его туда можно так, как показано на первом примере из постмана). Также может показаться, что у заказа не хватает указания сервисов, которые к нему относятся, но для этого есть отдельная таблица, отображающая связь многие-ко-многим (order_service)</p>
<p align=center><img src="https://user-images.githubusercontent.com/67076111/198812892-20ac6c22-a7d3-4d2a-ac1e-5f0cec085fcd.png"></p>
<p align=center>Пример добавления услуги к заказу</p>
<p align=center><img src="https://user-images.githubusercontent.com/67076111/198813072-380441f0-576c-4bf1-8192-60d806f58965.png"></p>
<p align=center>Пример запроса для резервирования денег за услугу на отдельном счете (можно проверить, что работает корректно, с помощью запроса баланса пользователя до и после запроса резервирования)</p>
<p align=center><img src="https://user-images.githubusercontent.com/67076111/198813214-5266aaa7-f464-48d8-aa4e-5d507ac8d99d.png"></p>
<p align=center>Пример запроса для признания выручки (отправляет данные в отдельную таблицу (accounting_report) и удаляет услугу из резерва)</p>
<p align=center><img src="https://user-images.githubusercontent.com/67076111/198813405-885b5dbe-e10c-4bd0-ae9e-e056135d80d3.png"></p>
<p align=center>Пример запроса для отмены резерва на случай, когда услугу оказать может не получиться (возвращает деньги на счет пользователя и убирает услугу из резерва)</p>
<p align=center><img src="https://user-images.githubusercontent.com/67076111/198813557-40e3826f-7b55-4ab0-9c07-9e92c6100be7.png"></p>
<p align=center>Пример запроса для добавления средств на счет пользователя (пришлось добавить два поля для выполнения доп задания)</p>
<p align=center><img src="https://user-images.githubusercontent.com/67076111/198813628-0febc369-45ec-4fc6-86ac-33939996e03e.png"></p>
<p align=center>Пример запроса баланса пользователя (по его id)</p>
<p align=center><img src="https://user-images.githubusercontent.com/67076111/198814637-844b0dc7-444c-4569-a73d-339f0bf6f536.png"></p>
<p align=center>Пример запроса для постраничного вывода списка транзакций (limit - максимальное число транзакций в выводе, sort_dir - если прописать desc, то по убыванию, если ничего не прописать - по возрастанию, в sort_by указывается название колонки таблицы user_report для сортировки по этой колонке)</p>
<p align=center><img src="https://user-images.githubusercontent.com/67076111/198814021-fd9e6f68-06da-4076-9347-d9510d84b87b.png"></p>
<p align=center>Пример запроса для создания отчета в формате csv. Это я еще попробую доделать...</p>
<p>Ну, вроде основное показал, но, вообще, можно поэксперементировать и с другими роутами, и, как я уже говорил выше, они даже скорее всего сработают (но в некоторых местах может вылезти ошибка, говорящая о том, что строки не найдены, но на самом деле не всегда действительно ошибка)</p>
