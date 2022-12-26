package helper

import (
	"fmt"
	"unicode"
)

/**
 * Match atleast 1 lower case , 1 number and 1 special character
 * {domain}/admin/master/console/#/realms/{realm}/authentication/password-policy
 */
func ValidatePassword(password string) error {
next:
	for name, classes := range map[string][]*unicode.RangeTable{
		"LOWER_CASE":        {unicode.Lower},
		"NUMERIC":           {unicode.Number, unicode.Digit},
		"SPECIAL_CHARACTER": {unicode.Space, unicode.Symbol, unicode.Punct, unicode.Mark},
	} {
		for _, r := range password {
			if unicode.IsOneOf(classes, r) {
				continue next
			}
		}
		return fmt.Errorf("AT_LEAST_ONE_%s", name)
	}
	return nil
}
