package erro
import "fmt"

func ErrHandle(err error){
	if err!=nil{
		fmt.Println("Error Encountered.")
	}
}