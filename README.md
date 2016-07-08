Redmine Issues Tracker
======

Want to contribute ? Feel free to send pull requests!

Have problems, bugs, feature ideas?
We are using the github [issue tracker](https://github.com/opstalent/tracker/issues) to manage them.

ABOUT
==================================================

Example call: `http://localhost:8080/issues??offset=0&limit=1000`

### Install And Run

```shell
$ go get github.com/opstalent/tracker
$ tracker -port=8080 username password 
```

### Add tracker as service [Debian]

```shell
$ sudo cp tracker.service /var/init.d/tracker
$ sudo nano /var/init.d/tracker
```

Change:
- your-redmine-user - redmine user
- your-redmine-password - redmine password
- service-port - tracker port

```shell
$ sudo chmod +x /var/init.d/tracker
$ sudo update-rc.d tracker defaults
$ sudo update-rc.d tracker enable
```

now see you service in

```shell
$ service --status-all
```

License
-------
This bundle is released under the MIT license. See the complete license in the bundle:

[LICENSE](LICENSE)
