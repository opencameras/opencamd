package capture

// Format is capture format for one capture device
type Format struct {
	Codec int
	Width uint
	Height uint
	MinFrameRate float32
	MaxFrameRate float32
}

func (f Format) CodecName() string {
	name := make([]byte, 4)
	name[0] = byte(f.Codec >> 24 & 0xff)
	name[1] = byte(f.Codec >> 16 & 0xff)
	name[2] = byte(f.Codec >> 8 & 0xff)
	name[3] = byte(f.Codec & 0xff)
	return string(name)
}

// Device is capture device information
type Device struct {
	Name string
	Model string
	Manufacturer string
	Formats []Format
}
