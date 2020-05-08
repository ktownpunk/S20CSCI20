package main

import "fmt"

type eggCarton struct {
    whiteEggs, brownEggs int
}

func (e *eggCarton) addEggs(){
}

func (e *eggCarton) setWhiteEggs(whiteEggs){
  whiteEggs = 0
}
func (e *eggCarton) setBrownEggs(brownEggs){
  brownEggs = 0
}
func (e *eggCarton)  addEggs(eggs){
  whiteEggs = whiteEggs + 1
  brownEggs = brownEggs + 1
}
func (e *eggCarton) getBrownEggs() int{
return brownEggs
}
func (e *eggCarton) getWhiteEggs() int{
return whiteEggs
}
func (e *eggCarton) getTotalEggs() int{
return brownEggs+whiteEggs
}
func main() {
  fmt.Println(e.getTotalEggs)
}