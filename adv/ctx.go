package adv

import (
	"context"
	"fmt"
	"time"
)

func watch(ctx context.Context, key string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("End of watching", ctx.Value(key))
			return
		default:
			fmt.Println("Watching", ctx.Value(key))
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func CtxMain() {
	ctx, cancel := context.WithCancel(context.Background())
	valueCtx1 := context.WithValue(ctx, "name", "JIANGYAN")
	valueCtx2 := context.WithValue(ctx, "name", "PekingUniversity")
	go watch(valueCtx1, "name")
	go watch(valueCtx2, "name")

	time.Sleep(1 * time.Second)

	fmt.Println("OK, stop watching now!")
	cancel()
	time.Sleep(1 * time.Second)
}
