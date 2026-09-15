package clock

import ("fmt"
       "strconv")

// Define the Clock type here.
type Clock struct{
    Hour int
    Minute int
}

func New(h, m int) Clock {
    minutoSobra := ((m % 60) + (60)) % 60
    h += (m - minutoSobra) / 60
	horaSobra := ((h % 24) + (24)) % 24
	return Clock{Hour: horaSobra, Minute: minutoSobra}
}

func (c Clock) Add(m int) Clock {
    return New(c.Hour, c.Minute + m)
}

func (c Clock) Subtract(m int) Clock {
    return New(c.Hour, c.Minute - m)
}

func (c Clock) String() string {
    hours := strconv.Itoa(c.Hour)
    minutes := strconv.Itoa(c.Minute)
    if c.Hour < 10{
        hours = "0" + hours
    }
    if c.Minute < 10{
        minutes = "0" + minutes
    }
	return fmt.Sprintf("%s:%s", hours, minutes)
}
