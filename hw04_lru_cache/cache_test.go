package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(3)

		wasInCache := c.Set("farnsworth", 160)
		require.False(t, wasInCache)

		wasInCache = c.Set("fry", 1033)
		require.False(t, wasInCache)

		wasInCache = c.Set("bender", 4)
		require.False(t, wasInCache)

		c.Clear()
		require.Equal(t, 0, c.(*lruCache).capacity)
		require.Equal(t, &list{}, c.(*lruCache).queue)
		require.Equal(t, 0, len(c.(*lruCache).items))
	})

	t.Run("pushing out of the cache", func(t *testing.T) {
		c := NewCache(3)

		wasInCache := c.Set("farnsworth", 160)
		require.False(t, wasInCache)

		wasInCache = c.Set("fry", 1033)
		require.False(t, wasInCache)

		wasInCache = c.Set("bender", 4)
		require.False(t, wasInCache)

		wasInCache = c.Set("leela", 41)
		require.False(t, wasInCache)

		require.Equal(t, 3, len(c.(*lruCache).items))
	})

	t.Run("pushing out long-used elements", func(t *testing.T) {
		c := NewCache(3)

		wasInCache := c.Set("farnsworth", 160) // [ farnsworth ]
		require.False(t, wasInCache)

		wasInCache = c.Set("fry", 1033) // [ fry, farnsworth ]
		require.False(t, wasInCache)

		wasInCache = c.Set("bender", 4) // [ bender, fry, farnsworth ]
		require.False(t, wasInCache)

		_, wasInCache = c.Get("fry") // [ fry, bender, farnsworth ]
		require.True(t, wasInCache)

		wasInCache = c.Set("bender", 1061) // [ bender, fry, farnsworth ]
		require.True(t, wasInCache)

		wasInCache = c.Set("leela", 41) // [ leela, bender, fry ]
		require.False(t, wasInCache)

		require.Equal(t, &Pair{key: "leela", value: 41}, c.(*lruCache).queue.Front().Value)
		require.Equal(t, &Pair{key: "bender", value: 1061}, c.(*lruCache).queue.Front().Next.Value)
		require.Equal(t, &Pair{key: "fry", value: 1033}, c.(*lruCache).queue.Back().Value)
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
