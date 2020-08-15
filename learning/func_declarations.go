package learning

import (
	"fmt"
)

func CallFunc() {

	//The underscore is used to indicate that I won't use the value
	//It has to be declared and can't have a name if it will not be used
	_, err := StartWebServer(3000, 1)

	if err != nil {
		fmt.Println("Error occured")
	} else {
		fmt.Println("Success")
	}
}

//it's possible to declare a bunch of variables separated by comma before the type
//port and retries are int type
func StartWebServer(port, retries int) (bool, error) {
	fmt.Println("Starting server...")
	fmt.Println("Server started on port", port)

	return true, nil
	//an example with an error
	//return true, errors.New("Something went wrong")
}
