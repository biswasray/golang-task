package main

import "fmt"

func main() {
	// Normal String (Can not contain newlines, and can have escape characters like `\n`, `\t` etc)
	var github = "\thttps://github.com/biswasray\t\n"

	// Raw String (Can span multiple lines. Escape characters are not interpreted)
	var myDescription = `\t\tHey Guys... wave\n
						Welcome to my programming world\n\t\tI am Subhasish Biswasray
                        - A passionate coder
						- Being myself
						- Currently working on several languages
						- Don't have much time
						- Expanding my limits`

	fmt.Println(github, myDescription)
}
