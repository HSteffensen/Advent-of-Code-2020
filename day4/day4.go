package day4

import (
	"regexp"
	"strconv"
)

func IsValidPassport(passport string) bool {
	pattern := `(\S+):(\S+)`
	matches := regexp.MustCompile(pattern).FindAllStringSubmatch(passport, -1)
	fieldMap := make(map[string]string)
	for _, match := range matches {
		fieldMap[match[1]] = match[2]
	}
	wantFields := [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, field := range wantFields {
		if _, ok := fieldMap[field]; !ok {
			return false
		}
	}
	return true
}

var PassportValidators = map[string]func(string) bool{
	"byr": func(value string) bool {
		val, err := strconv.Atoi(value)
		if err != nil || len(value) != 4 {
			return false
		}
		return val >= 1920 && val <= 2002
	},
	"iyr": func(value string) bool {
		val, err := strconv.Atoi(value)
		if err != nil || len(value) != 4 {
			return false
		}
		return val >= 2010 && val <= 2020
	},
	"eyr": func(value string) bool {
		val, err := strconv.Atoi(value)
		if err != nil || len(value) != 4 {
			return false
		}
		return val >= 2020 && val <= 2030
	},
	"hgt": func(value string) bool {
		val, err := strconv.Atoi(value[:len(value)-2])
		if err != nil {
			return false
		}
		switch units := value[len(value)-2:]; units {
		case "cm":
			return val >= 150 && val <= 193
		case "in":
			return val >= 59 && val <= 76
		}
		return false
	},
	"hcl": func(value string) bool {
		result, err := regexp.MatchString(`^#[0-9a-f]{6}$`, value)
		return result && err == nil
	},
	"ecl": func(value string) bool {
		allowed := map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}
		_, ok := allowed[value]
		return ok
	},
	"pid": func(value string) bool {
		result, err := regexp.MatchString(`^\d{9}$`, value)
		return result && err == nil
	},
}

func IsValidPassportValues(passport string) bool {
	if !IsValidPassport(passport) {
		return false
	}
	pattern := `(\S+):(\S+)`
	matches := regexp.MustCompile(pattern).FindAllStringSubmatch(passport, -1)
	fieldMap := make(map[string]string)
	for _, match := range matches {
		fieldMap[match[1]] = match[2]
	}
	for field, validator := range PassportValidators {
		if value, ok := fieldMap[field]; !ok || !validator(value) {
			return false
		}
	}
	return true
}

func CountValidPassports(passports []string) int {
	count := 0
	for _, passport := range passports {
		if IsValidPassport(passport) {
			count++
		}
	}
	return count
}

func CountValidPassportsValues(passports []string) int {
	count := 0
	for _, passport := range passports {
		if IsValidPassportValues(passport) {
			count++
		}
	}
	return count
}
