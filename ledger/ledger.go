// package ledger implements routines to generate a ledger report
package ledger

import (
	"errors"
	"fmt"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

// AddHeader returns the correct header string given locale
func AddHeader(locale string) (string, error) {
	switch locale {
	case "nl-NL":
		return fmt.Sprintf("%-10s | %-25s | Verandering\n", "Datum", "Omschrijving"), nil
	case "en-US":
		return fmt.Sprintf("%-10s | %-25s | Change\n", "Date", "Description"), nil
	}
	return "", errors.New("unknown locale")
}

// FormatDate returns the correct date formatted according to given locale
// and date data to be extracted from string d
func FormatDate(d string, locale string) (string, error) {
	var year, month, day int
	fmt.Sscanf(d, "%d-%d-%d", &year, &month, &day)
	switch {
	case month < 1 || month > 12:
		return "", errors.New("invalid month")
	case day < 1 || day > 31:
		return "", errors.New("invalid day")
	}
	switch locale {
	case "nl-NL":
		return fmt.Sprintf("%02d-%02d-%d", day, month, year), nil
	case "en-US":
		return fmt.Sprintf("%02d/%02d/%d", month, day, year), nil
	}
	return "", errors.New("unknown locale")
}

// FormatChange returns the correct change formatted according to given locale,
// currency, and actual change value (in cents)
func FormatChange(change int, locale string, currency string) (string, error) {
	cs := ""
	switch currency {
	case "EUR":
		cs = "€"
	case "USD":
		cs = "$"
	}
	var sSym, eSym, cSym, dSym string
	switch locale {
	case "nl-NL":
		sSym, eSym = cs+" ", " "
		if change < 0 {
			change = -change
			eSym = "-"
		}
		cSym, dSym = ".", ","
	case "en-US":
		sSym, eSym = " "+cs, " "
		if change < 0 {
			change = -change
			sSym, eSym = "("+cs, ")"
		}
		cSym, dSym = ",", "."
	default:
		return "", errors.New("unknown locale")
	}
	if change < 100000 {
		return fmt.Sprintf("%s%d%s%02d%s", sSym, change/100, dSym, change%100, eSym), nil
	}
	return fmt.Sprintf("%s%d%s%03d%s%02d%s", sSym, change/100000, cSym, (change/100)%1000, dSym, change%100, eSym), nil
}

// FormatDesc returns a truncated string if the description is longer than 25
// else it returns a space padded description
func FormatDesc(description string) string {
	if len(description) > 25 {
		return description[:22] + "..."
	}
	return fmt.Sprintf("%-25s", description)
}

// SortEntries returns sorted version of the inputted Entry slice
// Note: I retained the original sorting method used in the original code
func SortEntries(entries []Entry) []Entry {
	if len(entries) == 0 {
		return entries
	}
	m1 := map[bool]int{true: 0, false: 1}
	m2 := map[bool]int{true: -1, false: 1}
	es := append([]Entry{}, entries...)
	orig_es := es
	for len(es) > 1 {
		first, rest := es[0], es[1:]
		success := false
		for !success {
			success = true
			for i, e := range rest {
				if (m1[e.Date == first.Date]*m2[e.Date < first.Date]*4 +
					m1[e.Description == first.Description]*m2[e.Description < first.Description]*2 +
					m1[e.Change == first.Change]*m2[e.Change < first.Change]*1) < 0 {
					es[0], es[i+1] = es[i+1], es[0]
					success = false
				}
			}
		}
		es = es[1:]
	}
	return orig_es
}

// FormatLedger returns formatted ledger report given currency, locale, and Entry records
func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	if currency != "EUR" && currency != "USD" {
		return "", errors.New("invalid currency")
	}
	s, err := AddHeader(locale)
	if err != nil {
		return "", err
	}
	es := SortEntries(entries)
	for _, e := range es {
		dateStr, err := FormatDate(e.Date, locale)
		if err != nil {
			return "", err
		}
		changeStr, err := FormatChange(e.Change, locale, currency)
		if err != nil {
			return "", err
		}
		changeStr = fmt.Sprintf("%13s", changeStr)
		descStr := FormatDesc(e.Description)
		s += dateStr + " | " + descStr + " | " + changeStr + "\n"
	}
	return s, nil
}
