package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
    "regexp"
)

var validPath = regexp.MustCompile("^/(ota)/([\\.a-zA-Z0-9]+)$")

type Artifact struct {
    Version string
    Content []byte
}

func (af *Artifact) save() error {
    artifact_name := af.Version + ".artf"
    return ioutil.WriteFile(artifact_name, af.Content, 0600)
}

func getArtifact(ver string) (*Artifact, error) {
    artifact_name := ver + ".artf"
    content, err := ioutil.ReadFile(artifact_name)
    if err != nil {
        return nil, err
    }
    return &Artifact{Version: ver, Content: content}, nil
}

func handler(w http.ResponseWriter, r *http.Request, version string) {
    atfcontent, err := ioutil.ReadFile(version + ".artf")
    fmt.Println("url path: ", r.URL.Path[:1])    
    fmt.Println("method: ", r.Method)
    fmt.Println("header: ", r.Header)
    fmt.Println("body: ", r.Body)
    fmt.Println("contentLength: ", r.ContentLength)
    fmt.Println("host: ", r.Host)
    fmt.Println("remoteAddr: ", r.RemoteAddr)
    fmt.Println("requestURI: ", r.RequestURI)
    rq := r.URL.RawQuery
    qv, err := url.ParseQuery(rq)
    fmt.Println("formValue: ", qv)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Hi visitor, fw content: ", string(atfcontent))
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        m := validPath.FindStringSubmatch(r.URL.Path)
        if m == nil {
            http.NotFound(w, r)
            return
        }
        fn(w, r, m[2])
    }
}


func main() {
    a1 := &Artifact{Version: "1.0.0", Content: []byte("This  is poc fw")}

    a1.save()
    a2, _ := getArtifact("1.0.0")

    fmt.Println(string(a2.Content))

    http.HandleFunc("/ota/", makeHandler(handler))
    log.Fatal(http.ListenAndServe(":8080", nil))
}


