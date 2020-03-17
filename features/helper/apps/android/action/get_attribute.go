package androidaction

import (
	"fmt"
	"log"

	"github.com/golang-automation/features/helper"
	android "github.com/golang-automation/features/helper/apps/android"
)

/*GetAttributeByXpath element visible by Xpath selector*/
func GetAttributeByXpath(locator string, attr string) error {
	_, err := android.Device.FindByXPath(locator).Attribute(attr)
	if err != nil {
		log.Panicln(fmt.Errorf("REASON: %s", err))
	}

	return nil
}

/*GetAttributeByButton element visible by button*/
func GetAttributeByButton(locator string, attr string) error {
	_, err := android.Device.FindByButton(locator).Attribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByClass element visible by class*/
func GetAttributeByClass(locator string, attr string) error {
	_, err := android.Device.FindByClass(locator).Attribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByID element visible by ID*/
func GetAttributeByID(locator string, attr string) error {
	_, err := android.Device.FindByID(locator).Attribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByLabel element visible by label*/
func GetAttributeByLabel(locator string, attr string) error {
	_, err := android.Device.FindByLabel(locator).Attribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByLink element visible by link*/
func GetAttributeByLink(locator string, attr string) error {
	_, err := android.Device.FindByLink(locator).Attribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByName element visible by class name*/
func GetAttributeByName(locator string, attr string) error {
	_, err := android.Device.FindByName(locator).Attribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByText element visible by Xpath selector*/
func GetAttributeByText(locator string, attr string) error {
	_, err := android.Device.FindByXPath("//*[contains(@text, '" + locator + "')]").Attribute(attr)
	helper.LogPanicln(err)

	return nil
}