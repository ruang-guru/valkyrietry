package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ruang-guru/valkyrietry"
)

func main() {
	ctx := context.Background()

	options := []valkyrietry.Option{
		valkyrietry.WithMaxRetryAttempts(1),
		valkyrietry.WithRetryDelay(2 * time.Second),
		valkyrietry.WithRetryBackoffMultiplier(2),
		valkyrietry.WithJitter(0.2),
	}

	retryFunc := func() error {
		resp, err := http.Get("http://testingexample.com")
		if err != nil {
			fmt.Println("Request failed, will retry:", err)
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 500 {
			// Simulate server-side error
			return errors.New("server error, retrying")
		}
		fmt.Println("Request succeeded")
		return nil
	}

	// Use Valkyrietry to handle the retry logic
	if err := valkyrietry.Do(ctx, retryFunc, options...); err != nil {
		fmt.Println("Operation failed after retries:", err)
		return
	}

	fmt.Println("Operation successful")
}
