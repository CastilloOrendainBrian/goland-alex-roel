import (
	"fmt"

	"github.com/CastilloOrendainBrian/goland-alex-roel"
)

func main() {
	messages, err := greetings.Hellos("Juan")

	if err != nil {
		fmt.Println("Ocurrió un error:", err)
		return
	}

	fmt.Println(messages)
}