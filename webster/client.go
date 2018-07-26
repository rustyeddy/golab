package magoo

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type MagooClient struct {
	baseurl        string
	url            string // full url
	*http.Response        // the response we will send
}

var (
	localurl string
)

func init() {
	localurl = "http://localhost:1199/magoo/"
}

// NewMagooClient returns a structure ready to communicate with
// your favorite magoo server, whereever that may be
func NewMagooClient(u string) *MagooClient {
	return &MagooClient{
		baseurl:  localurl,
		url:      "",
		Response: nil,
	}
}

// Geturl returns the baseurl + args
func (mc *MagooClient) Geturl(args string) string {
	return mc.baseurl + args
}

// Get sends a getrequest to the specfied server
func (mc *MagooClient) Get(args string) (resp *http.Response, err error) {
	mc.url = mc.baseurl + args
	resp, err = http.Get(mc.url)
	if err != nil {
		log.Printf("  failed to get %s -> %v", mc.url, err)
		return nil, err
	}
	defer resp.Body.Close()

	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, fmt.Errorf("  failed to read body of response %v", err)
	}
	log.Println("  <<<<<<<<< body >>>>>>>>>  ")
	log.Printf("\n%s\n", string(body))
	log.Println("  >>>>>>>>> end  <<<<<<<<<  ")
	return resp, nil
}

// Post a request to the specified server.
func (mc *MagooClient) Post(url string, args map[string][]string) (resp *http.Response, err error) {
	resp, err = http.PostForm(url, args)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, fmt.Errorf("  failed to read body of response %v", err)
	}
	log.Println("  <<<<<<<<< POST body >>>>>>>>>  ")
	log.Printf("\n%s\n", string(body))
	log.Println("  >>>>>>>>>  the end  <<<<<<<<<  ")
	return resp, nil
}

func (mc *MagooClient) GetMagoo() (*http.Response, error) {
	return mc.Get("")
}

func (mc *MagooClient) GetEntries() (*http.Response, error) {
	r, e := mc.Get("entry")
	return r, e
}

func (mc *MagooClient) GetEntry(id int64) (*http.Response, error) {
	r, e := mc.Get("entry/" + string(id))
	return r, e
}

func (mc *MagooClient) PostEntry(args map[string][]string) (*http.Response, error) {
	r, e := mc.Post("entry/", args)
	return r, e
}

func DoMagooClient() {

	var resp *http.Response
	var err error

	mc := NewMagooClient("")
	resp, err = mc.GetMagoo()
	if err != nil {
		log.Printf("failed to get a magoo client, exiting ... ")
		log.Fatal(err)
	}
	log.Printf("GET /magoo => %+v", resp)

	resp, err = mc.GetEntries()
	if err != nil {
		log.Printf("failed to get entries exiting ... ")
		log.Fatal(err)
	}
	log.Printf("GET /magoo/entry => %+v", resp)
	vals := map[string][]string{
		"MagooId":   {"MrMagoo"},
		"EntryId":   {"0"},
		"Name":      {"tstform"},
		"FormId":    {"0"},
		"TransId":   {string(GetNextId())},
		"Timestamp": {GetTimeStamp()},
	}

	resp, err = mc.Post("submit/", vals)
	if err != nil {
		log.Printf("failed to post an entry to a magoo client, exiting ... ")
		log.Fatal(err)
	}
	log.Printf("POST /magoo/submit/ => %+v", resp)

	// resp, err === return if error
	resp, err = mc.GetEntry(0)
	if err != nil {
		log.Printf("failed to get entry named tstform ... ")
		log.Fatal(err)
	}
	log.Printf("GET /magoo/entry/{NameId} => %+v", resp)
}
