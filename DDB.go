package dxb

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

func ReadDDB(f *os.File) ([]float64, error) {
	const b = 8

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read .DSB file at ioutil.ReadAll: %w", err)
	}
	length := len(buf) / b
	samples := make([]float64, length, length)
	var f64 float64
	for i := 0; i < length; i++ {
		bufReader := bytes.NewReader(buf[b*i : b*(i+1)])
		err := binary.Read(bufReader, binary.LittleEndian, &f64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse binary at binary.Read: %w", err)
		}
		samples[i] = f64
	}
	return samples, nil
}

func WriteDDA(f *os.File, samples []float64) error {
	for _, sample := range samples {
		_, err := f.Write(float64ToByte(sample))
		if err != nil {
			return err
		}
	}
	return nil
}

func float64ToByte(f float64) []byte {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}
