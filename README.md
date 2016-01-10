<<<<<<< HEAD
# uid [![Build Status](https://travis-ci.org/mdigger/uid.svg)](https://travis-ci.org/mdigger/uid)

Генератор глобальных уникальных идентификаторов
=======
# Генератор глобальных уникальных идентификаторов

[![GoDoc](https://godoc.org/github.com/geotrace/uid?status.svg)](https://godoc.org/github.com/geotrace/uid)
[![Build Status](https://travis-ci.org/geotrace/uid.svg)](https://travis-ci.org/geotrace/uid)
[![Coverage Status](https://coveralls.io/repos/geotrace/uid/badge.svg?branch=master&service=github)](https://coveralls.io/github/geotrace/uid?branch=master)

Алгоритм генерации глобальных уникальных идентификаторов основан на том же самом принципе, что используется для генерации уникальных идентификаторов в MongoDB. Уникальный идентификатор представляет из себя 12 байтовую последовательность, состоящую из времени генерации, идентификатора компьютера и процесса, а так же счетчика. 

Основное отличие состоит в том, что данный идентификатор сразу представлен в виде строки с использованием base64-encoding. Так же поддерживается функция для быстрого разбора такой строки, которая возвращает информацию о всех значениях, из которых собран данный идентификатор.
>>>>>>> origin/master
