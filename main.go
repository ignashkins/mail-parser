package main

import (
	"fmt"
	"io"
	"log"
	"net/mail"
	"strings"
)

func main() {

	msg := `2021-10-13 13:06:24
To: admin@local.ru
Subject: =?UTF-8?B?Qml0cml4IERvY2tlcjog0KDQtdCz0LjRgdGC0YDQsNGG0LjQvtC90L3QsNGPINC40L3RhNC+0YDQvNCw0YbQuNGP?=
From: bitrix.docker@gmail.com
Reply-To: bitrix.docker@gmail.com
X-EVENT_NAME: USER_INFO
X-Priority: 3 (Normal)
Date: Wed, 13 Oct 2021 13:06:24 +0300
MIME-Version: 1.0
X-MID: 0.2 (13.10.2021 13:06:24)
Content-Type: text/plain; charset=UTF-8
Content-Transfer-Encoding: 8bit

Информационное сообщение сайта Bitrix Docker
------------------------------------------
 ,

Администратор сайта изменил ваши регистрационные данные.

Ваша регистрационная информация:

ID пользователя: 1
Статус профиля: активен
Login: admin

Вы можете изменить пароль, перейдя по следующей ссылке:
http://bitrix.docker/auth/index.php?change_password=yes&lang=ru&USER_CHECKWORD=b77bb426989a319fcdc19c5938f80140&USER_LOGIN=admin

Сообщение сгенерировано автоматически.
`

	r := strings.NewReader(msg)
	m, err := mail.ReadMessage(r)
	if err != nil {
		log.Fatal(err)
	}

	header := m.Header
	fmt.Println("Date:", header.Get("Date"))
	fmt.Println("From:", header.Get("From"))
	fmt.Println("To:", header.Get("To"))
	fmt.Println("Subject:", header.Get("Subject"))

	body, err := io.ReadAll(m.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)

}