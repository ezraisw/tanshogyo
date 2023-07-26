package repository

type FindOptions struct {
	Relations []string
}

type FindManyOptions struct {
	Orderings []Ordering
	Limit     int
	Offset    int
	Relations []string
}

type Ordering struct {
	By   string
	Desc bool
}
