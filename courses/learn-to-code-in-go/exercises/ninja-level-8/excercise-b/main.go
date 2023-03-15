package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	str := `[{"FirstName":"James","LastName":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"FirstName":"Miss","LastName":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"FirstName":"M","LastName":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	// fmt.Println(str)
	bs := []byte(str)
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print("All the data:\n", people)
	for i, person := range people {
		fmt.Println("\nperson", i+1, "-->", person.FirstName, person.LastName)
	}
}
