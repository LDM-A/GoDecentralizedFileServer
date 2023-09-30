package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newStore() *Store {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	return NewStore(opts)
}

func teardown(t *testing.T, s *Store) {
	err := s.Clear()
	assert.Nil(t, err)
}

func TestPathTransformFunc(t *testing.T) {
	key := "momsbestpicture"
	pathName := CASPathTransformFunc(key)
	expectedOriginalKey := "6804429f74181a63c50c3d81d733a12f14a353ff"
	testPathName := "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff"
	assert.Equal(t, testPathName, pathName.Pathname)
	assert.Equal(t, expectedOriginalKey, pathName.Filename)
}

func TestStore(t *testing.T) {
	s := newStore()
	defer teardown(t, s)
	count := 10
	for i := 0; i < count; i++ {
		key := fmt.Sprintf("foo%d", i)

		data := []byte("Some jpg bytes")
		err := s.writeStream(key, bytes.NewReader(data))
		assert.Nil(t, err)

		ok := s.Has(key)
		assert.True(t, ok)

		r, err := s.Read(key)
		assert.Nil(t, err)

		b, err := ioutil.ReadAll(r)
		assert.Nil(t, err)

		assert.Equal(t, data, b)

		err = s.Delete(key)
		assert.Nil(t, err)

	}

}
