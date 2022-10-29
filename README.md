# balance_api
Сначала используем файл docker-compose.yml для запуска PostgreSQL (т.е. прописываем в терминале docker-compose up)
После этого открываем проект в GoLand, и в конфигурации запуска указываем следующие переменные окружения (значения, в случае чего, можно посмотреть в файле .env.example):
BALANCE_API_HOSTNAME=localhost;BALANCE_API_PORT=3001;BALANCE_API_USERNAME=balanceapi;BALANCE_API_PASSWORD=balanceapi;BALANCE_API_DATABASE_NAME=balanceapi
После этого можно запустить приложение, оно будет работать на localhost:8000 (конкретные роуты указаны в мэйне (как это выглядит на данный момент - см. на рисунке ниже), но я в основном проверял базовый функционал, так что что-то не из списка минимальных требований может не работать/работать некорректно (но скорее всего подавляющее большинство работает как надо))
![изображение](https://user-images.githubusercontent.com/67076111/198801546-0a04cbd2-7b5b-4c6b-84c2-76646f5d3654.png)

![изображение](https://user-images.githubusercontent.com/67076111/198798000-c7ce86b0-2e56-49ce-97b6-624da78319af.png)
