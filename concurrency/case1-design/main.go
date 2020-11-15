// https://ithelp.ithome.com.tw/articles/10208936
package main

func main() {
	// 	sw := &StatisticWorker{
	// 		quene: make(chan *http.Request, 1000),
	// 		num:   100,
	// 	}
	// 	go sw.Run()

	// 	http.HandleFunc("/api/query", func(w http.ResponseWriter, r *http.Request) {

	// 		//假設需要統計 request 相關數據，所以丟背景
	// 		//這邊使用這個非阻塞式的worker
	// 		err := sw.Add(r)
	// 		if err != nil {
	// 			log.Println(err)
	// 			return
	// 		}

	// 		u := &UserInfo{
	// 			Name: "syhlion",
	// 			Age:  18,
	// 		}
	// 		b, err := json.Marshal(u)
	// 		if err != nil {
	// 			log.Println(err)
	// 			return
	// 		}
	// 		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	// 		w.WriteHeader(http.StatusOK)
	// 		w.Write(b)
	// 	})

	// 	log.Fatal(http.ListenAndServe(":8080", nil))
	// }

	// type StatisticWorker struct {
	// 	quene chan *http.Request
	// 	num   int
	// }

	// //統一由這邊做 requset 加入的動作
	// func (s *StatisticWorker) AddRequest(r *http.Request) (err error) {
	// 	// select {
	// 	// case s.quene <- r:
	// 	// //當quene已滿時，則把 err 丟出，
	// 	// default:
	// 	// 	err = errors.New("buffer full")
	// 	// }
	// 	// return

	// }

	// //實際worker本體 運作邏輯的地方
	// func (s *StatisticWorker) Statistic() {
	// 	for r := range s.quene {
	// 		//處理 request 相關統計
	// 	}
	// }

	// //初始化到底要有幾個worker
	// func (s *StatisticWorker) Run() {
	// 	for i := 0; i < s.num; i++ {
	// 		go Statistic()
	// 	}
}
