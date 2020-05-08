//Create a functional Card object, that represents a playing chard.  Must include appropriate accessors and constructor.  It needs the following functions: toString, isFaceCard, isRed, isBlack.  Create a main() method for testing that demonstrates each of the above functions you have written.
//Add the following functions: equals, hasSameRank, hasSameSuit, hasGreaterRank
package main
import "fmt"
type Card struct { //Card fields (added some as needed)
    toString string
    isFaceCard bool
    isRed bool
    isBlack bool
    rank int
    suit string
}
// make a new card
func newCard(name string,face bool,red bool,black bool,rank int,suit string) *Card {
  c:=Card{}
  c.toString = name
  c.isFaceCard = face
  c.isRed = red
  c.isBlack = black
  c.rank = rank
  c.suit = suit
   return &c
}
//the "equals" method, I don't know how to use them or if I wrote any method right
func (c Card) equals (card1 Card,card2 Card) bool{
  var boolean bool
  if card1.toString == card2.toString{
    boolean = true
  }else{
    boolean = false
  }
  return boolean
}
//the "hasSameRank" method, I don't know how to use them or if I wrote any method right
func (c Card) hasSameRank(card1 Card,card2 Card) bool{
  var boolean bool
  if card1.rank == card2.rank{
    boolean = true
  }else{
    boolean = false
  }
return boolean
}
//the "hasSameSuit" method, I don't know how to use them or if I wrote any method right
func (c Card) hasSameSuit(card1 Card, card2 Card) bool{
  var boolean bool
  if card1.suit == card2.suit{
    boolean = true
  }else{
    boolean = false
  }
return boolean
}
//the "hasGreaterRank" method, I don't know how to use them or if I wrote any method right
func (c Card) hasGreaterRank(card1 Card, card2 Card) bool{
  var boolean bool
  if card1.rank >= card2.rank{
    boolean = true
  }else {
    boolean = false
  }
return boolean
}
func main() {
  var placeHolder bool //tried to use the placeHolder and recreate the example in gobyexample with no success.
  Card1:=Card{"Ace of Hearts",false,true,false,1,"Ace"}
  Card2:Card{"Ace of Spades",false,false,true,1,"Ace"}
}
//I don't know what I did wrong but everytime I change something I get more errors, I will await feedback and attempt to fix this later.