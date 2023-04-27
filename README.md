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
### HTTP Method
`Router` mirip dengan `ServeMux`, dimana kita bisa menambahkan `route` kedalam `Router`<br>
Kelebihkan dibandingkan dengan ServeMux adalah pada router, kita bisa menentukan `HTTP Method` yang ingin kita gunakan, misalnya `GET, POST, PUT` dan lain-lain<br>
Cara menambahkan route kedalam Router adalah dengan menggunakan function yang sama dengan `HTTP Method`nya, misalnya `router.GET()`, `router.POST()` dan lain-lain
#### httprouter.Handle
Saat kita menggunakan `ServeMux`, ketika menambahkan route, kita menambahakan `http.Handler` <br>
Berbeda dengan Router, pada router kita tidak bisa menggunakan `http.Handler` lagi, melainkan menggunakan  type `httprouter.Handle` <br>
Perbedaan dengan `http.Handler` adalah, pada `httprouter.Handle`, terdapat parameter ketiga