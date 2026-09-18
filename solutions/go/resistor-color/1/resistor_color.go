package resistorcolor

// Colors returns the list of all colors.
func Colors() []string {
    return []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
    code := make(map[string]int)
    for ix, val := range Colors(){
        code[val] = ix
    }
    return code[color]
}
