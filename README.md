# Golang HttpRouter
`HttpRouter` merupakan salah satu Open Source library yang populer di golang untuk Http Handler<br>
HttpRouter terkenal dengan kecepatannya dan juga sangat minimalis <br>
Hal ini karena HttpRouter Hanya memiliki fitur untuk routing saja<br>
https://github.com/julienschmidt/httprouter

## Menambahkan HttpRouter Ke Project 
```bash
# Install Library HttpRouter
go get github.com/julienschmidt/httprouter

# Install Library untuk unit test
go get github.com/stretchr/testify
```

## Router
Inti dari `HttpRouter` adalah `struct Router`<br>
Router merupakan implementasi dari http.Handler, sehingga kita bisa dengan mudah menambahkan kedalam `http.Server`<br>
Untuk membuat router, kita bisa menggunakan function `httprouter.New()`, yang akan mengembalikan `Router pointer`