package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func printGreeting(ctx context.Context) error {
	greeting, err := genGreeting(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", greeting)
	return nil
}

func genGreeting(ctx context.Context) (string, error) {
	// 使用WithTimeout包装，将在1s后自动取消返回的context，从而取消它传递该context的任何子函数
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	// ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()
	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "hello", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func printFarewell(ctx context.Context) error {
	farewell, err := genFarewell(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", farewell)
	return nil
}

func genFarewell(ctx context.Context) (string, error) {
	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "goodbye", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func locale(ctx context.Context) (string, error) {
	if deadline, ok := ctx.Deadline(); ok {
		if deadline.Sub(time.Now().Add(1*time.Minute)) <= 0 {
			return "", context.DeadlineExceeded
		}
	}
	select {
	case <-ctx.Done():
		// 这一行返回为什么Context被取消的原因。该错误会已知弹出到main，这会导致取消
		return "", ctx.Err()
	case <-time.After(1 * time.Minute):
		// case <-time.After(1 * time.Second):
	}
	return "EN/US", nil
}

func main() {
	var wg sync.WaitGroup
	// 创建context对象
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printGreeting(ctx); err != nil {
			fmt.Printf("cannot print greeting: %v\n", err)
			// 如果从打印问候语返回错误，main将取消这个context
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printFarewell(ctx); err != nil {
			fmt.Printf("cannot print farewell: %v\n", err)
		}
	}()
	wg.Wait()
}
