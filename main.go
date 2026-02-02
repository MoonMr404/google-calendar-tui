package main

import "github.com/rivo/tview"

func main() {
	app := tview.NewApplication()

	taskForm := tview.NewInputField().SetLabel("Task Title")
	timeStampForm := tview.NewInputField().SetLabel("Timestamp")
	dateForm := tview.NewInputField().SetLabel("Date")

	form := tview.NewForm().
		AddFormItem(taskForm).
		AddFormItem(timeStampForm).
		AddFormItem(dateForm).
		AddButton("Exit", func() {
			app.Stop()
		})

	form.SetBorder(true).SetTitle("Google Calendar TUI").SetTitleAlign(tview.AlignLeft)
	flex := tview.NewFlex().AddItem(form, 0, 1, true)
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}

}
