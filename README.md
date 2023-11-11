Instruction:
* Chạy server MySql:
  * name: root
  * password: pwd (nếu khác pwd thì đổi pwd thành password của mình trong file .env)\
  * port: 3306
* Chạy server:
  * Chạy lệnh "go run" với file main.go
* Một số api hiện tại:
  * POST: http://localhost:8080/api/auth/register
    * req
      ```
      {
            Username: string
            Password: string
            Role: uint32
            Telephone: string
      }
      ```
  * POST: http://localhost:8080/api/auth/login -> nếu thành công trả về id và role
    * req
      ```
      {
             Username: string
             Password: string
      }
      ```
  * POST: http://localhost:8080/api/auth/logout
    * req
      ```
      {
              Id: uint32
      }
      ```