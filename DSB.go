package dxb

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func ReadDSB(f *os.File) ([]int16, error) {
	const b = 2
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read .DSB file at ioutil.ReadAll: %w", err)
	}
	length := len(buf) / b
	samples := make([]int16, length, length)
	var i16 int16
	for i := 0; i < length; i++ {
		bufReader := bytes.NewReader(buf[b*i : b*(i+1)])
		err := binary.Read(bufReader, binary.LittleEndian, &i16)
		if err != nil {
			return nil, fmt.Errorf("failed to parse binary at binary.Read: %w", err)
		}
		samples[i] = i16
	}
	return samples, nil
}

func WriteDSA(f *os.File, samples []int16) error {
	for _, sample := range samples {
		_, err := f.WriteString(strconv.Itoa(int(sample)) + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
