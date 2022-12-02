package year2020

import (
	"regexp"
	"strconv"
	"strings"
)

type Day04 struct{}

func (p Day04) PartA(lines []string) any {
	passports := getPassports(lines)
	valid := 0
	for _, passport := range passports {
		if validatePassport(passport, false) {
			valid++
		}
	}
	return valid
}

func (p Day04) PartB(lines []string) any {
	passports := getPassports(lines)
	valid := 0
	for _, passport := range passports {
		if validatePassport(passport, true) {
			valid++
		}
	}
	return valid
}

func validatePassport(passport map[string]string, extended bool) bool {
	requiredFields := []string{
		"byr", // Birth year
		"iyr", // (Issue Year)
		"eyr", // (Expiration Year)
		"hgt", // (Height)
		"hcl", // (Hair Color)
		"ecl", // (Eye Color)
		"pid", // (Passport ID)
		//"cid",// (country id)
	}
	for _, fieldName := range requiredFields {
		if passport[fieldName] == "" {
			return false
		}
	}
	if extended {
		//byr (Birth Year) - four digits; at least 1920 and at most 2002.
		if !validateNumber(passport["byr"], 1920, 2002) ||
			!validateNumber(passport["iyr"], 2010, 2020) ||
			!validateNumber(passport["eyr"], 2020, 2030) ||
			!validateHeight(passport["hgt"]) ||
			!validateHairColour(passport["hcl"]) ||
			!validateEyeColour(passport["ecl"]) ||
			!validatePid(passport["pid"]) {
			return false
		}
		//iyr (Issue Year) - four digits; at least 2010 and at most 2020.
		//eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
		//hgt (Height) - a number followed by either cm or in:
		//If cm, the number must be at least 150 and at most 193.
		//If in, the number must be at least 59 and at most 76.
		//hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
		//	ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
		//	pid (Passport ID) - a nine-digit number, including leading zeroes.
	}
	return true
}

func validatePid(pid string) bool {
	expr := regexp.MustCompile("^[0-9]{9}$")
	return expr.Match([]byte(pid))
}

func validateHairColour(colour string) bool {
	expr := regexp.MustCompile("^#[0-9a-f]{6}$")
	return expr.Match([]byte(colour))
}

func validateEyeColour(colour string) bool {
	validColours := []string{
		"amb", "blu", "brn", "gry", "grn", "hzl", "oth",
	}
	for _, col := range validColours {
		if colour == col {
			return true
		}
	}
	return false
}

func validateNumber(number string, min int, max int) bool {
	n, err := strconv.Atoi(number)
	if err != nil {
		return false
	}
	return n >= min && n <= max
}

func validateHeight(height string) bool {
	expr := regexp.MustCompile("^([0-9]+)(cm|in)$")
	matches := expr.FindStringSubmatch(height)
	if len(matches) < 3 {
		return false
	}
	n, err := strconv.Atoi(matches[1])
	if err != nil {
		return false
	}
	switch matches[2] {
	case "cm":
		return n >= 150 && n <= 193
	case "in":
		return n >= 59 && n <= 76
	}
	return false
}

func getPassports(lines []string) (passports []map[string]string) {
	passport := make(map[string]string)
	for _, line := range lines {
		if line == "" {
			passports = append(passports, passport)
			passport = make(map[string]string)
			continue
		}
		fields := strings.Split(line, " ")
		for _, field := range fields {
			f := strings.Split(field, ":")
			if f[0] != "cid" {
				passport[f[0]] = f[1]
			}
		}
	}
	passports = append(passports, passport)
	return
}
