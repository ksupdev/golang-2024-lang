package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gofiber/fiber/v2"
)

// This Function will run before main
func init() {
	// hystrix.DefaultTimeout
	// hystrix.Configure()
	/*
		// DefaultTimeout is how long to wait for command to complete, in milliseconds
		DefaultTimeout = 1000
		// DefaultMaxConcurrent is how many commands of the same type can run at the same time
		DefaultMaxConcurrent = 10
		// DefaultVolumeThreshold is the minimum number of requests needed before a circuit can be tripped due to health
		DefaultVolumeThreshold = 20
		// DefaultSleepWindow is how long, in milliseconds, to wait after a circuit opens before testing for recovery
		DefaultSleepWindow = 5000
		// DefaultErrorPercentThreshold causes circuits to open once the rolling measure of errors exceeds this percent of requests
		DefaultErrorPercentThreshold = 50
		// DefaultLogger is the default logger that will be used in the Hystrix package. By default prints nothing.
	*/

	// Configuration for validating conditions to open the circuit breaker
	/*
		Open Circuit Breaker => Break requests to other APIs
		Close Circuit Breaker => Allow requests to other APIs
	*/
	hystrix.ConfigureCommand("api", hystrix.CommandConfig{
		Timeout:                500,
		RequestVolumeThreshold: 1,
		ErrorPercentThreshold:  100,
		SleepWindow:            15000,
	})

	hystrix.ConfigureCommand("api2", hystrix.CommandConfig{
		Timeout:                500,
		RequestVolumeThreshold: 1,
		ErrorPercentThreshold:  100,
		SleepWindow:            15000,
	})

	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(":8002", hystrixStreamHandler)
	//http://localhost:8002/
}

func main() {
	app := fiber.New()

	app.Get("/api", api)
	app.Get("/api2", api)

	app.Listen(":8001")

}

func api(c *fiber.Ctx) error {

	output := make(chan string, 1)
	hystrix.Go("api", func() error {
		res, err := http.Get("http://localhost:8000/api")
		if err != nil {
			return err
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		msg := string(data)
		fmt.Println(msg)

		output <- msg

		return nil
	}, func(err error) error {
		fmt.Println(err)
		return nil
	})

	out := <-output
	return c.SendString(out)

	// return c.SendString(msg)
}

/*
go get github.com/gofiber/fiber/v2
go get github.com/afex/hystrix-go/hystrix
go mod tidy
*/

/*
==> Install hystrix-dashboard
https://hub.docker.com/r/mlabouardy/hystrix-dashboard
docker pull mlabouardy/hystrix-dashboard
create docker-compose.yml
docker compose up -d

http://localhost:9002/hystrix
http://localhost:9002/hystrix/monitor?stream=http%3A%2F%2Fhost.docker.internal%3A8002&delay=100&title=Pu
*/

/*
The `CommandConfig` struct represents the configuration settings for a Hystrix command. Each command in Hystrix is associated with a specific piece of code that you want to isolate and protect, typically an external dependency call.

Here's the meaning of each field in the `CommandConfig` struct:

1. `Timeout`: This is the maximum time allowed for the execution of the command. If the command takes longer than this duration to execute, it will be considered as a timeout and Hystrix will attempt to cancel it.

2. `MaxConcurrentRequests`: This specifies the maximum number of concurrent requests allowed to execute the command. If this limit is reached, additional requests will be rejected or queued depending on Hystrix configuration.

3. `RequestVolumeThreshold`: This represents the minimum number of requests within a rolling window before Hystrix starts to apply its circuit breaker logic. If the number of requests within the window is less than this threshold, Hystrix will not take any action, regardless of errors or latency.

4. `SleepWindow`: This is the duration (in milliseconds) for which Hystrix will keep the circuit breaker open after tripping. During this period, Hystrix will reject requests without executing the command. After the sleep window expires, Hystrix will attempt to allow requests again to see if the service has recovered.

5. `ErrorPercentThreshold`: This specifies the error percentage threshold for the circuit breaker. If the percentage of failed requests within the rolling window exceeds this threshold, Hystrix will open the circuit breaker and start rejecting requests. This helps prevent cascading failures and allows the system to recover.

These configuration settings allow you to control the behavior of Hystrix commands and tune them according to the characteristics of your application and the external dependencies it relies on. By adjusting these parameters, you can achieve better resilience and fault tolerance in your system.
*/
