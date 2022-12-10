package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestDatabase(t *testing.T) {
	testM, testP := "testMail", "123"

	database_insert(fmt.Sprintf("INSERT INTO `users`(`Mail`, `Password`) VALUES ('%s','%s')", testM, testP))

	database_select("SELECT `Mail`,`Password` FROM `users` ORDER BY `ID` DESC LIMIT 1")

	if testM != user.Mail {
		t.Errorf("Incorrect result. Except %s, got %s", testM, user.Mail)
	}

	if testP != user.Password {
		t.Errorf("Incorrect result. Except %s, got %s", testP, user.Password)
	}
}

func TestBorderLenPassword(t *testing.T){
	testM, testP := "EightLenPassword", "12345678"

	database_insert(fmt.Sprintf("INSERT INTO `users`(`Mail`, `Password`) VALUES ('%s','%s')", testM, testP))

	database_select("SELECT `Mail`,`Password` FROM `users` ORDER BY `ID` DESC LIMIT 1")

	if testM != user.Mail {
		t.Errorf("Incorrect result. Except %s, got %s", testM, user.Mail)
	}

	if testP != user.Password {
		t.Errorf("Incorrect result. Except %s, got %s", testP, user.Password)
	}

}

func TestBorderLenLogin(t *testing.T){
	testM, testP := "HelloHelloHelloHelloHello", "password"

	database_insert(fmt.Sprintf("INSERT INTO `users`(`Mail`, `Password`) VALUES ('%s','%s')", testM, testP))

	database_select("SELECT `Mail`,`Password` FROM `users` ORDER BY `ID` DESC LIMIT 1")

	if testM != user.Mail {
		t.Errorf("Incorrect result. Except %s, got %s", testM, user.Mail)
	}

	if testP != user.Password {
		t.Errorf("Incorrect result. Except %s, got %s", testP, user.Password)
	}

}

func TestNonePassword(t *testing.T){
	testM, testP := "testNonePassw", ""

	database_insert(fmt.Sprintf("INSERT INTO `users`(`Mail`, `Password`) VALUES ('%s','%s')", testM, testP))

	database_select("SELECT `Mail`,`Password` FROM `users` ORDER BY `ID` DESC LIMIT 1")

	if testP != user.Password {
		t.Errorf("Incorrect result. Except %s, got %s", testP, user.Password)
	}

}

func TestNoneLogin(t *testing.T){
	testM, testP := "", "789"

	database_insert(fmt.Sprintf("INSERT INTO `users`(`Mail`, `Password`) VALUES ('%s','%s')", testM, testP))

	database_select("SELECT `Mail`,`Password` FROM `users` ORDER BY `ID` DESC LIMIT 1")

	if testP != user.Password {
		t.Errorf("Incorrect result. Except %s, got %s", testP, user.Password)
	}

}

func TestSelectAllData(t *testing.T) {
	testM, testP := "testMailalldata", "456"

	database_insert(fmt.Sprintf("INSERT INTO `users`(`Mail`, `Password`) VALUES ('%s','%s')", testM, testP))

	database_alldata()

	if strings.Contains(allrootUsers, testM) != true {
		t.Errorf("Fail. %s don't contains %s", allrootUsers, testM)
	}
	if strings.Contains(allrootUsers, testP) != true {
		t.Errorf("Fail. %s don't contains %s", allrootUsers, testP)
	}
}

