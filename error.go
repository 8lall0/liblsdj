package liblsdj

type Lsdj_error_t struct {
	message string
}

func Lsdj_error_new(t *Lsdj_error_t, message string) {
	if t == nil {
		return
	}
	t.message = message
}
