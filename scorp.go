package scorp

import 	"github.com/parnurzeal/gorequest"
import 	"encoding/json"
import 	"strings"
import 	"strconv"

type Api struct {
	Username string
	Password string
	limit int       //Default 1
	links [3]string
	ScorpApi Scorp
	request *gorequest.SuperAgent
}


type User struct {
	FirstName string `json:"first_name"`
	LastName string	`json:"last_name"`
	Id int	`json:"id"`
	Username string	`json:"username"`

}

type Scorp struct  {
	User User `json:"user"`
	profile_picture string `json:"profile_picture"`
	website string `json:"website"`
	received_like_count int `json:"received_like_count"`
	info string `json:"info"`
	small_picture string `json:"small_picture"`
	following_count int `json:"following_count"`
	follower_count int `json:"follower_count"`
	view_count int `json:"view_count"`
	post_count int `json:"post_count"`
	like_count int `json:"like_count"`
	gender string `json:"gender"`
}



func (api *Api) Start(username string,password string) Api  {

	//initialize Struct
	api.links[0] = "https://scorpapp.com/login" //Login link
	api.links[1] = "https://scorpapp.com/api/profile/" //Profile link
	api.links[2] = "https://scorpapp.com/logout" //logout link

	api.limit = 1; //Default 1 it can be change

	api.Username = username
	api.Password = password


	api.login()


	return *api

}



func (api *Api) login(){

	str := []string{`{"username":"`, api.Username, `","password":"`,api.Password,`"}`}
	api.request = gorequest.New()
	_,_,_=api.request.Post(api.links[0]).
	Send(strings.Join(str, "")).
	End()

}


func (api *Api) GetUser(id int) Scorp {

	ids := strconv.Itoa(id)
	str := []string{api.links[1],ids}
	_,body,_ := api.request.Get(strings.Join(str, "")).End()
	var dat Scorp

	byt := []byte(body)
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	api.ScorpApi = dat;
	return api.ScorpApi
}


func (api *Api) Close() {
	api.request.Get(api.links[2]).End()
}

