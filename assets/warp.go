package assets

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	warp "github.com/PierreZ/Warp10Exporter"
	"github.com/PierreZ/startup-clicker-backend/templates"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

// createGTS is ensuring all the GTS are existing by pushing a 0 in it.
// As we are computing using bucketizer.sum all the time
func createGTS() {

	log.Info("Creating assets")

	for name := range All {
		gts := warp.NewGTS("asset").WithLabels(warp.Labels{"name": name}).AddDatapoint(time.Now(), 0)
		err := gts.Push("https://"+os.Getenv("ENDPOINT"), os.Getenv("WTOKEN"))
		if err != nil {
			panic(err)
		}
	}

	// Creating Money
	gts := warp.NewGTS("money").AddDatapoint(time.Now(), 0)
	err := gts.Push("https://"+os.Getenv("ENDPOINT"), os.Getenv("WTOKEN"))
	if err != nil {
		panic(err)
	}
}

func refreshAssets() {

	log.Println("refresh assets")

	var client = &http.Client{
		Timeout: time.Second * 10,
	}

	body := templates.GetAssets(os.Getenv("RTOKEN"))

	resp, err := client.Post("https://"+os.Getenv("ENDPOINT")+"/api/v0/exec", "", strings.NewReader(body))
	if err != nil {
		panic(err)
	}

	if resp.StatusCode > 200 {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%q", dump)
		os.Exit(1)
	}

	defer resp.Body.Close()

	assetsResponse := make([]map[string]float64, len(All))

	err = json.NewDecoder(resp.Body).Decode(&assetsResponse)
	if err != nil {
		panic(err)
	}
	assets := assetsResponse[0]
	for name, number := range assets {
		SetAssetNumber(name, number)
	}
	log.Println("map updated")

}

func refreshMoney() {

	log.Println("refresh money")

	var client = &http.Client{
		Timeout: time.Second * 10,
	}

	body := templates.GetMoney(os.Getenv("RTOKEN"))

	resp, err := client.Post("https://"+os.Getenv("ENDPOINT")+"/api/v0/exec", "", strings.NewReader(body))
	if err != nil {
		panic(err)
	}

	if resp.StatusCode > 200 {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%q", dump)
		os.Exit(1)
	}

	defer resp.Body.Close()

	money := make([]float64, 1)

	err = json.NewDecoder(resp.Body).Decode(&money)
	if err != nil {
		panic(err)
	}
	log.Println("money is at", money[0])
	account.Set(money[0])

}

func watch() {

beginning:
	u := url.URL{Scheme: "wss", Host: os.Getenv("ENDPOINT"), Path: "/api/v0/plasma"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	defer c.Close()
	defer close(done)

	err = c.WriteMessage(websocket.TextMessage, []byte("SUBSCRIBE "+os.Getenv("RTOKEN")+" ~.*{}"))
	if err != nil {
		panic(err)
	}

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Error(err)
			break
		}
		var classname string
		var labels map[string]string
		var value float64
		classname, labels, value, err = parseInputFormat(string(message))
		if err != nil {
			log.Panic(err)
		}

		//log.Println(classname, labels, value)

		switch classname {
		case "money":
			account.Set(value)
			log.Println("money updated")
		case "asset":
			Inc(labels["name"])
			log.Println("asset '" + labels["name"] + "' updated")
		default:
		}
	}
	goto beginning
}
func parseInputFormat(message string) (string, map[string]string, float64, error) {

	message = strings.Replace(message, "\n", "", -1)

	var classname string
	labels := make(map[string]string)
	var value float64
	var err error

	elts := strings.Split(message, " ")
	value, err = strconv.ParseFloat(elts[2], 64)
	if err != nil {
		return classname, labels, value, err
	}
	classname = strings.Split(elts[1], "{")[0]

	selector := strings.Split(elts[1], "{")[1]
	selector = selector[0 : len(selector)-1]

	l := strings.Split(selector, ",")
	for _, la := range l {
		lab := strings.Split(la, "=")
		if lab[0] != ".app" {
			labels[lab[0]] = lab[1]
		}
	}

	return classname, labels, value, err
}
