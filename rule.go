// Package validation provides a set of rules for validating various input fields.
package validation

import (
	"fmt"
	"regexp"
)

// RuleFunc is a type definition for a function that returns a pointer to a RuleSet.
type RuleFunc func() *RuleSet

// Regular expressions for email and username validation.
var (
	emailRegex    = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	// Rules is a list of all the rule functions available for validation.
	Rules = []RuleFunc{
		MaxLength,
		MinLength,
		Email,
		Required,
		Username,
		MinNumeric,
		MaxNumeric,
	}
)

// Required returns a RuleSet for validating that a field is required (non-empty).
func Required() *RuleSet {
	return &RuleSet{
		Name: "required",
		MessageFunc: func(set *RuleSet) string {
			return fmt.Sprintf("%s is a required field", set.FieldName)
		},
		ValidateFunc: func(rule *RuleSet) bool {
			// Check if the field value is a non-empty string.
			str, ok := rule.FieldValue.(string)
			if !ok {
				return false
			}
			return len(str) > 0
		},
	}
}

// Email returns a RuleSet for validating that a field is a valid email address.
func Email() *RuleSet {
	return &RuleSet{
		Name: "email",
		MessageFunc: func(set *RuleSet) string {
			return "email address is invalid"
		},
		ValidateFunc: func(set *RuleSet) bool {
			// Check if the field value matches the email regex pattern.
			email, ok := set.FieldValue.(string)
			if !ok {
				return false
			}
			return emailRegex.MatchString(email)
		},
	}
}

// Username returns a RuleSet for validating that a field is a valid username.
func Username() *RuleSet {
	return &RuleSet{
		Name: "username",
		MessageFunc: func(set *RuleSet) string {
			return fmt.Sprintf("%s is not a valid username", set.FieldName)
		},
		ValidateFunc: func(set *RuleSet) bool {
			// Check if the field value is a valid username (alphanumeric and underscores, length 3-20).
			username, ok := set.FieldValue.(string)
			if !ok {
				return false
			}
			if len(username) < 3 || len(username) > 20 {
				return false
			}
			return usernameRegex.MatchString(username)
		},
	}
}

// MaxLength returns a RuleSet for validating that a field does not exceed a maximum length.
func MaxLength() *RuleSet {
	return &RuleSet{
		Name: "maxLength",
		ValidateFunc: func(set *RuleSet) bool {
			// Check if the field value length is less than or equal to the specified max length.
			n := set.RuleValue.(int)
			str, ok := set.FieldValue.(string)
			if !ok {
				return false
			}
			return len(str) <= n
		},
		MessageFunc: func(set *RuleSet) string {
			return fmt.Sprintf("%s should be maximum %d characters long", set.FieldName, set.RuleValue)
		},
	}
}

// MinLength returns a RuleSet for validating that a field meets a minimum length.
func MinLength() *RuleSet {
	return &RuleSet{
		Name: "minLength",
		ValidateFunc: func(set *RuleSet) bool {
			// Check if the field value length is greater than or equal to the specified min length.
			n := set.RuleValue.(int)
			str, ok := set.FieldValue.(string)
			if !ok {
				return false
			}
			return len(str) >= n
		},
		MessageFunc: func(set *RuleSet) string {
			return fmt.Sprintf("%s should be at least %d characters long", set.FieldName, set.RuleValue)
		},
	}
}

// MinNumeric returns a RuleSet for validating that a numeric field meets a minimum value.
func MinNumeric() *RuleSet {
	return &RuleSet{
		Name: "min",
		ValidateFunc: func(set *RuleSet) bool {
			// Check if the numeric field value is greater than or equal to the specified min value.
			n := set.RuleValue.(int)
			num, ok := set.FieldValue.(int)
			if !ok {
				return false
			}
			return num >= n
		},
		MessageFunc: func(set *RuleSet) string {
			return fmt.Sprintf("%s should be at least %d", set.FieldName, set.RuleValue)
		},
	}
}

// MaxNumeric returns a RuleSet for validating that a numeric field does not exceed a maximum value.
func MaxNumeric() *RuleSet {
	return &RuleSet{
		Name: "max",
		ValidateFunc: func(set *RuleSet) bool {
			// Check if the numeric field value is less than or equal to the specified max value.
			n := set.RuleValue.(int)
			num, ok := set.FieldValue.(int)
			if !ok {
				return false
			}
			return num <= n
		},
		MessageFunc: func(set *RuleSet) string {
			return fmt.Sprintf("%s should be at most %d", set.FieldName, set.RuleValue)
		},
	}
}