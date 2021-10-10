package main

import (
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

type Strategy interface {
	NextEndpoint() url.URL
	SetEndpoint([]url.URL)
}

type RandomStrategy struct {
	endpoints []url.URL
}

func (r *RandomStrategy) NextEndpoint() url.URL{
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	return r.endpoints[r1.Intn(len(r.endpoints))]
} // 랜던함 값의 엔드포인트를 선정해서 리턴


func (r *RandomStrategy) SetEndpoint(endpoints []url.URL) {
	r.endpoints = endpoints
} // 해당 엔드포인트를 배열에 담아둠 

type LoadBalancer struct {
	strategy Strategy
}

func NewLoadBalancer(s Strategy, endpoints []url.URL) *LoadBalancer {
	s.SetEndpoint(endpoints)
	return &LoadBalancer{strategy: s}
} // 신규 인스턴스 로드밸런서를 리턴

func (l *LoadBalancer) GetEndpoint() url.URL {
	return l.strategy.NextEndpoint()
} // 로드 밸런서에서 엔드포인트를 꺼내옴

func main() {
	endpoints := []url.URL{
		{Host: "www.google.com"},
	  {Host: "www.google.co.uk"},
	}

	lb := NewLoadBalancer(&RandomStrategy{}, endpoints)

	fmt.Println(lb.GetEndpoint())
}