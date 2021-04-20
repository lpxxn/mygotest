package demo1

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func TestName(t *testing.T) {
	//actual, err := ioutil.ReadFile("active.en.toml")
	//if err != nil {
	//	t.Fatal(err)
	//}

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("active.en.toml")
	bundle.MustLoadMessageFile("active.zh-cn.toml")

	lizer := i18n.NewLocalizer(bundle, "zh-cn")

	str := lizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:   "MyUnreadEmails",
		PluralCount: 1,
	})
	t.Log(str)

	s, err := lizer.Localize(&i18n.LocalizeConfig{MessageID: "abcde"})
	t.Log(s, err)

	str = lizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:      "MEALPLAN_TWO",
	})
	t.Log(str)
}
