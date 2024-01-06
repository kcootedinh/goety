package notify

import (
	"context"
	"testing"

	"github.com/code-gorilla-au/odize"
)

type testContextKey struct {
	TestKey string
}

func TestFromContextByKey(t *testing.T) {
	group := odize.NewGroup(t, nil)
	err := group.
		Test("should return a string value", func(t *testing.T) {

			ctx := AttachToContext[testContextKey, string](context.Background(), testContextKey{}, "value")

			result, err := FromContext[testContextKey, string](testContextKey{}, ctx)
			odize.AssertNoError(t, err)

			odize.AssertEqual(t, "value", result)
		}).
		Test("should return a channel value", func(t *testing.T) {
			testCh := make(chan int, 1)

			defer close(testCh)

			ctx := AttachToContext[testContextKey, chan int](context.Background(), testContextKey{}, testCh)

			notifyChannel, err := FromContext[testContextKey, chan int](testContextKey{}, ctx)
			odize.AssertNoError(t, err)

			notifyChannel <- 1
			odize.AssertEqual(t, 1, <-testCh)
		}).
		Run()

	odize.AssertNoError(t, err)
}
