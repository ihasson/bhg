package main

import "fmt"

func handler(i int, s *goquery.Selection) {
    url, ok := s.Find("a").Attr("href")
    if !ok {
        return
    }

    fmt.Printf("%d: %s\n", i, url)
    res, err := http.Get(url)
    if err != nil {
        return
    }
    buf, err := ioutil.ReadAll(re.Body)
    if err != nil {
        return
    }
    defer res.Body.Close()

    r, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
    if err != nil {
        return
    }

    cp, ap, err := metadata.NewProperties(r)
    if err != nil {
        return
    }

    log.Printf(
            "%25s %25s - %s %s\n",
            cp.Creator,
            cp.LastModifiedBy,
            ap.Application,
            ap.GetMajorVersion())
}

