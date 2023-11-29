<div align="center">

<img src="https://cdn.discordapp.com/attachments/1179065448428490932/1179066374962810951/valkyrietry-logo.png?ex=65786e21&is=6565f921&hm=5288a4f0a21464c2065c047f4b5b8c11346e1c6871333c6a5425f80e10d83f6a&" align="center" width="144px" height="144px"/>

# -- VALKYRIETRY --
### Fail, Retry, Succeed With Valkryrietry
</div>

## Overview
Valkyrietry is a robust and intuitive GoLang library designed to bring resilience to your applications with its powerful retry mechanisms. Inspired by the legendary Valkyries, symbols of strength and determination, Valkyrietry embodies these attributes to ensure your operations don't just fail, but have the strength to retry and ultimately succeed.

## Key Features
- **Simplicity**: Valkyrietry is built with a focus on simple API, making it straightforward to integrate into your existing GoLang projects.
- **Ease of Use**: With minimal configuration and sensible defaults, you can start implementing advanced retry logic in no time.
- **Resilience**: Crafted to withstand and recover from failures, Valkyrietry ensures your applications remain robust under various failure scenarios.
- **Flexible Retry Mechanism**: Customizable retry strategies, including configurable delay intervals, max retry attempts, backoff algorithms, and jitter for more natural retry patterns.

## Getting Started
To get started with Valkyrietry, simply install the library in your GoLang project and start wrapping your critical operations with Valkyrietry's retry functionality.

```go
package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/confus1on/valkyrietry"
)

func main() {
	options := []valkyrietry.Option{
		valkyrietry.WithMaxRetryAttempts(5),
		valkyrietry.WithRetryDelay(0.5 * time.Second) // Google use 0.5 as initial retry,
		valkyrietry.WithRetryBackoffMultiplier(1.5) // Google also use 1.5 for default multiplier,
		valkyrietry.WithJitter(0.5),
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
	if err := valkyrietry.Do(retryFunc, options...); err != nil {
		fmt.Println("Operation failed after retries:", err)
		return
	}

	fmt.Println("Operation successful")
}
```
---

### Why Valkyrietry?

When considering retry mechanisms in GoLang, there are several options available, each with its strengths. Here's how `Valkyrietry` stands out in:

1. **Simplicity and Ease of Use**: 
   - **Valkyrietry**: Designed with a strong emphasis on user-friendly interfaces. It makes implementing retry logic straightforward, offering sensible defaults and intuitive configuration.

2. **Customization and Flexibility**: 
   - **Valkyrietry**: Provides a balance of ease-of-use and customization. It allows fine-tuning of retry strategies, including jitter and backoff configurations, tailored to diverse operational requirements.

3. **Resilience and Robustness**: 
   - **Valkyrietry**: Built with resilience at its core. It ensures robust retry mechanisms even in complex failure scenarios, making your applications more reliable.

### Conclusion:
`Valkyrietry` is an excellent choice for developers looking for a retry library that combines simplicity with powerful customization options. It's particularly well-suited for those who appreciate a balance between easy-to-use interfaces and the ability to handle complex retry scenarios robustly. With its unique theme and user-focused design, Valkyrietry offers an appealing and reliable solution for managing retries in your GoLang applications.

---