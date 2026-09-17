package luhn

import("strings"
       "strconv")

func Valid(id string) bool {
	sequence := strings.ReplaceAll(id," ", "")
    if len(sequence) <= 1{
        return false
    }

    digits := make([]int, len(sequence))
    for i:= 0 ; i < len(sequence); i++{
        n, err := strconv.Atoi(string(sequence[i]))
        if err != nil{
            return false
        }
        digits[i] = n
    } 
    
    for i:= len(digits)-2; i >= 0; i-= 2{
        digits[i] *= 2 
        if digits[i] > 9{
           digits[i] -= 9
        }
    }
    total := 0
    for _,d := range digits{
        total += d
    }

    return  total % 10 == 0
}
