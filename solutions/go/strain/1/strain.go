package strain

func Keep[T any](collection []T, predicate func(T) bool) []T{
    var result []T

    for _, element := range collection{
        if predicate(element){
            result = append(result, element)
        }
    }
    return result
}

func Discard[T any](collection []T, predicate func(T) bool) []T{
    var result []T
    
    for _, element := range collection {
        if !predicate(element){
            result = append(result, element)
        }
    }
    return result
}
// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1
