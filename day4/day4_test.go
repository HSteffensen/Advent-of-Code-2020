package day4_test

import (
	"HSteffensen/AoC2020/day4"
	"testing"
)

func TestReadPassports(t *testing.T) {
	t.Logf("%#v", day4.ReadPassports(day4.ExampleInput))
}

func TestIsValidPassport(t *testing.T) {
	data := day4.ReadPassports(day4.ExampleInput)
	result := [4]bool{}
	for i, _ := range result {
		result[i] = day4.IsValidPassport(data[i])
	}
	expected := [4]bool{true, false, true, false}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Passport %v should give %v from IsValidPassport", i, expected[i])
		}
	}
}

func TestCountValidPassports(t *testing.T) {
	if day4.CountValidPassports([]string{}) != 0 {
		t.Error("Empty passports list should have 0 valid.")
	}
	if c := day4.CountValidPassports(day4.ReadPassports(day4.ExampleInput)); c != 2 {
		t.Errorf("Example passports list should have 2 valid. Got %v", c)
	}
}

func TestPassportValidators(t *testing.T) {
	if !day4.PassportValidators["byr"]("2002") {
		t.Error("byr valid")
	}
	if day4.PassportValidators["byr"]("2003") {
		t.Error("byr invalid")
	}

	if !day4.PassportValidators["hgt"]("60in") {
		t.Error("hgt valid 1")
	}
	if !day4.PassportValidators["hgt"]("190cm") {
		t.Error("hgt valid 2")
	}
	if day4.PassportValidators["hgt"]("190in") {
		t.Error("hgt invalid 1")
	}
	if day4.PassportValidators["hgt"]("190") {
		t.Error("hgt invalid 2")
	}

	if !day4.PassportValidators["hcl"]("#123abc") {
		t.Error("hcl valid")
	}
	if day4.PassportValidators["hcl"]("#123abz") {
		t.Error("hcl invalid 1")
	}
	if day4.PassportValidators["hcl"]("123abc") {
		t.Error("hcl invalid 2")
	}

	if !day4.PassportValidators["ecl"]("brn") {
		t.Error("ecl valid")
	}
	if day4.PassportValidators["ecl"]("wat") {
		t.Error("ecl invalid")
	}

	if !day4.PassportValidators["pid"]("000000001") {
		t.Error("pid valid")
	}
	if day4.PassportValidators["pid"]("0123456789") {
		t.Error("pid invalid")
	}
}

func TestIsValidPassportValues(t *testing.T) {
	invalidPassports := `eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007`
	validPassports := `pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`
	invalidData := day4.ReadPassports(invalidPassports)
	validData := day4.ReadPassports(validPassports)
	for i, passport := range invalidData {
		if day4.IsValidPassportValues(passport) {
			t.Errorf("Invalid passport %v was incorrectly marked as valid", i)
		}
	}
	for i, passport := range validData {
		if !day4.IsValidPassportValues(passport) {
			t.Errorf("Valid passport %v was incorrectly marked as invalid", i)
		}
	}
}
