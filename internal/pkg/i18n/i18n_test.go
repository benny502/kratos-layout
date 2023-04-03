package i18n

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func TestI18NEnglish(t *testing.T) {
	bundle, err := NewBundle()
	if err != nil {
		t.Log(err)
	}
	localizer := NewLocalizer(bundle, "en")
	message := localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "TEAM_10004",
	})
	if err != nil {
		t.Log(err)
	}
	fmt.Println(message)
	assert.Equal(t, message, "Request failed.team does not exist .")
	//panic("boom")
}
