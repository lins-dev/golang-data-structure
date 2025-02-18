package exercises

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"time"
)

func WithGoRoutines(){
	start := time.Now()
	const n = 10
	var wg sync.WaitGroup
	wg.Add(n)
	for range n {
		go func() {
			defer wg.Done()
			response, err := http.Get("https://google.com")
			if err != nil {
				panic(err)
			}
			defer response.Body.Close()
			fmt.Println("ok")
		}()
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}

func WithoutGoROutines(){
	start := time.Now()
	const n = 10
	for range n {
		response, err := http.Get("https://google.com")
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		fmt.Println("ok")
	}
	fmt.Println(time.Since(start))
}

func WithContext(){
	start := time.Now()
	const n = 10
	var wg sync.WaitGroup
	wg.Add(n)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()

	server := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(10*time.Second)
				fmt.Fprintln(w, "hello world")
			},
		),
	)

	for range n {
		go func(ctx context.Context) {
			defer wg.Done()

			request, err := http.NewRequestWithContext(
				ctx,
				"GET",
				server.URL,
				nil,
			)
			if err != nil {
				panic(err)
			}

			response, err := http.DefaultClient.Do(request)
			if err != nil{
				if errors.Is(err, context.DeadlineExceeded){
					fmt.Println("timeout")
					return
				}
				panic(err)
			}

			defer response.Body.Close()
		}(ctx)
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}