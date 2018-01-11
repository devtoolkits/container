package safequeue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewQueue(t *testing.T) {
	q := NewQueue()
	assert.NotNil(t, q)
	assert.Equal(t, 0, q.Size())
	assert.False(t, q.Full())
}

func Test_NewCappedQueue(t *testing.T) {
	cq := NewCappedQueue(1)
	assert.False(t, cq.Full())
	assert.True(t, cq.Push(1))
	assert.True(t, cq.Full())
	assert.False(t, cq.Push(2))
}

func Test_PushPop(t *testing.T) {
	q := NewQueue()
	q.Push("1")
	assert.Equal(t, "1", q.Pop())
	assert.Equal(t, nil, q.Pop())
}

func Test_PopN(t *testing.T) {
	q := NewQueue()
	q.Push("1")
	q.Push("2")
	q.Push("3")
	q.Push(nil)
	q.Push("4")

	assert.Equal(t, []interface{}{"1", "2", "3", nil, "4"}, q.PopN(6))
}

func Benchmark_Push(b *testing.B) {
	q := NewQueue()
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
}

func Benchmark_Pop(b *testing.B) {
	q := NewQueue()
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

func Benchmark_PopN(b *testing.B) {
	q := NewQueue()
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
	b.StartTimer()

	for i := 0; i < b.N/2; i++ {
		q.PopN(2)
	}
}
