package command

type HelperHandler struct {
	Disc  string
	Usage string
}

type Command[K comparable, C, I, O any] struct { // nolint: golint
	Key K
	Fn  func(*I) (*O, error)

	CustomField C
	Helper      HelperHandler
}

func CreateCommand[K comparable, C, I, O any](key K, fn func(*I) (*O, error), customField C, disc string, usage string) Command[K, C, I, O] {
	return Command[K, C, I, O]{Key: key, Fn: fn, CustomField: customField, Helper: HelperHandler{Disc: disc, Usage: usage}}
}
