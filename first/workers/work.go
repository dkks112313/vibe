package workers

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

type Workers struct {
	wg    sync.WaitGroup
	Count uint
	file  *os.File
}

func (w *Workers) StartupWorkers(job func(string, chan []net.IP, chan error), name string) {
	channelForErrors := make(chan error, 10)
	channelForNetIp := make(chan []net.IP, 10)

	for _, v := range w.ReadFile(name) {
		w.wg.Add(1)
		go func() {
			defer w.wg.Done()

			job(v, channelForNetIp, channelForErrors)
		}()
	}

	go func() {
		w.wg.Wait()
		close(channelForNetIp)
		close(channelForErrors)
	}()

	for channelForNetIp != nil && channelForErrors != nil {
		select {
		case ip, ok := <-channelForNetIp:
			if !ok {
				channelForNetIp = nil
				continue
			}
			fmt.Println(ip)
		case err, ok := <-channelForErrors:
			if !ok {
				channelForErrors = nil
				continue
			}
			fmt.Println("Error", err)
		}
	}
}

func (w *Workers) openFile(name string) {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	w.file = f
}

func (w *Workers) ReadFile(name string) []string {
	w.openFile(name)
	defer w.file.Close()

	var strokes []string
	reader := bufio.NewReader(w.file)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		strokes = append(strokes, string(line))
	}

	return strokes
}
