package vec3util

import (
	"fmt"

	"rtiow/vec3"
)

func Print(v vec3.Vec3) {
	fmt.Printf("%v %v %v", v.X, v.Y, v.Z)
}
