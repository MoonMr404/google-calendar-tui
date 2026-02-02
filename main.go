package main

import (
	"fmt"
	"github.com/rivo/tview"
)

type Task struct {
	Title string `json:"title"`
	Time  string `json:"time"`
	Date  string `json:"date"`
}

//add limit
//func getTaskList() Task[] {
//
//	return
//}

func main() {
	app := tview.NewApplication()

	taskForm := tview.NewInputField().SetLabel("Task Title")
	timeStampForm := tview.NewInputField().SetLabel("Timestamp")
	dateForm := tview.NewInputField().SetLabel("Date")
	textResultView := tview.NewTextView().SetDynamicColors(true). // Enable dynamic coloring of text
		SetRegions(true). // Allows regions for interaction (not used here)
		SetWordWrap(true) // Enables word wrapping to fit the TextView size

	textResultView.SetBorder(true).SetTitle("Results")

	form := tview.NewForm().
		AddFormItem(taskForm).
		AddFormItem(timeStampForm).
		AddFormItem(dateForm).
		AddButton("Get TaskList", func() {
			fmt.Fprintln(textResultView, "Lista Task")
		}).AddButton("Exit", func() {
		app.Stop()
	})

	form.SetBorder(true).SetTitle("Google Calendar TUI").SetTitleAlign(tview.AlignLeft)
	flex := tview.NewFlex().AddItem(form, 0, 1, true).AddItem(textResultView, 0, 1, false)

	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}

}
