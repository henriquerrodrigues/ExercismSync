package collatzconjecture

import("errors")

func CollatzConjecture(n int) (int, error) {
    if n <= 0{
        return 0, errors.New("o número deve ser um inteiro maior que zero")
    }
	passos := 0
	var helper func(int)
    helper = func(n int){
        if n == 1{
            return 
        }
        
        passos ++
        
		if n % 2 == 0{
            helper(n/2)
        }else{
            helper((n*3)+1)
        }
    }

    helper(n)

    return passos, nil        
}
