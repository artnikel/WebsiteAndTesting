# WebsiteAndTesting
To run this project you must:
• install Open Server. I use Open Server Panel (for Windows)
• run your Open Server and open phpMyAdmin
• create database (name my database in project is "golang")
• in main.go edit this string:"root:root@tcp(127.0.0.1:3306)/golang" on the "<login in phpMA>:<password in phpMA>@tcp(17.0.0.1:3306)/<name your database in 
   phpMA>
• in your database create Table "users"
• in this table create 3 field:
  1) ID (int, UNSIGNED, PRIMARY, AUTO_INCREMENT) //this field is dispensable in this project
  2)Mail (varchar(25))
  3)Password (varchar(8))
• run main.go
AND ALL IS DONE! (if you want check for correct operation your database, run first "go test")
