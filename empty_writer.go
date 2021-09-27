package logger

type emptyWriter struct{}

func (f emptyWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}
