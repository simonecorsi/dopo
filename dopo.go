package dopo

type dopo struct {
	await func() interface{}
}

func Exec(f func() interface{}) dopo {
	c := make(chan interface{})
	go func() {
		c <- f()
	}()
	return dopo{
		await: func() interface{} {
			return <-c
		},
	}
}
