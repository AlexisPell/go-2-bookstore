cmd -> main.go

pkg ->

- config -> app.go
- controllers -> book-controller
- models -> book.go
- routes -> bookstore-routes
- utils -> utils.go

routes ->

- GET /book - get books
- GET /book/{bookId} - get book
- POST /book - create book
- PUT /book/{bookId} - update book
- DELETE /book/{bookId} - delete book

go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/mysql
go get github.com/gorilla/mux
