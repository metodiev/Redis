import redis

#Connect to redis

r = redis.Redis(host='localhost', port=6379, db=0)

#String
r.set('name', 'Miroslav')
print(r.get('name'))

#Hash
r.hset('user:1', mapping={'name': 'Alice', 'age': 30})
print(r.hgetall('user:1'))

#List
r.lpush('mylist', 'a', 'b', 'c')
print(r.lrange('mylist', 0, -1))