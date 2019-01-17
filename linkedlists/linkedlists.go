package main

import (
	"fmt"
	"time"
)

type Feed struct {
	Length int
	Start  *Post
	End    *Post
}

type Post struct {
	data          string
	publishedDate int64
	next          *Post
}

var feed = &Feed{}

func Append(data string, timeStamp int64) {
	post := &Post{
		data:          data,
		publishedDate: timeStamp,
		next:          nil,
	}
	if feed.Length == 0 {
		feed.Start = post
		feed.End = post
	} else {
		feed.End.next = post
		feed.End = post
	}

	feed.Length++
}

func Insert(data string, timeStamp int64) {
	if feed.Length == 0 {
		feed.Start = &Post{
			data:          data,
			publishedDate: timeStamp,
			next:          nil,
		}

		feed.End = feed.Start
	} else {
		var previousPost, currentPost *Post
		currentPost = feed.Start
		for currentPost.publishedDate < timeStamp {
			if currentPost.next == nil {
				break
			}
			previousPost = currentPost
			currentPost = currentPost.next
		}

		previousPost.next = &Post{
			data:          data,
			publishedDate: timeStamp,
			next:          currentPost,
		}
		if currentPost.next == nil { // edge case
			feed.End = currentPost
		}

	}

	feed.Length++
}

func Remove(publishedDate int64) {
	if feed.Length == 0 {
		panic("Feed List is empty")
		return
	} else {
		var (
			previousPost *Post
			currentPost  *Post
		)

		if feed.Start.publishedDate == publishedDate { // edge case
			feed.Start = feed.Start.next
			feed.Length--
			return
		}

		previousPost = feed.Start
		currentPost = feed.Start.next

		for currentPost.publishedDate != publishedDate {
			if currentPost.next == nil { // edge case
				panic("No more posts with publishedDate")
				return
			}
			previousPost = currentPost
			currentPost = currentPost.next
		}

		previousPost.next = currentPost.next // Go runtime GC tick will remove unused varibles
		if currentPost.next == nil {         // edge case
			feed.End = currentPost
		}

		feed.Length--
	}
}

func Search(timeStamp int64) *Post {
	if feed.Length == 0 {
		panic("Feed list is empty")
		return nil
	} else {
		var currentPost *Post
		currentPost = feed.Start
		for currentPost.publishedDate != timeStamp {
			if currentPost.next == nil {
				panic("No more posts with publishedDate")
				return nil
			}
			currentPost = currentPost.next
		}
		return currentPost
	}
}

func PrintFeed() {
	if feed.Length == 0 {
		panic("Feed list is empty")
		return
	} else {
		var currentPost = feed.Start
		for currentPost.next != nil {
			fmt.Print(currentPost.data, currentPost.publishedDate, " ")
			currentPost = currentPost.next
		}
		fmt.Print(currentPost.data, currentPost.publishedDate, " ")
	}
	fmt.Println()
}

func main() {
	timeStamp1 := time.Now().Add(time.Second * 1).Unix()
	Append("happy birthday x!!!!!", timeStamp1)

	timeStamp2 := time.Now().Add(time.Second * 2).Unix()
	Append("happy birthday y!!!!!", timeStamp2)

	timeStamp3 := time.Now().Add(time.Second * 3).Unix()
	Append("happy birthday z!!!!!", timeStamp3)

	PrintFeed()

	Insert("happy birthday a!!!!!", timeStamp2)

	PrintFeed()

	Remove(timeStamp2)

	PrintFeed()
}
