package imagick

type Destroyer interface {
	Destroy()
}

func Destroy(d Destroyer) {
	d.Destroy()
}
