package raindrops

import "strconv"

func Convert(number int) string {
	result := ""

    factors := []struct {
        value int
        sound string
    } {
        {3, "Pling"},
        {5, "Plang"},
        {7, "Plong"},
    }

    for _, factor := range factors {
        if number%factor.value == 0 {
            result += factor.sound
        }
    }
    if result == "" {
        result = strconv.Itoa(number)
    }
    return result
}