package uid

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

// objectIDCounter является счетчиком, который автоматически увеличивается после каждой генерации
// нового уникального ключа. Начальное значение данного ключа устанавливается случайным.
var objectIDCounter = randInt()

// randInt возвращает случайное uint32 число.
func randInt() uint32 {
	b := make([]byte, 3)
	if _, err := rand.Reader.Read(b); err != nil {
		panic(fmt.Errorf("Cannot generate random number: %v;", err))
	}
	return uint32(b[0])<<16 | uint32(b[1])<<8 | uint32(b[2])
}

// machineId хранит идентификатор машины. Используется при генерации случайного идентификатора.
var machineID = readMachineID()

// readMachineId инициализирует значение идентификатора компьютера.
func readMachineID() []byte {
	id := make([]byte, 3)
	if hostname, err := os.Hostname(); err == nil {
		hw := md5.New()
		hw.Write([]byte(hostname))
		copy(id, hw.Sum(nil))
	} else {
		// Fallback to rand number if machine id can't be gathered
		if _, randErr := rand.Reader.Read(id); randErr != nil {
			panic(fmt.Errorf("Cannot get hostname nor generate a random number: %v; %v", err, randErr))
		}
	}
	return id
}

// New генерирует строку с глобальным уникальным идентификатором.
//
// Для генерации уникального идентификатора используется тот же алгоритм, что и у MongoDB. Отличие
// состоит только в формате представления сгенерированного идентификатора в строковом виде:
// для этого данная библиотека использует base64, вместо hex представления.
func New() string {
	id := make([]byte, 12)
	// Timestamp, 4 bytes, big endian
	binary.BigEndian.PutUint32(id, uint32(time.Now().Unix()))
	// Machine, first 3 bytes of md5(hostname)
	id[4] = machineID[0]
	id[5] = machineID[1]
	id[6] = machineID[2]
	// Pid, 2 bytes, specs don't specify endianness, but we use big endian.
	pid := os.Getpid()
	id[7] = byte(pid >> 8)
	id[8] = byte(pid)
	// Increment, 3 bytes, big endian
	i := atomic.AddUint32(&objectIDCounter, 1)
	id[9] = byte(i >> 16)
	id[10] = byte(i >> 8)
	id[11] = byte(i)
	return base64.RawURLEncoding.EncodeToString(id)
}
