package printer

type Printer interface {
	Write(buf []byte) error
	Close() error
}

