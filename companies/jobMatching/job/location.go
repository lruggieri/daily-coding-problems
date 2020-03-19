package job

// for simplicity
type Location string
func (l Location) String() string{
	return string(l)
}

func (l *Location) IsEqual(otherLocation *Location) bool{
	return l.String() == otherLocation.String()
}