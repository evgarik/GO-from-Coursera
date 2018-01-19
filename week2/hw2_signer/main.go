package main

// func main() {
// 	res := ""

// 	inputData := []int{0, 1, 1, 2, 3, 5, 8}

// 	hashSignJobs := []job{
// 		job(func(in, out chan interface{}) {
// 			for _, fibNum := range inputData {
// 				out <- fibNum
// 			}
// 		}),
// 		job(SingleHash),
// 		job(MultiHash),
// 		job(CombineResults),
// 		job(func(in, out chan interface{}) {
// 			dataRaw := <-in
// 			data, _ := dataRaw.(string)

// 			res = data
// 			fmt.Println(res)
// 		}),
// 	}

// 	ExecutePipeline(hashSignJobs...)

// 	time.Sleep(10 * time.Second)
// }

func main() {

	f := []job{
		job(func(in, out chan interface{}) {
			inputData := []int{0, 1}
			// inputData := []int{0, 1, 1, 2, 3, 5, 8}

			for _, fibNum := range inputData {
				out <- fibNum
			}
		}),
		job(SingleHash),
		job(MultiHash),
		job(CombineResults),
	}
	ExecutePipeline(f...)
}
