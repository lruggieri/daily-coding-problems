package job

type Title string
func (l Title) String() string{
	return string(l)
}

func (l *Title) IsEqual(otherTitle *Title) bool{
	return l.String() == otherTitle.String()
}