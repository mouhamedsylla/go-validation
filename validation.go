// Package validation provides a framework for validating struct fields based on custom rules.
package validation

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// RuleSet represents a set of validation rules for a specific field.
type RuleSet struct {
	Name         string                 // Name of the rule.
	RuleValue    any                    // Value associated with the rule (e.g., length, min/max value).
	FieldValue   any                    // Value of the field being validated.
	FieldName    string                 // Name of the field being validated.
	MessageFunc  func(*RuleSet) string  // Function to generate an error message if validation fails.
	ValidateFunc func(*RuleSet) bool    // Function to validate the field value against the rule.
}

// Target represents a field to be validated, including its name, value, and associated validation tags.
type Target struct {
	Name  string   // Name of the field.
	Value any      // Value of the field.
	Tag   []string // Validation tags for the field.
}

// Validator contains a list of Targets to be validated.
type Validator struct {
	Targets []*Target // List of targets to be validated.
}

// NewValidator creates a new instance of Validator.
func NewValidator() *Validator {
	return &Validator{}
}

// NewTarget creates a new Target with the given name, value, and tags.
func NewTarget(name string, value any, tags []string) *Target {
	return &Target{
		Name:  name,
		Value: value,
		Tag:   tags,
	}
}

// GetRuleByTagRuleName retrieves the RuleSet corresponding to the provided tag rule name.
func (tg *Target) GetRuleByTagRuleName(tagRule string) (*RuleSet, error) {
	for _, rule := range Rules { // Iterate through available rules.
		ruleset := rule()
		switch {
		case strings.Contains(tagRule, "minL") && ruleset.Name == "minLength",
			strings.Contains(tagRule, "maxL") && ruleset.Name == "maxLength",
			strings.Contains(tagRule, "min") && ruleset.Name == "min",
			strings.Contains(tagRule, "max") && ruleset.Name == "max":
			// Extract numeric value from tag and set as RuleValue.
			value, err := GetNum(tagRule)
			if err != nil {
				log.Fatalln(err)
			}
			ruleset.RuleValue = value
			return ruleset, nil
		case ruleset.Name == tagRule:
			// Direct match of rule name.
			return ruleset, nil
		}
	}
	return nil, errors.New("rule not found") // Return error if rule is not found.
}

// Init initializes the Validator with the fields of the provided models.
func (v *Validator) Init(models ...interface{}) {
	for _, m := range models {
		val := reflect.ValueOf(m)
		for i := 0; i < val.NumField(); i++ { // Iterate through fields of the model.
			field := val.Field(i)
			target := NewTarget(
				val.Type().Field(i).Name,
				field.Interface(),
				strings.Split(val.Type().Field(i).Tag.Get("validate"), " "),
			)
			v.Targets = append(v.Targets, target)
		}
	}
}

// Validate checks all targets against their respective validation rules.
func (v *Validator) Validate() error {
	for _, target := range v.Targets {
		for _, tag := range target.Tag { // Iterate through each validation tag of the target.
			ruleset, err := target.GetRuleByTagRuleName(tag)
			if err != nil {
				fmt.Printf("Warning: %s for %s\n", err, target.Name) // Log warning if rule is not found.
				continue
			}
			ruleset.FieldValue = target.Value // Set the field value for the rule.
			ruleset.FieldName = target.Name   // Set the field name for the rule.
			if !ruleset.ValidateFunc(ruleset) { // Perform validation.
				return errors.New(ruleset.MessageFunc(ruleset)) // Return error if validation fails.
			}
		}
	}
	return nil
}

// GetNum extracts a number from the given input string.
func GetNum(input string) (int, error) {
	re := regexp.MustCompile(`\d+`)
	match := re.FindString(input) // Find the first numeric substring.
	if match == "" {
		return 0, fmt.Errorf("no number found in the input string")
	}
	return strconv.Atoi(match) // Convert the extracted substring to an integer.
}