<<<<<<< HEAD
=======
// Генератор глобальных уникальных идентификаторов.
//
// Алгоритм генерации глобальных уникальных идентификаторов основан на том же самом принципе, что
// используется для генерации уникальных идентификаторов в MongoDB. Уникальный идентификатор
// представляет из себя 12 байтовую последовательность, состоящую из времени генерации,
// идентификатора компьютера и процесса, а так же счетчика.
//
// Основное отличие состоит в том, что данный идентификатор сразу представлен в виде строки с
// использованием base64-encoding. Так же поддерживается функция для быстрого разбора такой строки,
// которая возвращает информацию о всех значениях, из которых собран данный идентификатор.
>>>>>>> origin/master
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

<<<<<<< HEAD
// objectIDCounter является счетчиком, который автоматически увеличивается после
// каждой генерации нового уникального ключа. Начальное значение данного ключа
// устанавливается случайным.
=======
// objectIDCounter является счетчиком, который автоматически увеличивается после каждой генерации
// нового уникального ключа. Начальное значение данного ключа устанавливается случайным.
>>>>>>> origin/master
var objectIDCounter = randInt()

// randInt возвращает случайное uint32 число.
func randInt() uint32 {
	b := make([]byte, 3)
	if _, err := rand.Reader.Read(b); err != nil {
		panic(fmt.Errorf("Cannot generate random number: %v;", err))
	}
	return uint32(b[0])<<16 | uint32(b[1])<<8 | uint32(b[2])
}

<<<<<<< HEAD
// machineId хранит идентификатор машины. Используется при генерации случайного
// идентификатора.
=======
// machineId хранит идентификатор машины. Используется при генерации случайного идентификатора.
>>>>>>> origin/master
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
<<<<<<< HEAD
			panic(
				fmt.Errorf("Cannot get hostname nor generate a random number: %v; %v",
					err, randErr))
=======
			panic(fmt.Errorf("Cannot get hostname nor generate a random number: %v; %v", err, randErr))
>>>>>>> origin/master
		}
	}
	return id
}

// New генерирует строку с глобальным уникальным идентификатором.
//
<<<<<<< HEAD
// Для генерации уникального идентификатора используется тот же алгоритм, что и
// у MongoDB. Отличие состоит только в формате представления сгенерированного
// идентификатора в строковом виде: для этого данная библиотека использует
// base64, вместо hex представления.
=======
// Для генерации уникального идентификатора используется тот же алгоритм, что и у MongoDB. Отличие
// состоит только в формате представления сгенерированного идентификатора в строковом виде:
// для этого данная библиотека использует base64, вместо hex представления.
>>>>>>> origin/master
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
