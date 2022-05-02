package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Register() error {
	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8000/register", nil)
	if err != nil {
		return err
	}
	req.Header.Set("address", "127.0.0.1:8000")
	req.Header.Set("service", "testClient")
	cli := http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(res))
	fmt.Println(resp.Status)
	return nil
}

func (c *Client) GetServe() error {
	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8000/get-serve", nil)
	if err != nil {
		return err
	}
	req.Header.Set("service", "testClient")
	cli := http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	address, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(address))
	fmt.Println(resp.Status)
	return nil
}
