package stepdefinitions

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/golang-automation/features/helper/apps/android"
	"github.com/golang-automation/features/helper/apps/ios"
	"github.com/golang-automation/features/helper/web"
	"github.com/golang-automation/features/helper/web/desktophelper"
	"github.com/golang-automation/features/helper/web/mobilehelper"
)

/*Username variabel*/
var Username string

/*Password variabel*/
var Password string

/*LoginData is function to set login data*/
func LoginData(user string, platform string) error {
	userENV := strings.HasPrefix(user, "ENV:")
	readENV := strings.TrimPrefix(user, "ENV:")

	if userENV {
		Username = os.Getenv(readENV + "_USERNAME")
		Password = os.Getenv(readENV + "_PASSWORD")
	} else {
		Username = os.Getenv(user + "_USERNAME")
		Password = os.Getenv(user + "_PASSWORD")
	}

	switch platform {
	case "desktop":
		web.DriverConnect()
		desktophelper.GoToURL(os.Getenv("URL_2"))
	case "mobile":
		web.DriverConnect()
		mobilehelper.GoToURL(os.Getenv("URL_2"))
	case "android":
		android.DriverConnect()
		android.OpenApps()
	case "ios":
		ios.DriverConnect()
		ios.OpenApps()
	default:
		log.Panicln(fmt.Errorf("REASON: Driver not found"))
	}

	return nil
}