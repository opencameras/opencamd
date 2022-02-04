package media

import (
	"fmt"
	"github.com/pion/webrtc/v3/pkg/media/h264reader"
	"github.com/pion/webrtc/v3/pkg/media/oggreader"
	"os"
	"path"
)

type Reader interface {
}

type Datasource interface {
	GetCodecs() ([]string, error)
	OpenH264() (*h264reader.H264Reader, error)
	OpenOgg() (*oggreader.OggReader, error)
	Close() error
	String() string
}

type LocalFile struct {
	Filename string
	Codec string
	f *os.File
}

func (lf *LocalFile) GetCodecs() ([]string, error) {
	e := path.Ext(lf.Filename)
	if e != ".h264" && e != ".ogg" {
		return nil, nil
	}
	_, err := os.Stat(lf.Filename)
	if err != nil {
		fmt.Println(err)
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	if lf.Codec != "" {
		return []string{lf.Codec}, nil
	}
	return []string{e}, nil
}

func (lf *LocalFile) OpenH264() (reader *h264reader.H264Reader, err error) {
	lf.f, err = os.Open(lf.Filename)
	if err != nil {
		return
	}

	return h264reader.NewReader(lf.f)
}

func (lf *LocalFile) OpenOgg() (reader *oggreader.OggReader, err error){
	lf.f, err = os.Open(lf.Filename)
	if err != nil {
		return
	}

	// Open on oggfile in non-checksum mode.
	reader, _, err = oggreader.NewWith(lf.f)
	return
}

func (lf *LocalFile) Close() error {
	return lf.f.Close()
}

func (lf *LocalFile) String() string {
	return lf.Filename
}
