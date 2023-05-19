# passbook

Регистрация, авторизация происходят через Redis (запуск redis: redis-server)
остальной функционал через Postgres.

Запуск сервера с postgres: go run ./cmd/passbook/main.go

Примеры комманд клиента(клиентский main.go находиться в папке cli):

    go run main.go register --login leo --password qwerty
    go run main.go login --login leo --password qwerty
    go run main.go logout
    go run main.go account --title vk.ru --login leo --password qwerty
    go run main.go text --title mybook --data "lorem iposum dolor"
    go run main.go card --title sber --name "leo di catrio" --card 24234436456457 --date 21/23 --cvc sd
    go run main.go get_title --title "mybook"
    go run main.go full