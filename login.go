package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
)

type OAuth2Request struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (self *GlobalConfig) OAuth2Request() *OAuth2Request {
	return &OAuth2Request{
		self.ClientId,
		self.ClientSecret,
	}
}

type OAuth2Response struct {
	Id          string `json:"id"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func (self *OAuth2Response) Bytes() []byte {
	it, err := json.Marshal(self)
	if err != nil {
		panic(err)
	}
	return it
}

func (self *OAuth2Response) Save(filename string) {
	err := ioutil.WriteFile(filename, self.Bytes(), 0755)
	if err != nil {
		panic(err)
	}
}

func (self *OAuth2Response) Load() *OAuth2Response {
	self.LoadFrom(metaFilePath(".token"))
	return self
}

func (self *OAuth2Response) LoadFrom(path string) *OAuth2Response {
	file, _ := os.Open(path)
	body, err := ioutil.ReadAll(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	json.Unmarshal(body, self)
	return self
}

func (self *OAuth2Response) Decode(res *HttpResponse) *OAuth2Response {
	d := json.NewDecoder(res.Body)
	err := d.Decode(&self)
	if err != nil {
		panic(err)
	}
	return self
}

func retrieveAppAdminAccessToken() *OAuth2Response {
	token := globalConfig.OAuth2Request()
	headers := globalConfig.HttpHeaders("application/json")
	res := HttpPostJson("/oauth2/token", headers, token)
	oauth2res := &OAuth2Response{}
	return oauth2res.Decode(res)
}

func LoginAsAppAdmin() error {
	res := retrieveAppAdminAccessToken()
	res.Save(metaFilePath(".token"))
	return nil
}