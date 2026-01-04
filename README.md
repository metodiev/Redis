# Redis
Redis is an in-memory key-value store used for caching, fast lookups, and temporary data storage. It’s often used in repositories to speed up data access or reduce load on the primary database.

## Install Redis in Ubuntu or Debian based distro:

```bash
sudo apt update
sudo apt install redis-server -y
sudo systemctl enable redis-server
sudo systemctl start redis-server
redis-cli
```
### Test Redis:
```bash
SET mykey "Hello Redis"
GET mykey
```

## Install redis Using Docker contaienr

```bash
docker pull redis
docker run --name my-redis -p 6379:6379 -d redis
```

- Test with Redis CLI (inside the container or another terminal):

```bash
docker exec -it my-redis redis-cli
SET mykey "Hello Docker Redis"
GET mykey

```

## Install redis using precompiled version form GitHub

https://github.com/tporadowski/redis/releases
