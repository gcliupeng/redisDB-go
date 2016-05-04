package main
import (
	"fmt"
	"redisDB/lib"
)
func main() {
	r,_:=redisDB.NewRedis("","6379")
	s,_:=r.Get("name")
	fmt.Printf("%s\n",s)

	r.Set("name2","liupeng")

	r.HSet("key","age",28)

	smap,_:=r.HGetAll("key")
	for k,v := range smap{
		fmt.Printf("%s-----%s\n",k,v)
	}

	r.HMSet("key","home","guizhou","like","eat")
	smap,_=r.HGetAll("key")
	for k,v := range smap{
		fmt.Printf("%s-----%s\n",k,v)
	}
}