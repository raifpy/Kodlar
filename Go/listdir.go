package main
import(
	"os"
	"fmt"
)

func main(){
	abc , err:= os.readDir()
	fmt.Println(abc,err)
}