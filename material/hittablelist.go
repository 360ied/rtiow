package material

type HittableList struct {
	Objects []Hittable
	box     *Box
}

func (h *HittableList) Clear() {
	h.Objects = []Hittable{}
	h.box = nil
}

func (h *HittableList) Add(object Hittable) {
	h.Objects = append(h.Objects, object)
	if h.box == nil {
		b := object.Box()
		h.box = &b
	} else {
		*h.box = h.box.Surrounding(object.Box())
	}
}

func (h HittableList) Box() Box {
	if h.box == nil {
		return Box{}
	} else {
		return *h.box
	}
}

func (h HittableList) Hit(r Ray, tMin float64, tMax float64) (rec HitRecord, hitAnything bool) {
	closestSoFar := tMax
	for _, object := range h.Objects {
		if !object.Box().Hit(r, tMin, closestSoFar) {
			continue
		}
		if tempRec, hit := object.Hit(r, tMin, closestSoFar); hit {
			hitAnything = true
			closestSoFar = tempRec.T
			rec = tempRec
		}
	}
	return
}
