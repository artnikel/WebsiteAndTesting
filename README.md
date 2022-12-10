# WebsiteAndTesting
To run this project you must:
* install Open Server. I use Open Server Panel (for Windows)<br>
* run your Open Server and open phpMyAdmin<br>
* create database (name my database in project is "golang")<br>
* in main.go edit this string:"root:root@tcp(127.0.0.1:3306)/golang" <br>on the "(login in phpMA):(password in phpMA)@tcp(17.0.0.1:3306)/(name your database in 
   phpMA)<br>
* in your database create Table "users"<br>
* in this table create 3 field:<br>
1. ID (int, UNSIGNED, PRIMARY, AUTO_INCREMENT) //this field is dispensable in this project<br>
2. Mail (varchar(25))<br>
3. Password (varchar(8))<br>
* run main.go<br><br>
AND ALL IS DONE! (if you want check for correct operation your database, run first "go test")
