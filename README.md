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

## Params
`httprouter.Handle` memiliki parameter ketiga yaitu Params<br>
`Params` merupakan tempat untuk menyimpan parameter yang dikirim dari client<br>
Namun `Params` ini bukan `Query Parameter`, melainkan parameter di `URL` <br>
Kadang kita butuh membuat `URL` yang tidak fix, alias bisa berubah-ubah, misalnya `/produk/1`, `/produk/2`, dan seterusnya<br>
ServeMux tidak mendukung hal ini, Namun `Router` mendukung hal tersebut<br>
Parameter yang dinamis  yang terdapat di URL, secara otomatis dikumpulkan dalam `Params` <br>
Namun kita harus memberi tahu ketika menambahkan `Route`, dibagian mana kita akan membuat `URL Pathn`nya menjadi dinamis ``/produk/:name``

## Router Pattern
### Named Parameter
`Named Parameter` adalah pola pembuatkan parameter dengan menggunakan nama<br>
Setiap nama parameter harus diawali dengan `:`(titik dua), lalu diikuti dengan nama parameter<br>
contoh: <br>
`Pattern` 				`/user/:user`
`/user/luna`		 =>	`match`
`/user/you`	 		 => `match`
`/user/luna/profile` => `not match`
`/user/`			 => `not match`
### Catch All Parameter
Selain `Named Parameter`, ada juga `catch all parameter`, yaitu menangkap semua parameter <br>
`Catch All Parameter` harus diawali  dengan `*`(Bintang), lalu diikuti dengan nama parameter<br>
`Catch All Parameter` harus berada diposisi akhir `URL`
`Pattern` 				`/src/*filepath`
`/src/`				 	=> `not match`
`/src/somefile`		 	=> `match`
`/src/subdir/somefile`	=> `match`
