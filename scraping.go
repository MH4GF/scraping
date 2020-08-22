package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/sclevine/agouti"
	"log"
	"os"
	"strings"
	"time"
)

var ua = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.116 Safari/537.36"
var lpPath = "https://moneyforward.com/"

func scparing() *goquery.Selection {
	page := signInPage()

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

	return result
}

func signInPage() *agouti.Page {
	driver := newDriver()

	page, err := driver.NewPage()
	if err != nil {
		log.Fatal(err)
	}

	if err := page.Navigate(lpPath); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
	if err := page.FindByLink("ログイン").Click(); err != nil {
		log.Fatal(err)
	}
	time.Sleep(1 * time.Second)
	if err := page.FindByLink("メールアドレスでログイン").Click(); err != nil {
		log.Fatal(err)
	}
	time.Sleep(1 * time.Second)

	emailInput := page.FindByName("mfid_user[email]")
	emailSubmit := page.FindByClass("submitBtn")
	if err := emailInput.Fill(os.Getenv("SIGN_IN_EMAIL")); err != nil {
		log.Fatal(err)
	}
	if err := emailSubmit.Submit(); err != nil {
		log.Fatalf("Failed to login:%v", err)
	}
	time.Sleep(1 * time.Second)

	passwordInput := page.FindByName("mfid_user[password]")
	passwordSubmit := page.FindByClass("submitBtn")
	if err := passwordInput.Fill(os.Getenv("SIGN_IN_PASSWORD")); err != nil {
		log.Fatal(err)
	}
	if err := passwordSubmit.Submit(); err != nil {
		log.Fatalf("Failed to login:%v", err)
	}
	time.Sleep(1 * time.Second)

	return page
}

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
