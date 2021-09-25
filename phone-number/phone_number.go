// Package phonenumber implements routines to process
// phone number strings
package phonenumber

import (
    "fmt"
    "unicode"
)

// Number takes in a string and checks if it is a valid 
// phone number or not. If it is return the cleaned up
// phone number.
func Number(phone string) (string, error) {
    phoneNumber := ""
    for _, c := range phone {
        if unicode.IsDigit(c) {
            phoneNumber += string(c)
        }
    }
    n := len(phoneNumber)
    if n == 10 {
        if phoneNumber[0] == '0' || phoneNumber[0] == '1' {
            return "", fmt.Errorf("invalid area code")
        }
        if phoneNumber[3] == '0' || phoneNumber[3] == '1' {
            return "", fmt.Errorf("invalid exchange code")
        }
        return phoneNumber, nil
    }
    if n == 11 && phoneNumber[0] == '1' {
        if phoneNumber[1] == '0' || phoneNumber[1] == '1' {
            return "", fmt.Errorf("invalid area code")
        }
        if phoneNumber[4] == '0' || phoneNumber[4] == '1' {
            return "", fmt.Errorf("invalid exchange code")
        }
        return phoneNumber[1:], nil
    }
    return "", fmt.Errorf("invalid phone number")
}

// AreaCode takes in a phone number and returns the area code
func AreaCode(phone string) (string, error) {
    phoneNumber, err := Number(phone)
    if err == nil {
        return phoneNumber[:3], nil
    }
    return "", err
}

// Format takes in a phone number and returns the phone number 
// in this format: (NXX) NXX-XXXX.
func Format(phone string) (string, error) {
    phoneNumber, err := Number(phone)
    if err == nil {
        return "("+phoneNumber[:3]+") "+phoneNumber[3:6]+"-"+phoneNumber[6:], nil
    }
    return "", err
}
