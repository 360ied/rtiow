package material

type HittableList struct {
	Objects []Hittable
}

func (h *HittableList) Clear() {
	h.Objects = []Hittable{}
}

func (h *HittableList) Add(object Hittable) {
	h.Objects = append(h.Objects, object)
}

func (h HittableList) Hit(r Ray, tMin float64, tMax float64) (rec HitRecord, hitAnything bool) {
	closestSoFar := tMax
	for _, object := range h.Objects {
		if tempRec, hit := object.Hit(r, tMin, closestSoFar); hit {
			hitAnything = true
			closestSoFar = tempRec.T
			rec = tempRec
		}
	}
	return
}
