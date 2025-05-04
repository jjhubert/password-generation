package main

import (
	"crypto/rand"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"math/big"
	"strconv"
)

var (
	lowerChars  = "abcdefghijklmnopqrstuvwxyz"
	upperChars  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars = "0123456789"
	symbolChars = "!@#$%^&*()-_=+[]{}|;:,.<>/?"
)

func generatePassword(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		b[i] = charset[num.Int64()]
	}
	return string(b)
}

func main() {
	a := app.New()
	w := a.NewWindow("Password Generator")
	w.Resize(fyne.NewSize(500, 400))
	
	// Options
	lowerCheck := widget.NewCheck("a-z", nil)
	lowerCheck.SetChecked(true)
	upperCheck := widget.NewCheck("A-Z", nil)
	upperCheck.SetChecked(true)
	numberCheck := widget.NewCheck("0-9", nil)
	numberCheck.SetChecked(true)
	symbolCheck := widget.NewCheck("!@#$%", nil)
	symbolCheck.SetChecked(true)

	lengthEntry := widget.NewEntry()
	lengthEntry.SetText("16")
	countEntry := widget.NewEntry()
	countEntry.SetText("1")

	resultBox := container.NewVBox()

	genBtn := widget.NewButton("生成密码", func() {
		resultBox.Objects = nil
		// Assemble character set
		charset := ""
		if lowerCheck.Checked {
			charset += lowerChars
		}
		if upperCheck.Checked {
			charset += upperChars
		}
		if numberCheck.Checked {
			charset += numberChars
		}
		if symbolCheck.Checked {
			charset += symbolChars
		}
		if charset == "" {
			resultBox.Add(widget.NewLabel("请至少选择一种字符集"))
			w.Content().Refresh()
			return
		}
		// 长度和数量
		length, err1 := strconv.Atoi(lengthEntry.Text)
		count, err2 := strconv.Atoi(countEntry.Text)
		if err1 != nil || length < 6 {
			resultBox.Add(widget.NewLabel("密码长度最小为6"))
			w.Content().Refresh()
			return
		}
		if err2 != nil || count < 1 || count > 10 {
			resultBox.Add(widget.NewLabel("生成数量范围1-10"))
			w.Content().Refresh()
			return
		}
		// Generate passwords
		for i := 0; i < count; i++ {
			pwd := generatePassword(length, charset)
			pwdEntry := widget.NewEntry()
			pwdEntry.SetText(pwd)
			pwdEntry.Disable()
			copyBtn := widget.NewButton("复制", func(p string) func() {
				return func() {
					w.Clipboard().SetContent(p)
				}
			}(pwd))
			// row := container.NewHBox(nil,nil,nil,pwdEntry, copyBtn)
			row := container.NewGridWithColumns(2,pwdEntry, copyBtn)
			resultBox.Add(row)
		}
		w.Content().Refresh()
	})

	// Layout
	options := container.NewVBox(
		container.NewHBox(lowerCheck, upperCheck, numberCheck, symbolCheck),
		container.NewHBox(widget.NewLabel("密码长度："), lengthEntry, widget.NewLabel("生成数量："), countEntry),
		genBtn,
	)

	mainBox := container.NewVBox(options, resultBox)
	w.SetContent(mainBox)
	w.ShowAndRun()
}
