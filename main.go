package main

import (
    "fmt"
    "time"
    "log"
    "net/http"
    "github.com/PuerkitoBio/goquery"
    "github.com/reujab/wallpaper"
)

func main(){
    currentTime := time.Now()
    date := currentTime.Format("060102")
    url := fmt.Sprintf("https://apod.nasa.gov/apod/ap%s.html", date)

    resp, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    if resp.StatusCode != 200 {
            log.Fatalf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
        }

        doc, err := goquery.NewDocumentFromReader(resp.Body)

        if err != nil {
            log.Fatal(err)
        }
        var wallImage string
        var lastContext string
        doc.Find("center").Each(func(i int, a *goquery.Selection) {
            doc.Find("p").Each(func(i int, x *goquery.Selection) {
                doc.Find("a").Each(func(i int, y *goquery.Selection) {
                    doc.Find("img").Each(func(i int, s *goquery.Selection) {
                        context, ok := s.Attr("src")
                        if ok {
                            
                            if context != lastContext {
                                wallImage = fmt.Sprintf("https://apod.nasa.gov/apod/%s", context)
                            }
                            lastContext = context
                            
                        }
            })
            })
            })
            
            

        })
        wallpaper.SetFromURL(wallImage)
}



