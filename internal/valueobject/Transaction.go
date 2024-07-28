package valueobject

import "time"

type Transaction struct {
	amount    int
	from      int
	to        int
	createdAt time.Time
}
