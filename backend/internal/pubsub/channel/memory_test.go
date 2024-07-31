package pubsubchannel

import (
	"context"
	"sync"
	"testing"
)

/*
BenchmarkHighLoadPublish-16    	 1331211	       928.6 ns/op	       0 B/op	       0 allocs/op
*/
func BenchmarkHighLoadPublish(b *testing.B) {
	ctx := context.Background()
	psub := NewMemoryPubsub()
	go psub.Listen(ctx)

	numSubscriptions := 1
	numPublishs := b.N // Use b.N to automatically adjust the number of iterations based on time

	wg := sync.WaitGroup{}
	wg.Add(numPublishs * numSubscriptions)

	for range numSubscriptions {
		r, err := psub.Sub(ctx, "testTopic")
		if err != nil {
			b.Fatal(err)
		}

		go func() {
			for msg := range r.Data() {
				if msg != "Test" {
					b.Error("Received unexpected message:", msg)
				}
				wg.Done()
			}
		}()
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	defer b.StopTimer()
	for range numPublishs {
		psub.Pub(ctx, "testTopic", "Test")
	}
	wg.Wait()
}

/*
cpu: AMD Ryzen 7 6800U with Radeon Graphics
BenchmarkHighLoadSubsciption-16    	  288598	      4468 ns/op	     890 B/op	       9 allocs/op
*/
func BenchmarkHighLoadSubsciption(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	ctx := context.Background()
	psub := NewMemoryPubsub()
	go psub.Listen(ctx)

	numSubscriptions := b.N
	numPublishs := 1

	wg := sync.WaitGroup{}
	wg.Add(numPublishs * numSubscriptions)

	for range numSubscriptions {
		r, err := psub.Sub(ctx, "testTopic")
		if err != nil {
			b.Fatal(err)
		}

		go func() {
			b.StartTimer()
			defer b.StopTimer()
			for msg := range r.Data() {
				if msg != "Test" {
					b.Error("Received unexpected message:", msg)
				}
				wg.Done()
			}
		}()
	}

	for range numPublishs {
		psub.Pub(ctx, "testTopic", "Test")
	}
	wg.Wait()
}
