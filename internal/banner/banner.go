package banner

import "fmt"

const Logo = `
   ___       _    _ _          ___      _ _   
  / __|___  | |__(_) |_ ___   |_ _|_ _ (_) |_ 
 | (_ / _ \ | / /| |  _|___|   | || ' \| |  _|
  \___\___/ |_\_\|_|\__|     |___|_||_|_|\__|
                                              
  Go Project Generator - Fast & Clean
`

func Print() {
	fmt.Println(Logo)
}
