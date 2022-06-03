package telegram

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"simple-web-golang-server/lib/e"
	"strconv"
)

type Client struct{
	host string
	basePath string
	client http.Client
}

func New(host string,token string) Client{
	return Client{
		host: host,
		basePath: NewBasePath(token),
		client: http.Client{},
	}
}

func NewBasePath(token string)string{
	return "bot" + token
}

func (c *Client)Updates(offset int, limit int)([]Update,error){
	q:= url.Values{}
	q.Add("offset",strconv.Itoa(offset))
	q.Add("limit",strconv.Itoa(limit))
	// do request <- GetUpdates
}


func (c *Client)DoRequest(method string,query url.Values)(data []byte,err error){
	defer func(){err = e.WrapIfErr("cant do request",err) }()
	u:= url.URL{
		Scheme: "https",
		Host: c.host,
		Path: path.Join(c.basePath,method),
	}
	req,err := http.NewRequest(http.MethodGet, u.String(),nil)
	if err != nil{
		return nil,err
	}

	req.URL.RawQuery= query.Encode()

	resp,err := c.client.Do(req)
	if err != nil{
		return nil,err
	}
	defer func(){_=resp.Body.Close()
	}()
	body,err := io.ReadAll(resp.Body)
	if err != nil{
		return nil,err
	}
	return body,nil
}

func (c *Client)SendMessage(){

}

