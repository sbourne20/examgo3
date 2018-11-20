/*
Documentation :
http://golang.org
https://medium.com/@kelvin_sp/building-and-testing-a-rest-api-in-golang-using-gorilla-mux-and-mysql-1f0518818ff6
https://www.youtube.com/watch?v=SqrbIlUwR0U&t=315s
https://www.youtube.com/watch?v=DWNozbk_fuk&t=421s
https://www.linkedin.com/learning/learning-go-for-web-development/using-gorilla-mux

*/

package main

func main() {
	a := App{}
	a.Initialize("root", "@Bourne20", "bkd")
	//a.Initialize("root", "Master199", "bkd")
	a.Run(":8080")
}
