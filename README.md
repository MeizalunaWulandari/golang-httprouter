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

## Serve File
`Router` juga mendukung static file menggunakan function function `ServeFiles(Path, FileSystem)`<br>
Dimana pada `Path`, kita harus menggunakan Catch All Parameter<br>
Sedangkan pada `FileSystem` kita harus melakukan manaual load dari folder atau menggunakan `golang embed` 

## Panic Handler
Jika terjadi `Panic` secara otomatis akan error dan web akan berhenti mengembalikan response<br>
Kadang saat terjadi `panic`, kita ingin melakukan sesuatu, misalnya memberitahu terjadi kesalahan di web, atau bahkan mengirim informasi log kesalahan yang terjadi<br>
Jika kita ingin menangani panic, kita harus membuat `Middleware` khusus secara manual, Namun di `Router` sudah disediakan untuk menangani panic, caranya dengan menggunakan atribute `PanicHandler:func(http.ResponseWriter, *http.Request, interface{})`

## Not Found Handler
Selain `panic handler`, Router juga memiliki `not found handler`<br>
`Not Found Handler` adalah handler yang dieksekusi ketika client mencoba melakukan Request URL yang yang memang tidak terdapat di Router<br>
Secara Default, jika ada route tidak ditemukan, Router akan melanjutkan request ke `http.NotFound`, namun kita bisa mengubahnya<br>
Caranya dengan mengubah `router.NotFound = http.Handler` 

## Method Not Allowed Handler
Saat menggunakan `ServeMux`, kita tidak bisa menentukan `HTTP Method` apa yang digunakan untuk Handler<br>
Namun pada `Router`, kita bisa menentukan `HTTP Method` yang ingin kita gunakan<br>
Jika client tidak mengirim sesuai HTTP Method yang kita tentukan akan akan terjadi error `Method Not Allowed`<br>
Secara default, jika terjadi error seperti itu, maka Router akan memanggil function `http.Error`<br>
Jika ingin mengubahnya kita bisa gunakan `router.MethodNotAllowed = http.Handler`