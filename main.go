package main

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/http2"
)

var (
	ArrCoin = []string{"bitcoin", "litecoin", "dash",
		"ethereum", "ripple", "cardano",
		"eos", "monero", "bitcoin-cash"}
)

func HttpGet(name_symbol string, hostt int) ([]Ticker, error) {
	var hostname string
	client := http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	if hostt == 0 {
		hostname = "https://api.coinmarketcap.com/v1/ticker/" + name_symbol + "/"
	} else if hostt == 1 {
		hostname = "https://api.coinmarketcap.com/v1/ticker/"
	}
	resp, err := client.Get(hostname)
	if err != nil {
		return []Ticker{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("error: ", err)
	}
	res := make([]Ticker, 0)
	err = json.NewDecoder(strings.NewReader(string(body))).Decode(&res)
	if err != nil {
		log.Fatalln(err)
	}
	return res, err
}
func getCoins(arr0 []string) ([]Ticker, error) {
	var wg sync.WaitGroup
	var arrTicker []Ticker
	var err error
	if len(arr0) == 0 {
		res, err := HttpGet("", 1)
		if err != nil {
			return []Ticker{}, err
		}
		return res, err
	}
	wg.Add(len(arr0))
	for _, NameCoins := range arr0 {
		go func(n string) {
			defer wg.Done()
			res, err := HttpGet(n, 0)
			if err != nil {
				return
			}
			arrTicker = append(arrTicker, res[0])
		}(NameCoins)
	}
	wg.Wait()
	return arrTicker, err
}
func uniqueInt(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
func uniqueFloat64(intSlice []float64) []float64 {
	keys := make(map[float64]bool)
	list := []float64{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
func main() {
	log.Println("start...")
	arr0 := make([]float64, 0, 10000)
	cnt := 0
	for {
		cnt++
		start := time.Now()
		res, err := getCoins([]string{"ripple"})
		if err != nil {
			log.Println("error: ", err)
		}
		for _, n := range res {
			arr0 = append(arr0, n.Price_usd)
			sort.Float64s(arr0)
			log.Printf("name,price-arr0-max: %f", arr0[len(arr0)-1])
			sort.Sort(sort.Reverse(sort.Float64Slice(arr0)))
			log.Printf("name,price-arr0-min: %f", arr0[0])
			log.Println("name,price: ", n.Name, ":", n.Price_usd)
			if n.Price_usd > 5 || n.Price_usd > 10 {
				sendMessageBot(">4")

			}
		}
		log.Println("time elapsed: ", time.Since(start))
		log.Println("index for: ", cnt)
		time.Sleep(time.Minute * 10)
	}
}
