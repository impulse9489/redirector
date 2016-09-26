package main

import (
    "testing"
    "os"
    "net/http/httptest"
    "fmt"
    "strings"
)

func TestUnit_Redirect_WithPath(t *testing.T) {
    port := "8082"
    redirect_url := "https://www.wbplay.com"
    os.Args = []string{"cmd", port, redirect_url}
    recorder := httptest.NewRecorder()


    orig_url :="http://example.com"
    path := "/foo"
    orig_url_full := orig_url + path
    req := httptest.NewRequest("GET", orig_url_full, nil)
    redirect(recorder, req);
    
    new_url := recorder.HeaderMap["Location"][0]

    url_test := strings.Replace(orig_url_full, orig_url, redirect_url, 1)

    if new_url == url_test {
        fmt.Println("URL was redirected properly. URL: " + new_url) 

    } else {
       t.Errorf("URL not redirected properly. URL: " + new_url) 
    }


    fmt.Printf("HTTP Status Code: %v\n", recorder.Code)
    if recorder.Code != 301 {
        t.Errorf("HTTP Status incorrect.")
    }
}

func TestUnit_Redirect_LongPath(t *testing.T) {
    port := "8081"
    redirect_url := "https://www.wbplay.com"
    os.Args = []string{"cmd", port, redirect_url}
    recorder := httptest.NewRecorder()


    orig_url :="http://www.thisisalongpath.com"
    path := "/test1/test2/test3/"
    orig_url_full := orig_url + path
    req := httptest.NewRequest("GET", orig_url_full, nil)
    redirect(recorder, req);
    
    new_url := recorder.HeaderMap["Location"][0]

    url_test := strings.Replace(orig_url_full, orig_url, redirect_url, 1)

    if new_url == url_test {
        fmt.Println("URL was redirected properly. URL: " + new_url) 

    } else {
       t.Errorf("URL not redirected properly. URL: " + new_url) 
    }


    fmt.Printf("HTTP Status Code: %v\n", recorder.Code)
    if recorder.Code != 301 {
        t.Errorf("HTTP Status incorrect.")
    }
}

func TestUnit_Redirect_Base(t *testing.T) {
    port := "8086"
    redirect_url := "https://www.wbplay.com"
    os.Args = []string{"cmd", port, redirect_url}
    recorder := httptest.NewRecorder()


    orig_url :="http://www.thisisalongpath.com"
    path := ""
    orig_url_full := orig_url + path
    req := httptest.NewRequest("GET", orig_url_full, nil)
    redirect(recorder, req);
    
    new_url := recorder.HeaderMap["Location"][0]

    url_test := strings.Replace(orig_url_full, orig_url, redirect_url, 1)

    if new_url == url_test {
        fmt.Println("URL was redirected properly. URL: " + new_url) 

    } else {
       t.Errorf("URL not redirected properly. URL: " + new_url) 
    }


    fmt.Printf("HTTP Status Code: %v\n", recorder.Code)
    if recorder.Code != 301 {
        t.Errorf("HTTP Status incorrect.")
    }
}

func TestUnit_Redirect_Infinite(t *testing.T) {
    port := "8082"
    redirect_url := "https://www.wbplay.com"
    os.Args = []string{"cmd", port, redirect_url}
    recorder := httptest.NewRecorder()

    orig_url := redirect_url
    path := "/foo"
    orig_url_full := orig_url + path

    req := httptest.NewRequest("GET", orig_url_full, nil)
    redirect(recorder, req);
    
    fmt.Printf("HTTP Status Code: %v\n", recorder.Code)
    if recorder.Code != 400 {
        t.Errorf("HTTP Status incorrect.")
    }

}

func TestUnit_Redirect_OSArgs(t *testing.T) {
    //port := "8082"
    redirect_url := "https://www.wbplay.com"
    os.Args = []string{"cmd", redirect_url}
    recorder := httptest.NewRecorder()

    orig_url := redirect_url
    path := "/foo"
    orig_url_full := orig_url + path

    req := httptest.NewRequest("GET", orig_url_full, nil)
    redirect(recorder, req);
    body := recorder.Body.String
    fmt.Println(body)
    fmt.Printf("HTTP Status Code: %v\n", recorder.Code)
    if recorder.Code != 400 {
        t.Errorf("HTTP Status incorrect.")
    }
}

