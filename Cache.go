package main

import (
    "log"
)

type HourCache struct {
    Minutes map[int]*MinuteCache
}

func NewHourCache() *HourCache {
    minuteMap := map[int]*MinuteCache{}
    for i:=0; i<60; i++ {
        minuteMap[i] = NewMinuteCache()
    }
    return &HourCache{Minutes: minuteMap}
}

type MinuteCache struct {
    Seconds map[int]float64
    written bool
}

func NewMinuteCache() *MinuteCache {
    secondMap := map[int]float64{}
    return &MinuteCache{Seconds: secondMap, written: true}
}

type Cache struct {
	Content map[string]*HourCache
}

func NewCache() *Cache {
    cache := &Cache{Content: map[string]*HourCache{}}
    return cache
}

func (self *Cache) Insert( Second int, Minute int, Id string, Data float64 ) {
    if self.Content[Id] == nil {
        self.Content[Id] = NewHourCache()
    }
    self.Content[Id].Minutes[Minute].Seconds[Second] = Data
    self.Content[Id].Minutes[Minute].written = false
}


func (self *Cache) print( Id string ) {
    for i:= 0; i<60; i++ {
        for j:=0; j<60; j++ {
            if self.Content[Id].Minutes[i].Seconds[j] != 0 {
                log.Println( self.Content[Id].Minutes[i].Seconds[j] )
                log.Printf( "%d, %d, %s \n", i, j, Id)
            }
        }
    }
}
