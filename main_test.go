package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	a := assert.New(t)
	main()
	a.Equal(1, 1)
}
