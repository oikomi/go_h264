package h264

import (
	"bufio"
	"github.com/oikomi/go_h264/nalu"
	"log"
	"os"
)

type H264FileHandler struct {
	file *os.File
	r    *bufio.Reader
}

func New() *H264FileHandler {
	fh := &H264FileHandler{}
	return fh
}

func FindStartCode3(b []byte) int {
	if b[0] != 0 || b[1] != 0 || b[2] != 1 {
		return 0
	} else {
		return 1
	}
}

func FindStartCode4(b []byte) {
	if b[0] != 0 || b[1] != 0 || b[2] != 0 || b[3] != 1 {
		return 0 //0x00000001?
	} else {
		return 1
	}
}

func (self *H264FileHandler) Open(filePath string) (err error) {
	self.file, err = os.Open(filePath)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	self.r = bufio.NewReader(self.file)
	return nil
}

func (self *H264FileHandler) Read(size int64) ([]byte, error) {
	buf := make([]byte, size)
	_, err := self.file.Read(buf)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}
	return buf, nil
}

func (self *H264FileHandler) Seek(offset int64, whence int) error {
	_, err := self.file.Seek(offset, whence)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	return nil
}

func (self *H264FileHandler) GetAnnexbNALU(nalu *nalu.Nalu) (err error) {
	var (
		b                                 []byte
		pos, StartCodeFound, info3, info4 int64
	)
	if b, err = self.Read(3); err != nil {
		log.Fatalln(err.Error())
		return
	}
	if FindStartCode3(b) != 1 {
		if b, err = self.Read(4); err != nil {
			log.Fatalln(err.Error())
			return
		}
		if FindStartCode4(b) != 1 {
			return
		} else {
			pos = 4
		}
	} else {
		pos = 3
	}

	

	return
}
