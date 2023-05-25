# Passbook

Passbook представляет собой клиент-серверную систему, позволяющую пользователю надёжно и безопасно хранить логины, пароли, бинарные данные и прочую приватную информацию..

## Features

- регистрация, аутентификация и авторизация пользователей
- хранение приватных данных
- передача приватных данных владельцу по запросу

Типы хранимой информации:

- пары логин/пароль
- извольные текстовые данные
- произвольные бинарные данные
- данные банковских карт


Регистрация, авторизация происходят через Redis (запуск redis: redis-server)
остальной функционал через Postgres.

Запуск сервера с postgres из папки server: go run ./cmd/passbook/main.go
Запуск клиента из папки server: go run ./cmd/passbook/main.go

Примеры комманд клиента(клиентский main.go находиться в папке cli):

   - go run main.go register --login leo --password qwerty
   - go run main.go login --login leo --password qwerty
   - go run main.go logout
   - go run main.go me
   - go run main.go account --title vk.ru --login leo --password qwerty
   - go run main.go text --title mybook --data "lorem iposum dolor"
   - go run main.go bin --title mybin --path hello.txt
   - go run main.go card --title sber --name "leo di catrio" --card 24234436456457 --date 21/23 --cvc sd
   - go run main.go get_title --title "mybook"
   - go run main.go full

## License



**Free Software**
