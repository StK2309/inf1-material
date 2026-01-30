package duration

import "fmt"

type duration int

func FromSeconds(t int) duration {
	return duration(t)
}

func Minute(t int) duration {
	return duration(t * 60)
}

func Hours(t int) duration {
	return duration(t * 3600)
}

func Day(t int) duration {
	return duration(t * 86400)
}

func Year(t int) duration {
	return duration(t * 31536000)
}

func (d duration) Scale(factor int) duration {
	return d * duration(factor)
}

// Seconds returns the duration in seconds as int.
func (d duration) Seconds() int {
	return int(d)
}
func (d duration) Minute() int {
	return int(d)
}
func (d duration) Hours() int {
	return int(d)
}
func (d duration) Day() int {
	return int(d)
}
func (d duration) Year() int {
	return int(d)
}
func Example() {
	a := FromSeconds(300)

	fmt.Println(a.Seconds())
	fmt.Println(a.Minute())

	// Output:

}
