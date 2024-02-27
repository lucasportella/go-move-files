package types

type Date interface {
	GetYear()
	GetMonth()
	GetDay()
}

func (d DateValues) GetYear() string {
	return d.Year
}
func (d DateValues) GetMonth() string {
	return d.Month
}

func (d DateValues) GetDay() string {
	return d.Day
}

type DateValues struct {
	Year  string
	Month string
	Day   string
}

type DateYear struct {
	Year string
}

type DateMonth struct {
	Year  string
	Month string
}

type DateDay struct {
	Year  string
	Month string
	Day   string
}

type DateOption string

const (
	Yearly  string = "yearly"
	Monthly string = "monthly"
	Daily   string = "daily"
)
