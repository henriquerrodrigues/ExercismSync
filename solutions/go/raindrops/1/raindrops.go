package raindrops

import("strconv")

func Convert(number int) string {
    var result string
	if number % 3 == 0 || number % 5 == 0 || number % 7 == 0{
        if (number % 3 ) == 0{
            result = result + "Pling"
        }
        if (number % 5 ) == 0{
            result = result + "Plang"
        }
        if (number % 7 ) == 0{
            result = result + "Plong"
        }
        return result
    }
    return strconv.Itoa(number)
}
