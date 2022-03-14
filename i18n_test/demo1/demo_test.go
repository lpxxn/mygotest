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
	bundle.MustLoadMessageFile("active.zh-cn.toml") // 中文只有other

	lizer := i18n.NewLocalizer(bundle, "zh-cn")

	str := lizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:   "MyUnreadEmails",
		PluralCount: 1,
	})
	t.Log(str)

	s, err := lizer.Localize(&i18n.LocalizeConfig{MessageID: "abcde"})
	t.Log(s, err)

	str = lizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "MEALPLAN_TWO",
	})
	t.Log(str)

	str = lizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    "PersonUnreadEmails",
		PluralCount:  2,
		TemplateData: map[string]interface{}{"Name": "Nick", "UnreadEmailCount": 3},
	})
	t.Log(str)

	str = lizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    "PersonUnreadEmails",
		PluralCount:  1,
		TemplateData: map[string]interface{}{"Name": "Nick", "UnreadEmailCount": 1},
	})
	t.Log(str)

}

func Test1(t *testing.T) {
	bundle := i18n.NewBundle(language.English)
	localizer := i18n.NewLocalizer(bundle, "en")
	personCatsMessage := &i18n.Message{
		ID:    "PersonCats",
		One:   "{{.Name}} has {{.Count}} cat.",
		Other: "{{.Name}} has {{.Count}} cats.",
	}
	bundle.AddMessages(language.English, personCatsMessage)
	t.Log(localizer.MustLocalize(&i18n.LocalizeConfig{
		//DefaultMessage: personCatsMessage,
		MessageID:   "PersonCats",
		PluralCount: 1,
		TemplateData: map[string]interface{}{
			"Name":  "Nick",
			"Count": 1,
		},
	}))
	t.Log(localizer.MustLocalize(&i18n.LocalizeConfig{
		//DefaultMessage: personCatsMessage,
		MessageID:   "PersonCats",
		PluralCount: 2,
		TemplateData: map[string]interface{}{
			"Name":  "Nick",
			"Count": 2,
		},
	}))
	t.Log(localizer.MustLocalize(&i18n.LocalizeConfig{
		//DefaultMessage: personCatsMessage,
		MessageID:   "PersonCats",
		PluralCount: "2.5",
		TemplateData: map[string]interface{}{
			"Name":  "Nick",
			"Count": "2.5",
		},
	}))
	// Output:
	// Nick has 1 cat.
	// Nick has 2 cats.
	// Nick has 2.5 cats
}
