package colour

import (
	"fmt"

	"rtiow/constants"
	"rtiow/vec3"
)

func WriteColour(v vec3.Colour) {
	fmt.Printf("%v %v %v\n", int(constants.Almost256*v.X), int(constants.Almost256*v.Y), int(constants.Almost256*v.Z))
}
