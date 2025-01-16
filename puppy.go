package puppy
import (
"github.com/nazrawigedion123/dog"
)

func Bark() string {
	return "woof"
}

func Barks() string {
	return "woof! woof! woof!"
}

func BigBark() string{
        return dog.WhenGrownUp(Bark())
}
func BigBarks() string{
	return dog.WhenGrownUp(Barks())
}
func Form11(){
	return(Println("Im'from version 1.1.0")
}
func Form12(){
	return(Println("I'm from version1.2.0"))
}
