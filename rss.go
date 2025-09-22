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
            Created:     time.Date(2009, time.November, 10, 22, 0, 0, 0, time.UTC),
        },
        &feeds.Item{
            Title:       "Logic-less Template Redux",
            Link:        &feeds.Link{Href: "https://www.ojn.ovh/logicless-template-redux/"},
            Description: "More thoughts on logicless templates",
            Author:      &feeds.Author{Name: "Gleb Toit", Email: "ojn@mailbox.org"},
            Created:     time.Date(2009, time.November, 11, 12, 0, 0, 0, time.UTC),
        },
        &feeds.Item{
            Title:       "Idiomatic Code Reuse in Go",
            Link:        &feeds.Link{Href: "https://www.ojn.ovh/idiomatic-code-reuse-in-go/"},
            Description: "How to use interfaces <em>effectively</em>",
            Author:      &feeds.Author{Name: "Gleb Toit", Email: "ojn@mailbox.org"},
            Created:     time.Date(2009, time.November, 14, 05, 0, 0, 0, time.UTC),
        },
    }

    rss, err := feed.ToRss()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(rss)
}
