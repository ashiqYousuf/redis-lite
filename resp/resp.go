package resp

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

const (
	STRING  = '+'
	ERROR   = '-'
	INTEGER = ':'
	BULK    = '$'
	ARRAY   = '*'
)

type Value struct {
	typ   string
	str   string
	num   int
	bulk  string
	array []Value
}

type Resp struct {
	reader *bufio.Reader
}

func NewResp(r io.Reader) *Resp {
	return &Resp{reader: bufio.NewReader(r)}
}

func (r *Resp) readLine() (line []byte, n int, err error) {
	for {
		b, err := r.reader.ReadByte()
		if err != nil {
			return nil, 0, err
		}

		n += 1
		line = append(line, b)

		if len(line) >= 2 && line[len(line)-2] == '\r' {
			// A  H \r \n
			// 0  1  2  3
			// 1  2  3  4
			break
		}
	}

	return line[:len(line)-2], n, nil
}

func (r *Resp) readInteger() (x int, n int, err error) {
	line, n, err := r.readLine()
	if err != nil {
		return 0, 0, err
	}

	i64, err := strconv.ParseInt(string(line), 10, 64)
	if err != nil {
		return 0, 0, err
	}

	return int(i64), n, nil
}

func (r *Resp) Read() (Value, error) {
	_type, err := r.reader.ReadByte()
	if err != nil {
		return Value{}, nil
	}

	switch _type {
	case ARRAY:
		return r.readArray()
	case BULK:
		return r.readBulk()
	default:
		fmt.Printf("unknown type: %v", string(_type))
		return Value{}, nil
	}
}

/*
RESP Array looks like: *2\r\n$5\r\nhello\r\n$5\r\nworld\r\n
OR
*2
$5
hello
$5
world

Reading Array:-
	i. Skip the first byte because we have already read it in
	the Read method.
	ii. Read the integer that represents the number of elements
	in the array.
	iii. Iterate over the array and for each line call Read
	method to parse the type according to symbol.
	iv. With each iteration, append the parsed value to the
	array in the Value object and return it
*/

func (r *Resp) readArray() (Value, error) {
	v := Value{}
	v.typ = "array"

	// read length of array
	length, _, err := r.readInteger()
	if err != nil {
		return v, err
	}

	v.array = make([]Value, length)
	for i := 0; i < length; i++ {
		val, err := r.Read()
		if err != nil {
			return v, err
		}

		v.array = append(v.array, val)
	}

	return v, nil
}

/*
RESP Bulk string: $5\r\nhello\r\n
OR
$5
hello

Reading Bulk strings:
	i. Skip the first byte because we have already read it in
	the Read method.
	ii. Read the integer that represents the number of bytes
	in the bulk string.
	iii. Read the bulk string, followed by the ‘\r\n’ that
	indicates the end of the bulk string
*/

func (r *Resp) readBulk() (Value, error) {
	v := Value{}
	v.typ = "bulk"

	length, _, err := r.readInteger()
	if err != nil {
		return v, err
	}

	bulk := make([]byte, length)
	r.reader.Read(bulk)
	v.bulk = string(bulk)

	// read trailing CRLF
	r.readLine()

	return v, nil
}
