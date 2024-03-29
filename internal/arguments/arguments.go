package arguments

import (
	"Rail-Ticket-Notifier/utils/constants"
	"flag"
	"strings"
	"time"
)

var (
	FROM                   string
	TO                     string
	DATE                   string
	SEAT_COUNT             uint
	PHONE_NUMBER           string
	SEAT_TYPE_ARRAY        = []string{"F_BERTH", "AC_B"} //, "S_CHAIR", "F_BERTH"}
	SPECIFIC_TRAIN_ARRAY   = []string{"TURNA"}           //{"SONAR", "TURNA", "SUBORNO"}
	RECEIVER_EMAIL_ADDRESS string
	GO_TO_BOOK_PAGE        uint
)

func init() {
	now := time.Now()
	twoDaysLater := now.AddDate(0, 0, 9)
	formattedDateAfterTwoDays := twoDaysLater.Format("02-Jan-2006")

	flag.StringVar(&FROM, "from", "Dhaka", "From city")
	flag.StringVar(&TO, "to", "Chattogram", "To city")
	flag.StringVar(&PHONE_NUMBER, "phone", "+8801555555555", "Phone")
	flag.StringVar(&DATE, "date", formattedDateAfterTwoDays, "Date of travel")
	flag.StringVar(&RECEIVER_EMAIL_ADDRESS, "email", "atiabintiaziz@gmail.com", "Email address")
	flag.UintVar(&SEAT_COUNT, "seatCount", 2, "Seat count")
	flag.UintVar(&GO_TO_BOOK_PAGE, "purchasePage", 1, "Go to purchase page")

	flag.Func("seatTypes", "Seat types", func(s string) error {
		SEAT_TYPE_ARRAY = strings.Split(s, ",")
		return nil
	})

	flag.Func("trains", "Specific trains", func(s string) error {
		SPECIFIC_TRAIN_ARRAY = strings.Split(s, ",")
		return nil
	})

	flag.Parse()
}

func UpdateArguments(from, to, date, email string, goToBookPage bool, seatCount uint, seatTypes, trains []string) {
	FROM = from
	TO = to
	DATE = date
	RECEIVER_EMAIL_ADDRESS = email
	SEAT_COUNT = seatCount
	SEAT_TYPE_ARRAY = seatTypes
	SPECIFIC_TRAIN_ARRAY = trains
	if goToBookPage {
		GO_TO_BOOK_PAGE = 1
	} else {
		GO_TO_BOOK_PAGE = 0
	}
}

func GenerateURL() string {
	return constants.BASE_URL + constants.FROM_KEY + FROM + constants.TO_KEY + TO + constants.DATE_KEY + DATE + constants.CLASS_KEY + SEAT_TYPE_ARRAY[0]
}
