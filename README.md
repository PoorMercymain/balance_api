# balance_api
Сначала используем файл docker-compose.yml для запуска PostgreSQL (т.е. прописываем в терминале docker-compose up)
После этого открываем проект в GoLand, и в конфигурации запуска указываем следующие переменные окружения (значения, в случае чего, можно посмотреть в файле .env.example):
BALANCE_API_HOSTNAME=localhost;BALANCE_API_PORT=3001;BALANCE_API_USERNAME=balanceapi;BALANCE_API_PASSWORD=balanceapi;BALANCE_API_DATABASE_NAME=balanceapi
![изображение](https://user-images.githubusercontent.com/67076111/198798000-c7ce86b0-2e56-49ce-97b6-624da78319af.png)
