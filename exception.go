package exception

func Throw(err error) {
	if err != nil {
		panic(err)
	}
}
