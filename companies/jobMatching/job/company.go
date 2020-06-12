package job

type Company string

func (c Company) String() string {
	return string(c)
}

func (c *Company) IsEqual(otherCompany *Company) bool {
	return c.String() == otherCompany.String()
}
