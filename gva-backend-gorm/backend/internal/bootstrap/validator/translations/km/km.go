package km

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/locales"
	ut "github.com/go-playground/universal-translator"

	"github.com/go-playground/validator/v10"
)

// RegisterDefaultTranslations registers a set of default translations
// for all built in tag's in validator; you may add your own as desired.
func RegisterDefaultTranslations(v *validator.Validate, trans ut.Translator) (err error) {
	translations := []struct {
		tag             string
		translation     string
		override        bool
		customRegisFunc validator.RegisterTranslationsFunc
		customTransFunc validator.TranslationFunc
	}{
		{
			tag:         "required",
			translation: "{0} គឺជាវាលដែលត្រូវការ",
			override:    false,
		},
		{
			tag:         "required_if",
			translation: "{0} គឺជាវាលដែលត្រូវការ",
			override:    false,
		},
		{
			tag:         "required_unless",
			translation: "{0} គឺជាវាលដែលត្រូវការ",
			override:    false,
		},
		{
			tag:         "required_with",
			translation: "{0} គឺជាវាលដែលត្រូវការ",
			override:    false,
		},
		{
			tag:         "required_with_all",
			translation: "{0} គឺជាវាលដែលត្រូវការ",
			override:    false,
		},
		{
			tag:         "required_without",
			translation: "{0} គឺជាវាលដែលត្រូវការ",
			override:    false,
		},
		{
			tag:         "required_without_all",
			translation: "{0} គឺជាវាលដែលត្រូវការ",
			override:    false,
		},
		{
			tag:         "excluded_if",
			translation: "{0} គឺជាវាលដែលមិនរាប់បញ្ចូល",
			override:    false,
		},
		{
			tag:         "excluded_unless",
			translation: "{0} គឺជាវាលដែលមិនរាប់បញ្ចូល",
			override:    false,
		},
		{
			tag:         "excluded_with",
			translation: "{0} គឺជាវាលដែលមិនរាប់បញ្ចូល",
			override:    false,
		},
		{
			tag:         "excluded_with_all",
			translation: "{0} គឺជាវាលដែលមិនរាប់បញ្ចូល",
			override:    false,
		},
		{
			tag:         "excluded_without",
			translation: "{0} គឺជាវាលដែលមិនរាប់បញ្ចូល",
			override:    false,
		},
		{
			tag:         "excluded_without_all",
			translation: "{0} គឺជាវាលដែលមិនរាប់បញ្ចូល",
			override:    false,
		},
		{
			tag:         "isdefault",
			translation: "{0} ត្រូវតែជាតម្លៃលំនាំដើម",
			override:    false,
		},
		{
			tag: "len",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("len-string", "{0} ត្រូវតែមានប្រវែង {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("len-string-character", "{0} តួអក្សរ", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("len-string-character", "{0} តួអក្សរ", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("len-number", "{0} ត្រូវតែស្មើនឹង {1}", false); err != nil {
					return
				}

				if err = ut.Add("len-items", "{0} ត្រូវតែមាន {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("len-items-item", "ធាតុ {0}", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("len-items-item", "ធាតុ {0}", locales.PluralRuleOther, false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string

				var digits uint64
				var kind reflect.Kind

				if idx := strings.Index(fe.Param(), "."); idx != -1 {
					digits = uint64(len(fe.Param()[idx+1:]))
				}

				f64, err := strconv.ParseFloat(fe.Param(), 64)
				if err != nil {
					goto END
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					c, err = ut.C("len-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("len-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					c, err = ut.C("len-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("len-items", fe.Field(), c)

				default:
					t, err = ut.T("len-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "min",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("min-string", "{0} ត្រូវតែមានប្រវែងយ៉ាងតិច {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("min-string-character", "{0} តួអក្សរ", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("min-string-character", "{0} តួអក្សរ", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("min-number", "{0} ត្រូវតែ {1} ឬធំជាងនេះ។", false); err != nil {
					return
				}

				if err = ut.Add("min-items", "{0} ត្រូវតែមានយ៉ាងតិច {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("min-items-item", "ធាតុ {0}", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("min-items-item", "ធាតុ {0}", locales.PluralRuleOther, false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string

				var digits uint64
				var kind reflect.Kind

				if idx := strings.Index(fe.Param(), "."); idx != -1 {
					digits = uint64(len(fe.Param()[idx+1:]))
				}

				f64, err := strconv.ParseFloat(fe.Param(), 64)
				if err != nil {
					goto END
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					c, err = ut.C("min-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("min-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					c, err = ut.C("min-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("min-items", fe.Field(), c)

				default:
					t, err = ut.T("min-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "max",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("max-string", "{0} must be a maximum of {1} in length", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("max-string-character", "{0} តួអក្សរ", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("max-string-character", "{0} តួអក្សរs", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("max-number", "{0} ត្រូវតែ {1} ឬតិចជាងនេះ។", false); err != nil {
					return
				}

				if err = ut.Add("max-items", "{0} ត្រូវតែមានអតិបរមា {1}", false); err != nil {
					return
				}
				// if err = ut.AddCardinal("max-items-item", "ធាតុ {0}", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("max-items-item", "ធាតុ {0}", locales.PluralRuleOther, false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string

				var digits uint64
				var kind reflect.Kind

				if idx := strings.Index(fe.Param(), "."); idx != -1 {
					digits = uint64(len(fe.Param()[idx+1:]))
				}

				f64, err := strconv.ParseFloat(fe.Param(), 64)
				if err != nil {
					goto END
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					c, err = ut.C("max-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("max-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					c, err = ut.C("max-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("max-items", fe.Field(), c)

				default:
					t, err = ut.T("max-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "eq",
			translation: "{0} is not equal to {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ne",
			translation: "{0} should not be equal to {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "lt",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("lt-string", "{0} ត្រូវតែតិចជាង {1} in length", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("lt-string-character", "{0} តួអក្សរ", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("lt-string-character", "{0} តួអក្សរs", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-number", "{0} ត្រូវតែតិចជាង {1}", false); err != nil {
					return
				}

				if err = ut.Add("lt-items", "{0} ត្រូវតែមានតិចជាង {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("lt-items-item", "ធាតុ {0}", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("lt-items-item", "ធាតុ {0}", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-datetime", "{0} ត្រូវតែតិចជាងកាលបរិច្ឆេទ និងពេលវេលាបច្ចុប្បន្ន", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				fn := func() (err error) {
					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lt-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lt-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lt-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lt-items", fe.Field(), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
						goto END
					}

					t, err = ut.T("lt-datetime", fe.Field())

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("lt-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "lte",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("lte-string", "{0} ត្រូវតែមានប្រវែងអតិបរមា {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("lte-string-character", "{0} តួអក្សរ", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("lte-string-character", "{0} តួអក្សរs", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-number", "{0} ត្រូវតែ {1} ឬតិចជាងនេះ។", false); err != nil {
					return
				}

				if err = ut.Add("lte-items", "{0} ត្រូវតែមានអតិបរមា {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("lte-items-item", "ធាតុ {0}", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("lte-items-item", "ធាតុ {0}", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-datetime", "{0} ត្រូវតែតិចជាង ឬស្មើនឹងកាលបរិច្ឆេទ និងពេលវេលាបច្ចុប្បន្ន", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				fn := func() (err error) {
					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lte-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lte-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lte-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lte-items", fe.Field(), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
						goto END
					}

					t, err = ut.T("lte-datetime", fe.Field())

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("lte-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "gt",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("gt-string", "{0} ត្រូវតែធំជាង {1} in length", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("gt-string-character", "{0} តួអក្សរ", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("gt-string-character", "{0} តួអក្សរs", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-number", "{0} ត្រូវតែធំជាង {1}", false); err != nil {
					return
				}

				if err = ut.Add("gt-items", "{0} must contain more than {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("gt-items-item", "ធាតុ {0}", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("gt-items-item", "ធាតុ {0}", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-datetime", "{0} ត្រូវតែធំជាងកាលបរិច្ឆេទ និងពេលវេលាបច្ចុប្បន្ន", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				fn := func() (err error) {
					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gt-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gt-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gt-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gt-items", fe.Field(), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
						goto END
					}

					t, err = ut.T("gt-datetime", fe.Field())

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("gt-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "gte",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("gte-string", "{0} ត្រូវតែមានប្រវែងយ៉ាងតិច {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("gte-string-character", "{0} តួអក្សរ", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("gte-string-character", "{0} តួអក្សរs", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-number", "{0} ត្រូវតែ {1} ឬធំជាងនេះ។", false); err != nil {
					return
				}

				if err = ut.Add("gte-items", "{0} ត្រូវតែមានយ៉ាងតិច {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("gte-items-item", "ធាតុ {0}", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("gte-items-item", "ធាតុ {0}", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-datetime", "{0} must be greater than or equal to the current Date & Time", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				fn := func() (err error) {
					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gte-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gte-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gte-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gte-items", fe.Field(), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
						goto END
					}

					t, err = ut.T("gte-datetime", fe.Field())

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("gte-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "eqfield",
			translation: "{0} ត្រូវតែស្មើនឹង {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "eqcsfield",
			translation: "{0} ត្រូវតែស្មើនឹង {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "necsfield",
			translation: "{0} មិនអាចស្មើនឹង {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtcsfield",
			translation: "{0} ត្រូវតែធំជាង {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtecsfield",
			translation: "{0} ត្រូវតែធំជាង ឬស្មើ {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltcsfield",
			translation: "{0} ត្រូវតែតិចជាង {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltecsfield",
			translation: "{0} ត្រូវតែតិចជាង ឬស្មើ {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "nefield",
			translation: "{0} មិនអាចស្មើនឹង {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtfield",
			translation: "{0} ត្រូវតែធំជាង {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtefield",
			translation: "{0} ត្រូវតែធំជាង ឬស្មើ {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltfield",
			translation: "{0} ត្រូវតែតិចជាង {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltefield",
			translation: "{0} ត្រូវតែតិចជាង ឬស្មើ {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "alpha",
			translation: "{0} អាច​មាន​តែ​តួអក្សរ​អក្ខរក្រម​ប៉ុណ្ណោះ។",
			override:    false,
		},
		{
			tag:         "alphanum",
			translation: "{0} ត្រូវតែមានតែតួអក្សរលេខប៉ុណ្ណោះ។",
			override:    false,
		},
		{
			tag:         "numeric",
			translation: "{0} ត្រូវតែជាតម្លៃលេខគត់ទុក្ខជាក់ស្ដែង",
			override:    false,
		},
		{
			tag:         "number",
			translation: "{0} ត្រូវតែជាតម្លៃលេខគត់ទុក្ខជាក់ស្ដែង",
			override:    false,
		},
		{
			tag:         "hexadecimal",
			translation: "{0} ត្រូវតែជាតម្លៃលេខអក្សរប៉ុណ្ណោះ។",
			override:    false,
		},
		{
			tag:         "hexcolor",
			translation: "{0} ត្រូវតែជាតម្លៃពណ៌ HEX ប៉ុណ្ណោះ។",
			override:    false,
		},
		{
			tag:         "rgb",
			translation: "{0} ត្រូវតែជាតម្លៃពណ៌ RGB ប៉ុណ្ណោះ។",
			override:    false,
		},
		{
			tag:         "rgba",
			translation: "{0} ត្រូវតែជាតម្លៃពណ៌ RGBA ប៉ុណ្ណោះ។",
			override:    false,
		},
		{
			tag:         "hsl",
			translation: "{0} ត្រូវតែជាតម្លៃពណ៌ HSL ប៉ុណ្ណោះ។",
			override:    false,
		},
		{
			tag:         "hsla",
			translation: "{0} ត្រូវតែជាតម្លៃពណ៌ HSLA ប៉ុណ្ណោះ។",
			override:    false,
		},
		{
			tag:         "e164",
			translation: "{0} ត្រូវតែជាតម្លៃលេខទូរស័ព្ទ E.164 ប៉ុណ្ណោះ។",
			override:    false,
		},
		{
			tag:         "email",
			translation: "{0} ត្រូវតែជាអាសយដ្ឋានអ៊ីមែលប៉ុណ្ណោះ។",
			override:    false,
		},
		{
			tag:         "url",
			translation: "{0} ត្រូវតែជាតម្លៃ URL ប៉ុណ្ណោះ។",
			override:    false,
		},
		{
			tag:         "uri",
			translation: "{0} ត្រូវតែជាតម្លៃ URI ប៉ុណ្ណោះ។",
			override:    false,
		},
		{
			tag:         "base64",
			translation: "{0} ត្រូវតែជាតម្លៃអក្សរ Base64 ប៉ុណ្ណោះ។",
			override:    false,
		},
		{
			tag:         "contains",
			translation: "{0} ត្រូវតែមានអត្ថបទ '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "containsany",
			translation: "{0} ត្រូវតែមានយ៉ាងហោចណាស់មួយក្នុងចំនោមតួអក្សរខាងក្រោម '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "excludes",
			translation: "{0} មិនអាចមានអត្ថបទ '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "excludesall",
			translation: "{0} មិនអាចមានតួអក្សរ '{1}' ណាមួយខាងក្រោម",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "excludesrune",
			translation: "{0} មិនអាចមាន '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "isbn",
			translation: "{0} ត្រូវតែជាកូដ ISBN ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "isbn10",
			translation: "{0} ត្រូវតែជាកូដ ISBN-10 ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "isbn13",
			translation: "{0} ត្រូវតែជាកូដ ISBN-13 ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "issn",
			translation: "{0} ត្រូវតែជាកូដ ISSN ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "uuid",
			translation: "{0} ត្រូវតែជាកូដ UUID ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "uuid3",
			translation: "{0} ត្រូវតែជាកូដ UUID កំណែ 3 ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "uuid4",
			translation: "{0} ត្រូវតែជាកូដ UUID កំណែ 4 ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "uuid5",
			translation: "{0} ត្រូវតែជាកូដ UUID កំណែ 5 ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "ulid",
			translation: "{0} ត្រូវតែជាកូដ ULID ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "ascii",
			translation: "{0} ត្រូវតែមានតែតួអក្សរ ascii ប៉ុណ្ណោះ",
			override:    false,
		},
		{
			tag:         "printascii",
			translation: "{0} ត្រូវតែមានតែតួអក្សរអាចបោះពុម្ព ascii ប៉ុណ្ណោះ",
			override:    false,
		},
		{
			tag:         "multibyte",
			translation: "{0} ត្រូវតែមានតួអក្សរសម្រាប់ multibyte",
			override:    false,
		},
		{
			tag:         "datauri",
			translation: "{0} ត្រូវតែមាន Data URI ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "latitude",
			translation: "{0} ត្រូវតែមានទីតាំងលើទទឹងត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "longitude",
			translation: "{0} ត្រូវតែមានទីតាំងបណ្តោយត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "ssn",
			translation: "{0} ត្រូវតែជាប៉ាសស្នូតលេខសង្គមត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "ipv4",
			translation: "{0} ត្រូវតែជាអាសយដ្ឋាន IPv4 ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "ipv6",
			translation: "{0} ត្រូវតែជាអាសយដ្ឋាន IPv6 ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "ip",
			translation: "{0} ត្រូវតែជាអាសយដ្ឋាន IP ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "cidr",
			translation: "{0} ត្រូវតែមានចំនួន CIDR ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "cidrv4",
			translation: "{0} ត្រូវតែមានចំនួន CIDR សម្រាប់អាសយដ្ឋាន IPv4 ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "cidrv6",
			translation: "{0} ត្រូវតែមានចំនួន CIDR សម្រាប់អាសយដ្ឋាន IPv6 ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "tcp_addr",
			translation: "{0} ត្រូវតែជាអាសយដ្ឋាន TCP ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "tcp4_addr",
			translation: "{0} ត្រូវតែជាអាសយដ្ឋាន TCP IPv4 ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "tcp6_addr",
			translation: "{0} ត្រូវតែជាអាសយដ្ឋាន TCP IPv6 ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "udp_addr",
			translation: "{0} ត្រូវតែជាអាសយដ្ឋាន UDP ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "udp4_addr",
			translation: "{0} ត្រូវតែជាអាសយដ្ឋាន UDP IPv4 ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "udp6_addr",
			translation: "{0} ត្រូវតែជាអាសយដ្ឋាន UDP IPv6 ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "ip_addr",
			translation: "{0} ត្រូវតែជាអាសយដ្ឋាន IP ដែលអាចដោះស្រាយបាន",
			override:    false,
		},
		{
			tag:         "ip4_addr",
			translation: "{0} ត្រូវតែជាអាសយដ្ឋាន IPv4 ដែលអាចដោះស្រាយបាន",
			override:    false,
		},
		{
			tag:         "ip6_addr",
			translation: "{0} ត្រូវតែជាអាសយដ្ឋាន IPv6 ដែលអាចដោះស្រាយបាន",
			override:    false,
		},
		{
			tag:         "unix_addr",
			translation: "{0} ត្រូវតែជាអាសយដ្ឋាន UNIX ដែលអាចដោះស្រាយបាន",
			override:    false,
		},
		{
			tag:         "mac",
			translation: "{0} ត្រូវតែមានអាសយដ្ឋាន MAC ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "fqdn",
			translation: "{0} ត្រូវតែជាឈ្មោះ FQDN ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "unique",
			translation: "{0} ត្រូវតែមានតម្លៃតែមួយ",
			override:    false,
		},
		{
			tag:         "iscolor",
			translation: "{0} ត្រូវតែជាពណ៌ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "cron",
			translation: "{0} ត្រូវតែជាការបញ្ចូល cron ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "oneof",
			translation: "{0} ត្រូវតែជាផ្នែកមួយនៃ [{1}]",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				s, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}
				return s
			},
		},
		{
			tag:         "json",
			translation: "{0} ត្រូវតែជាខ្សែអក្សររូបភាព json ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "jwt",
			translation: "{0} ត្រូវតែជាខ្សែអក្សររូបភាព jwt ត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "lowercase",
			translation: "{0} ត្រូវតែជាខ្សែអក្សរអក្សរតូច",
			override:    false,
		},
		{
			tag:         "uppercase",
			translation: "{0} ត្រូវតែជាខ្សែអក្សរអក្សរធំ",
			override:    false,
		},
		{
			tag:         "datetime",
			translation: "{0} មិនត្រឹមត្រូវទៅនឹងរបៀបបញ្ចូលនៅក្នុង {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "postcode_iso3166_alpha2",
			translation: "{0} មិនត្រឹមត្រូវទៅនឹងរបៀបបញ្ចូលនៅក្នុងប្រទេស {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "postcode_iso3166_alpha2_field",
			translation: "{0} មិនត្រឹមត្រូវទៅនឹងរបៀបបញ្ចូលនៅក្នុងប្រទេសនៅក្នុងវិសាលគណនី {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "boolean",
			translation: "{0} ត្រូវតែជាតម្លៃពិតឬមិនពិត",
			override:    false,
		},
		{
			tag:         "image",
			translation: "{0} ត្រូវតែជារូបភាពត្រឹមត្រូវ",
			override:    false,
		},
		{
			tag:         "cve",
			translation: "{0} ត្រូវតែជាអត្ថបទវិទ្យាសាស្ត្រត្រឹមត្រូវ",
			override:    false,
		},
	}

	for _, t := range translations {

		if t.customTransFunc != nil && t.customRegisFunc != nil {
			err = v.RegisterTranslation(t.tag, trans, t.customRegisFunc, t.customTransFunc)
		} else if t.customTransFunc != nil && t.customRegisFunc == nil {
			err = v.RegisterTranslation(t.tag, trans, registrationFunc(t.tag, t.translation, t.override), t.customTransFunc)
		} else if t.customTransFunc == nil && t.customRegisFunc != nil {
			err = v.RegisterTranslation(t.tag, trans, t.customRegisFunc, translateFunc)
		} else {
			err = v.RegisterTranslation(t.tag, trans, registrationFunc(t.tag, t.translation, t.override), translateFunc)
		}

		if err != nil {
			return
		}
	}

	return
}

func registrationFunc(tag string, translation string, override bool) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) (err error) {
		if err = ut.Add(tag, translation, override); err != nil {
			return
		}

		return
	}
}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		log.Printf("warning: error translating FieldError: %#v", fe)
		return fe.(error).Error()
	}

	return t
}
