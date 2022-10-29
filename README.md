# balance_api
<h3>Как запустить?</h3>
<p>Сначала используем файл docker-compose.yml для запуска PostgreSQL в Docker (т.е. просто <b>прописываем в терминале docker-compose up, находясь в директории с данным файлом</b>)
После этого <b>открываем проект в GoLand</b>, и <b>в конфигурации запуска указываем следующие переменные окружения</b> (в случае чего, данные значения можно посмотреть в файле .env.example):
<br><br><b>BALANCE_API_HOSTNAME=localhost;BALANCE_API_PORT=3001;BALANCE_API_USERNAME=balanceapi;BALANCE_API_PASSWORD=balanceapi;BALANCE_API_DATABASE_NAME=balanceapi</b><br><br>
Далее, <b>используя файл scheme.sql, создаем таблицы в БД</b> (можно это сделать прямо из GoLand, открыв данный файлик, настроив data source и нажав на соответствующую кнопку в интерфейсе IDE)
После этого <b>можно запустить файл main.go</b>, сервис будет работать на localhost:8000 (конкретные роуты указаны в мэйне (как это выглядит на данный момент - см. на рисунке ниже), но я в основном тестил базовый функционал, так что что-то не из списка минимальных требований может не работать/работать некорректно (но скорее всего подавляющее большинство работает как надо))</p>
<p align=center><img src="https://user-images.githubusercontent.com/67076111/198839929-a61931fb-9551-40a0-8203-3bf812611a79.png"></p>
<p align=center>Роуты в мэйне</p>
<p></p>
<h3>Примеры запросов и комментарии к ним</h3>
<b align=center><p>Ниже приведены примеры из Postman, показывающие, как можно пользоваться разработанным сервисом.</p></b>
<p align=center><img src="https://user-images.githubusercontent.com/67076111/198809445-acff72ba-4b26-47b1-855e-dcef8f43ecb5.png"></p>
<p align=center>Пример добавления пользователя</p>
<p align=center><img src="https://user-images.githubusercontent.com/67076111/198808135-de46f6e6-9e35-4c1d-8c7d-8195fce22919.png"></p>
<p align=center>Пример создания услуги</p>
<p align=center><img src="https://user-images.githubusercontent.com/67076111/198810982-57fd4682-ca5a-4645-a9a7-f949957f5129.png"></p>
<p align=center>Пример создания заказа</p>
<p>Тут уже стоит учитывать, что id пользователя должно присутствовать в таблице user (добавить его туда можно так, как показано на первом примере из постмана). Также может показаться, что у заказа не хватает указания сервисов, которые к нему относятся (т.к. заказ, как я это понял, может включать в себя несколько услуг), но для этого есть отдельная таблица, отображающая связь многие-ко-многим (order_service, структура каждой таблицы указана в файле scheme.sql)</p>
<p align=center><img src="https://user-images.githubusercontent.com/67076111/198812892-20ac6c22-a7d3-4d2a-ac1e-5f0cec085fcd.png"></p>
<p align=center>Пример добавления услуги к заказу, тут нужно учесть, что должна существовать как соответствующая услуга, так и заказ</p>
<p align=center><img src="https://user-images.githubusercontent.com/67076111/198813072-380441f0-576c-4bf1-8192-60d806f58965.png"></p>
<p align=center>Пример запроса для резервирования денег за услугу на отдельном счете, пользователь, услуга и заказ, соответственно, к этому моменту уже должны существовать (проверить корректность работы можно с помощью запроса баланса пользователя до и после запроса резервирования). <b>Является частью основного задания</b></p>
<p align=center><img src="https://user-images.githubusercontent.com/67076111/198813214-5266aaa7-f464-48d8-aa4e-5d507ac8d99d.png"></p>
<p align=center>Пример запроса для признания выручки (отправляет данные в отдельную таблицу (accounting_report) и удаляет услугу из резерва). Тут нужно учитывать, что если в один и тот же заказ входит несколько одинаковых услуг, признание выручки произойдет по им всем (т.е. если в заказ входит услуга1, стоящая 500, два раза, при признании выручки лучше указать 1000) <b>Является частью основного задания</b></p>
<p align=center><img src="https://user-images.githubusercontent.com/67076111/198813405-885b5dbe-e10c-4bd0-ae9e-e056135d80d3.png"></p>
<p align=center>Пример запроса для отмены резерва на случай, когда услугу оказать не получится (возвращает деньги на счет пользователя и убирает услугу из резерва). <i>Вроде это нужно было в пункте "Будет плюсом"</i></p>
<p align=center><img src="https://user-images.githubusercontent.com/67076111/198813557-40e3826f-7b55-4ab0-9c07-9e92c6100be7.png"></p>
<p align=center>Пример запроса для добавления средств на счет пользователя (пришлось добавить два поля для выполнения доп задания). <b>Является частью основного задания</b></p>
<p align=center><img src="https://user-images.githubusercontent.com/67076111/198813628-0febc369-45ec-4fc6-86ac-33939996e03e.png"></p>
<p align=center>Пример запроса баланса пользователя (по его id, в примере id=2). <b>Является частью основного задания</b></p>
<p align=center><img src="https://user-images.githubusercontent.com/67076111/198814637-844b0dc7-444c-4569-a73d-339f0bf6f536.png"></p>
<p align=center>Пример запроса для постраничного вывода списка транзакций (limit - максимальное число транзакций в выводе, sort_dir - если прописать desc, то по убыванию, если ничего не прописать - по возрастанию, в sort_by указывается название колонки таблицы user_report (варианты: либо money, либо transaction_date, либо оба вместе - money, transaction_date) для сортировки по этой колонке). <b><i>Является доп. заданием</i></b></p>
<p align=center><img src="https://user-images.githubusercontent.com/67076111/198838458-61ceb420-136e-429d-a680-43ef019d0999.png"></p>
<p align=center>Пример запроса для создания отчета в формате csv. На выход дается ссылка, по которой можно сделать запрос содержимого файла csv. <b><i>Это доп. задание</i></b></p>
<p>Ну, вроде основное показал, но, вообще, можно поэксперементировать и с другими роутами, и, как я уже говорил выше, они даже скорее всего сработают (но в некоторых местах может вылезти ошибка, говорящая о том, что строки не найдены, но на самом деле не всегда действительно ошибка)</p>
