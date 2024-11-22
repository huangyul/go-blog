package validator

import (
	"regexp"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTran "github.com/go-playground/validator/v10/translations/en"
	chTran "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

func init() {
	local := "zh"
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New()
		enT := en.New()
		uni := ut.New(enT, zhT, enT)

		var ok bool
		trans, ok = uni.GetTranslator("zh")

		if !ok {
			panic("uni.GetTranslator('zh') err")
		}
		var err error
		switch local {
		case "zh":
			err = chTran.RegisterDefaultTranslations(v, trans)
		case "en":
			err = enTran.RegisterDefaultTranslations(v, trans)
		default:
			err = enTran.RegisterDefaultTranslations(v, trans)
		}
		if err != nil {
			panic(err)
		}
	}
}

func Translate(err error) string {

	if err == nil {
		return ""
	}

	typ, ok := err.(validator.ValidationErrors)

	if !ok {
		return "no json structure"
	}

	raw := typ.Translate(trans)

	for _, v := range raw {
		return convertToSnakeCase(v)
	}
	return ""
}

func pascalToSnake(s string) string {
	re := regexp.MustCompile("([A-Z])")
	snake := re.ReplaceAllStringFunc(s, func(match string) string {
		return "_" + strings.ToLower(match)
	})

	if len(snake) > 0 && snake[0] == '_' {
		snake = snake[1:]
	}

	return snake
}

func convertToSnakeCase(input string) string {
	wordRegex := regexp.MustCompile(`[A-Za-z]+`)

	result := wordRegex.ReplaceAllStringFunc(input, func(word string) string {
		re := regexp.MustCompile("([A-Z])")
		snake := re.ReplaceAllStringFunc(word, func(match string) string {
			return "_" + strings.ToLower(match)
		})
		if len(snake) > 0 && snake[0] == '_' {
			snake = snake[1:]
		}
		return snake
	})

	return result
}
