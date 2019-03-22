package main

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestUlid(t *testing.T) {
	var err error
	hash := fnv.New64()

	// hostname + pidをシードにする
	b := new(bytes.Buffer)
	_, err = fmt.Fprintf(b, "%s:%d", os.Getenv("HOSTNAME"), os.Getpid())
	checkError(t, err)
	t.Logf("HashSource: %s", b.String())
	_, err = hash.Write(b.Bytes())
	checkError(t, err)
	// 先の値の64bithashからentropy sourceを作成
	entropy := ulid.Monotonic(rand.New(rand.NewSource(int64(hash.Sum64()))), 0)

	// ID生成
	tx := time.Unix(1000000, 0)
	id := ulid.MustNew(ulid.Timestamp(tx), entropy)

	// text表現
	t.Logf("%s", id.String())

	// text 復元
	idParsed, err := ulid.ParseStrict(id.String())
	checkError(t, err)
	assert.Equal(t, id, idParsed)

	// bynary表現
	idb, err := id.MarshalBinary()
	checkError(t, err)
	t.Logf("%v", idb)

}

func checkError(t *testing.T, err error) {
	if err != nil {
		t.Logf("%+v", err)
		t.FailNow()
	}
}
