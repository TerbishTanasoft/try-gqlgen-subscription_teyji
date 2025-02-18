package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
	"time"
	"try-gqlgen-subscription_teyji/graph/model"
)

// Placeholder is the resolver for the placeholder field.
func (r *queryResolver) Placeholder(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented: Placeholder - placeholder"))
}

// CurrentTime is the resolver for the currentTime field.
func (r *subscriptionResolver) CurrentTime(ctx context.Context) (<-chan *model.Time, error) {
	// First you'll need to `make()` your channel. Use your type here!
	ch := make(chan *model.Time)

	// You can (and probably should) handle your channels in a central place outside of `schema.resolvers.go`.
	// For this example we'll simply use a Goroutine with a simple loop.
	go func() {
		// Handle deregistration of the channel here. Note the `defer`
		defer close(ch)
		flag := true
		for {
			// In our example we'll send the current time every second.
			time.Sleep(1 * time.Second)
			if flag {
				fmt.Println("Tick")
			} else {
				fmt.Println("Tock")
			}
			flag = !flag
			// Prepare your object.
			currentTime := time.Now()
			t := &model.Time{
				UnixTime:  int(currentTime.Unix()),
				TimeStamp: currentTime.Format(time.RFC3339),
			}

			// The subscription may have got closed due to the client disconnecting.
			// Hence we do send in a select block with a check for context cancellation.
			// This avoids goroutine getting blocked forever or panicking,
			select {
			case <-ctx.Done(): // This runs when context gets cancelled. Subscription closes.
				fmt.Println("Subscription Closed")
				// Handle deregistration of the channel here. `close(ch)`
				return // Remember to return to end the routine.

			case ch <- t: // This is the actual send.
				// Our message went through, do nothing
			}
		}
	}()

	// We return the channel and no error.
	return ch, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
