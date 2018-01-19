package main

import (
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

func SingleHash(in, out chan interface{}) {
	mux := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	do := func(num int) {
		wg.Add(1)
		data := strconv.Itoa(num)
		res := DataSignerCrc32(data) + "~"

		mux.Lock()
		md5 := DataSignerMd5(data)
		mux.Unlock()

		res += DataSignerCrc32(md5)
		out <- res
		wg.Done()
	}

	timeout := time.After(time.Millisecond)

LOOP:
	for {
		select {
		case input := <-in:
			num, _ := input.(int)
			go do(num)
		case <-timeout:
			break LOOP
		}
	}

	wg.Wait()
	close(out)
}

func MultiHash(in, out chan interface{}) {
	do := func(str string) {
		wg := &sync.WaitGroup{}
		wg.Add(6)

		var tmp [6]string

		hesh := func(i int) {
			tmp[i] = DataSignerCrc32((strconv.Itoa(i) + str))
			wg.Done()
		}

		for i := 0; i < 6; i++ {
			go hesh(i)
		}

		wg.Wait()
		res := tmp[0] + tmp[1] + tmp[2] + tmp[3] + tmp[4] + tmp[5]

		out <- res
	}

	// timeout := time.After(time.Second)

	// LOOP:

	for {
		select {
		case msg := <-in:

			str, _ := msg.(string)
			go do(str)
			// case <-timeout:
			// break LOOP
		}
	}
}

func CombineResults(in, out chan interface{}) {
	var tmp []string

	for i := 0; i < 7; {
		select {
		case msg := <-in:
			str, ok := msg.(string)
			if ok {
				tmp = append(tmp, str)
			}
			i++
			// case <-timer.C:
			// 	break LOOP
		}
	}

	sort.Strings(tmp[:])

	res := strings.Join(tmp[:], "_")
	out <- res
}

func ExecutePipeline(freeFlowJobs ...job) {

	in := make(chan interface{}, 1)
	out := in

	for _, v := range freeFlowJobs {
		out = make(chan interface{}, 1)
		go v(in, out)
		in = out
	}

	time.Sleep(4 * time.Second)
}
