package job

type Description string

func (d *Description) String() string {
	return string(*d)
}
