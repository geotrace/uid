package uid

import (
	"encoding/base64"
	"encoding/binary"
	"time"
)

// Info описывает разобранную информацию об уникальном идентификаторе.
type Info struct {
	Time    time.Time // время создания идентификатора
	Machine []byte    // три байта с идентификатором компьютера
	Pid     uint16    // идентификатор процесса
	Counter uint32    // счетчик
}

// Parse разбирает и возвращает информацию об уникальном идентификаторе. Если
// строка не является уникальным идентификатором, то возвращается nil.
func Parse(uid string) *Info {
	if len(uid) != 16 {
		return nil
	}
	data, err := base64.RawURLEncoding.DecodeString(uid)
	if err != nil {
		return nil
	}
	counter := data[9:12]
	return &Info{
		Time:    time.Unix(int64(binary.BigEndian.Uint32(data[0:4])), 0),
		Machine: data[4:7],
		Pid:     binary.BigEndian.Uint16(data[7:9]),
		Counter: uint32(uint32(counter[0])<<16 | uint32(counter[1])<<8 |
			uint32(counter[2])),
	}
}
