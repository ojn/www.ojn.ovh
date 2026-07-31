package main

import (
    "fmt"
    "log"
    "time"
    "github.com/gorilla/feeds"
)

func main() {
    now := time.Now()
    feed := &feeds.Feed{
        Title:       "ojn website",
        Link:        &feeds.Link{Href: "https://www.ojn.ovh"},
        Description: "ojn notes and rants about technology and open source",
        Author:      &feeds.Author{Name: "Gleb Toit", Email: "ojn@mailbox.org"},
        Created:     now,
    }

    feed.Items = []*feeds.Item{
        &feeds.Item{
            Title:       "Lorem Ipsum",
            Link:        &feeds.Link{Href: "https://www.ojn.ovh/lorem-ipsum/"},
            Description: "Lorem ipsum",
            Author:      &feeds.Author{Name: "Gleb Toit", Email: "ojn@mailbox.org"},
            Created:     time.Date(2026, time.July, 10, 22, 0, 0, 0, time.UTC),
        },
    }

    rss, err := feed.ToRss()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(rss)
}
