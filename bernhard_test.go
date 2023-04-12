package bernhard

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"net/http"
	"testing"
	"time"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_PORT_OK    = "50001"
	_PORT_NOTOK = "50002"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type BernSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&BernSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *BernSuite) SetUpSuite(c *C) {
	go runServer(c, _PORT_OK, handlerOK)
	go runServer(c, _PORT_NOTOK, handlerNotOK)

	time.Sleep(time.Second)
}

func (s *BernSuite) TestClientCreation(c *C) {
	client, err := NewClient("0.0.1:" + _PORT_OK)

	c.Assert(err, NotNil)
	c.Assert(client, IsNil)
}

func (s *BernSuite) TestOk(c *C) {
	client, err := NewClient("http://127.0.0.1:" + _PORT_OK)

	c.Assert(err, IsNil)
	c.Assert(client, NotNil)

	err = client.Add(
		Alert{
			Service: "MyService",
			Host:    "app1",
			TTL:     600,
			State:   STATE_PROBLEM,
			Desc:    "Some error",
		},
	)

	c.Assert(err, IsNil)

	err = client.Send()

	c.Assert(err, IsNil)
}

func (s *BernSuite) TestSendError(c *C) {
	client, err := NewClient("http://127.0.0.1:50000")

	c.Assert(err, IsNil)
	c.Assert(client, NotNil)

	err = client.Add(
		Alert{
			Service: "MyService",
			Host:    "app1",
			TTL:     600,
			State:   STATE_PROBLEM,
			Desc:    "Some error",
		},
	)

	err = client.Send()

	c.Assert(err, NotNil)
}

func (s *BernSuite) TestServerError(c *C) {
	client, err := NewClient("http://127.0.0.1:" + _PORT_NOTOK)

	c.Assert(err, IsNil)
	c.Assert(client, NotNil)

	err = client.Add(
		Alert{
			Service: "MyService",
			Host:    "app1",
			TTL:     600,
			State:   STATE_PROBLEM,
			Desc:    "Some error",
		},
	)

	err = client.Send()

	c.Assert(err, NotNil)
}

func (s *BernSuite) TestEmptyStackSend(c *C) {
	client, err := NewClient("http://127.0.0.1:" + _PORT_OK)

	c.Assert(err, IsNil)
	c.Assert(client, NotNil)

	err = client.Send()

	c.Assert(err, IsNil)
}

func (s *BernSuite) TestAlertValidator(c *C) {
	a := Alert{}
	c.Assert(validateAlert(a), DeepEquals, ErrServiceNotSet)

	a = Alert{Service: "abcd"}
	c.Assert(validateAlert(a), DeepEquals, ErrHostNotSet)

	a = Alert{Service: "abcd", Host: "abcd"}
	c.Assert(validateAlert(a), DeepEquals, ErrStateNotSet)

	a = Alert{Service: "abcd abcd", Host: "abcd", State: "ok"}
	c.Assert(validateAlert(a), DeepEquals, ErrBadServiceName)

	client, _ := NewClient("http://127.0.0.1:" + _PORT_OK)

	c.Assert(client, NotNil)
	c.Assert(client.Add(Alert{}), NotNil)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func runServer(c *C, port string, handler func(http.ResponseWriter, *http.Request)) {
	server := &http.Server{
		Addr:    ":" + port,
		Handler: http.NewServeMux(),
	}

	server.Handler.(*http.ServeMux).HandleFunc("/events", handler)

	server.ListenAndServe()
}

func handlerOK(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Accepted\n"))
}

func handlerNotOK(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
	w.Write([]byte("Server error\n"))
}
