# minit

**WORKING IN PROGRESS**

Minimal init daemon for container, support systemd service files

## Usage

```bash
minit redis nginx php-fpm
```

`minit` searchs `/etc/systemd/system`, `/lib/systemd/system`, `/usr/lib/systemd/system` for `redis.service`, `nginx.service` and `php-fpm.service` and run them.

`minit` IGNORE dependencies (`Requires=`, `After=`) BY DESIGN, because in most circumstances, they are not suitable for container environment.

## License

Yanke Guo <guoyk.cn@gmail.com>, MIT License
