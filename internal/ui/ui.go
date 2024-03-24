package ui

import (
	"Rail-Ticket-Notifier/cmd/handlers"
	"Rail-Ticket-Notifier/internal/arguments"
	"Rail-Ticket-Notifier/internal/models"
	"Rail-Ticket-Notifier/utils/constants"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	xwidget "fyne.io/x/fyne/widget"
	"strconv"
	"strings"
	"time"
)

func InitializeUIAndForm() models.ElementsOfUI {
	a := app.New()

	window := a.NewWindow("Automated Rail Ticket Booker & Notifier")
	window.Resize(fyne.NewSize(800, 600))
	introLabel := widget.NewLabel(constants.INTRO_MSG)
	// Create form fields with default values
	fromEntry := widget.NewEntry()
	fromEntry.SetText(arguments.FROM) // Default value from arguments package
	toEntry := widget.NewEntry()
	toEntry.SetText(arguments.TO) // Default value from arguments package
	dateEntry := widget.NewEntry()
	dateEntry.SetText(arguments.DATE) // Default value from arguments package
	seatCountEntry := widget.NewEntry()
	seatCountEntry.SetText(strconv.Itoa(int(arguments.SEAT_COUNT))) // Convert uint to string
	seatTypesEntry := widget.NewEntry()
	seatTypesEntry.SetText(strings.Join(arguments.SEAT_TYPE_ARRAY, ",")) // Default value from arguments package
	trainsEntry := widget.NewEntry()
	trainsEntry.SetText(strings.Join(arguments.SPECIFIC_TRAIN_ARRAY, ","))
	emailEntry := widget.NewEntry()
	emailEntry.SetText(arguments.RECEIVER_EMAIL_ADDRESS)
	phoneEntry := widget.NewEntry()
	phoneEntry.SetText(arguments.PHONE_NUMBER)
	phoneEntry.Disable()

	content := container.NewVBox(introLabel, fromEntry, toEntry, dateEntry, seatCountEntry, seatTypesEntry, trainsEntry)

	window.SetContent(content)

	uiElements := models.ElementsOfUI{
		Window:         window,
		IntroLabel:     introLabel,
		FromEntry:      fromEntry,
		ToEntry:        toEntry,
		DateEntry:      dateEntry,
		SeatCountEntry: seatCountEntry,
		SeatTypesEntry: seatTypesEntry,
		TrainsEntry:    trainsEntry,
		EmailEntry:     emailEntry,
		PhoneEntry:     phoneEntry,
	}

	return uiElements
}

func CreateForm(uiElements models.ElementsOfUI) *fyne.Container {

	calendar := GetCalendar(func(t time.Time) {
		uiElements.DateEntry.SetText(t.Format("02-Jan-2006"))
	})

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "From (Capital Case)", Widget: uiElements.FromEntry},
			{Text: "To (Capital Case)", Widget: uiElements.ToEntry},
			{Text: "Date Of Journey (Choose From Calender)", Widget: uiElements.DateEntry},
			{Text: "(Only from current date to next 10 days)", Widget: calendar},
			{Text: "Seat Count (1 to Max 4)", Widget: uiElements.SeatCountEntry},
			{Text: "Seat Types (Will Prioritize Serial Wise)", Widget: uiElements.SeatTypesEntry},
			{Text: "Trains (Choose only One.)", Widget: uiElements.TrainsEntry},
			{Text: "Email address (To receive mail after done)", Widget: uiElements.EmailEntry},
			{Text: "Phone Number (Currently unavailable)", Widget: uiElements.PhoneEntry},
		},
	}

	submitButton := getSubmitButton()

	submitButton.OnTapped = func() {
		handlers.HandleFormSubmission(uiElements, submitButton)
	}

	return container.NewVBox(
		form,
		submitButton,
	)
}

func getSubmitButton() *widget.Button {
	submitButton := widget.NewButton("Start Search", func() {})
	return submitButton
}

func GetCalendar(onSelected func(time.Time)) *xwidget.Calendar {
	startingDate := time.Now()
	return xwidget.NewCalendar(startingDate, onSelected)
}
