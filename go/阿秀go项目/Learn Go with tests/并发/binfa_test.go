package concurrency

import (
    "reflect"
    "testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
    if url == "waat://furhurterwe.geds" {
        return false
    }
    return true
}

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
    results := make(map[string]bool)

	for _, url := range urls {
        go func(u string) {
            results[u] = wc(u)
        }(url)
    }
	time.Sleep(2 * time.Second)
    return results
}

func TestCheckWebsites(t *testing.T) {
    websites := []string{
        "http://google.com",
        "http://blog.gypsydave5.com",
        "waat://furhurterwe.geds",
    }

    actualResults := CheckWebsites(mockWebsiteChecker, websites)

    want := len(websites)
    got := len(actualResults)
    if want != got {
        t.Fatalf("Wanted %v, got %v", want, got)
    }

    expectedResults := map[string]bool{
        "http://google.com":          true,
        "http://blog.gypsydave5.com": true,
        "waat://furhurterwe.geds":    false,
    }

    if !reflect.DeepEqual(expectedResults, actualResults) {
        t.Fatalf("Wanted %v, got %v", expectedResults, actualResults)
    }
}


func slowStubWebsiteChecker(_ string) bool {
    time.Sleep(20 * time.Millisecond)
    return true
}

func BenchmarkCheckWebsites(b *testing.B) {
    urls := make([]string, 100)
    for i := 0; i < len(urls); i++ {
        urls[i] = "a url"
    }

    for i := 0; i < b.N; i++ {
        CheckWebsites(slowStubWebsiteChecker, urls)
    }
}
