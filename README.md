## TODO

---

1. Ubah yang belum di service handler ✅
2. Do all function in handler (Create ✅, Reads ✅, Update ✅, Delete ✅, ReadByID ✅)
3. Sesuain response dengan schema ✅
4. Error handler di Findbyid, delete, update pada resultnya ✅
5. Use environment ✅

---

- Table user

---

1. Create table ✅
2. Do all function in handler (Create ✅, Reads ✅, Update ✅, Delete ✅, ReadByID ✅)
3. Make email unique ✅
4. Dont showing password in APIResponse ✅
5. Password hashing to store in DB ✅
6. Authentication (Bearer) (doing)
   - Login ✅
   - GenToken ✅
   - DecodeToken ✅
   - restrictedRoute ✅
7. Relation each other (one-to-many) -> user has many books ✅
8. Make Code better (dry, etc)
9. Update book with image store in AWS S3

## Flow App

---

    server -> repository -> service -> handler

## How to run

---

- `swag init -g server.go`
- `go run server.go`

or

- `swag init -g server.go`
- `go build server.go`
- `./server`

visit

- `http://127.0.0.1:5000/swagger/index.html` to see Doc API

### Tutorial from

---

[YT Agung Setiawan - Tutorial Golang Web API Bahasa Indonesia - Full Course](https://www.youtube.com/watch?v=GjI0GSvmcSU&t=290s)
