package redisDB
import(
	"github.com/garyburd/redigo/redis"
)

type Redis struct{
	redis.Conn
}

func NewRedis(ip ,port string) (*Redis,error){
	c,err:=redis.Dial("tcp",ip+":"+port)
	if err!=nil{
		return nil,err
	}
	return &Redis{c},nil
}

func (r *Redis) Get(key string) (string,error) {
	reply,err:=r.Do("get",key)
	if err!=nil{
		return "",err
	}
	s,err:=redis.String(reply,nil)
	if err!=nil{
		return "",err
	}
	return s,nil
}

func (r *Redis) Set(key,value string) error {
	_,err:=r.Do("set",key,value)
	return err
}

func (r *Redis) HGet(key,field string) (string,error) {
	reply,err:=r.Do("hget",key,field)
	if err!=nil{
		return "",err
	}
	s,err:=redis.String(reply,nil)
	if err!=nil{
		return "",err
	}
	return s,nil
}

func (r *Redis) HGetAll(key string) (map[string]string,error) {
	reply,err:=r.Do("hgetall",key)
	if err!=nil{
		return nil,err
	}
	smap,err:=redis.StringMap(reply,nil)
	if err!=nil{
		return nil,err
	}
	return smap,nil
	
}

func (r *Redis) HSet(key,field,value interface{}) (error) {
	_,err:=r.Do("hset",key,field,value)
	return err
}

func (r *Redis) HMSet(key string, fieldvalue ...interface{}) (error) {
	params:=[]interface{}{key}
	params=append(params,fieldvalue...)
	_,err:=r.Do("hmset",params...)
	return err
}
