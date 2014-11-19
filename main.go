package main

import "github.com/dustin/go-humanize"
import "github.com/tj/nsqtop/pkg/nsqd"
import "github.com/segmentio/go-log"
import "github.com/tj/go-gracefully"
import "github.com/tj/docopt"
import stdlog "log"
import "io/ioutil"
import "time"
import "fmt"

var Version = "0.1.0"

const Usage = `
  Usage:
    nsqtop [--interval n] [--nsqd-http-address a...]
    nsqtop -h | --help
    nsqtop --version

  Options:
    -a, --nsqd-http-address a  nsqd http address [default: 0.0.0.0:4151]
    -i, --interval n           refresh interval [default: 1s]
    -h, --help                 output help information
    -v, --version              output version

`

func main() {
	stdlog.SetOutput(ioutil.Discard)

	args, err := docopt.Parse(Usage, nil, true, Version, false)
	log.Check(err)

	addrs := args["--nsqd-http-address"].([]string)

	d, err := time.ParseDuration(args["--interval"].(string))
	log.Check(err)

	go loop(d, addrs)

	gracefully.Shutdown()
}

func loop(d time.Duration, addrs []string) {
	for _ = range time.Tick(d) {
		for _, addr := range addrs {
			nsq := nsqd.New(addr)

			stats, err := nsq.Stats()
			log.Check(err)

			fmt.Printf("\033[2J\033[0f")
			fmt.Printf("\n\n\n\033[1m%30s\033[0m\n", addr)
			fmt.Printf("%30s %30s %15s %15s %15s %15s\n", "topic", "channel", "depth", "in-flight", "deferred", "timeouts")
			for _, topic := range stats.Topics {
				fmt.Printf("%30s %30s %15s %15s %15s %15s\n",
					topic.Name,
					"∙",
					humanize.Comma(topic.Depth),
					humanize.Comma(topic.InFlightCount),
					humanize.Comma(topic.DeferredCount),
					humanize.Comma(topic.TimeoutCount))

				for _, channel := range topic.Channels {
					fmt.Printf("%30s %30s %15s %15s %15s %15s\n",
						"∙",
						channel.Name,
						humanize.Comma(channel.Depth),
						humanize.Comma(channel.InFlightCount),
						humanize.Comma(channel.DeferredCount),
						humanize.Comma(channel.TimeoutCount))
				}
			}
		}
	}
}
