package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
	"github.com/sclevine/agouti"
	"log"
	"os"
	"strings"
	"time"
)

var ua = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.116 Safari/537.36"
var loginPath = "https://moneyforward.com/users/sign_in"

func newDriver() *agouti.WebDriver {
	driver := agouti.ChromeDriver(agouti.ChromeOptions("args", []string{
		"--headless",
		"--user-agent=" + ua,
	}))

	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}

	return driver
}

func renderOutputJson() {
	driver := newDriver()

	page, err := driver.NewPage()
	if err != nil {
		log.Fatal(err)
	}

	if err := page.Navigate(loginPath); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}

	email := page.FindByID("sign_in_session_service_email")
	password := page.FindByID("sign_in_session_service_password")
	submit := page.FindByID("login-btn-sumit")

	email.Fill(os.Getenv("SIGN_IN_EMAIL"))
	password.Fill(os.Getenv("SIGN_IN_PASSWORD"))

	if err := submit.Submit(); err != nil {
		log.Fatalf("Failed to login:%v", err)
	}

	// 画面遷移のための時間を待つ
	time.Sleep(3 * time.Second)

	getSource, err := page.HTML()
	if err != nil {
		log.Fatalf("Failed to get HTML:%v", err)
	}

	readerCurContents := strings.NewReader(getSource)
	doc, err := goquery.NewDocumentFromReader(readerCurContents)
	if err != nil {
		log.Fatal(err)
	}

	result := doc.Find("#monthly_total_table_home tbody tr td")
	monthlyTotalTable := newTotalTable(result.Eq(0).Text(), result.Eq(1).Text(), result.Eq(2).Text())
	outputJson, err := json.Marshal(&monthlyTotalTable)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(outputJson))
}

func setupEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	setupEnv()
	renderOutputJson()
}
