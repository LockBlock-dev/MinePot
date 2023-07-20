package typings

import "io"

const SEGMENT_BITS = 0x7F
const CONTINUE_BIT = 0x80

type CustomReader struct {
    reader io.Reader
}

type CustomWriter struct {
    writer io.Writer
}

func NewCustomReader(reader io.Reader) *CustomReader {
    return &CustomReader{
        reader: reader,
    }
}

func (cr *CustomReader) ReadVarInt() (int, error) {
    value := 0
    position := 0
    currentByte := make([]byte, 1)

    for {
        _, err := cr.reader.Read(currentByte)
		if err != nil {
			return -1, err
		}

        value |= (int(currentByte[0]) & SEGMENT_BITS) << position;

        if ((int(currentByte[0]) & CONTINUE_BIT) == 0) {
			break
		}

        position += 7

        if (position >= 32) {
			return -1, err
		}
    }

    return value, nil
}

func NewCustomWriter(writer io.Writer) *CustomWriter {
    return &CustomWriter{
        writer: writer,
    }
}

func (cw *CustomWriter) WriteVarInt(value int) error {
    for {
		buf := make([]byte, 1)

        if (value & ^SEGMENT_BITS) == 0 {
			buf[0] = byte(value)
            _, err := cw.writer.Write(buf)
            return err
        }

		buf[0] = byte((value & SEGMENT_BITS) | CONTINUE_BIT)
       	if _, err := cw.writer.Write(buf); err != nil {
			return err
	   	}
    
		// equivalent of: value >>>= 7
		uvalue := uint(value)
		value = int(uvalue >> 7)
    }
}
