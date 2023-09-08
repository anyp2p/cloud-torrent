package server

/*
handle realtime client library:
/js/velox.js

handle realtime client connections:
/sync

搜索资源:
/search

获取(GET)/删除(DELETE)已下载文件
/download/<filepath>

-- API [POST] --

convert url into torrent bytes:
/api/url (POST body: torrent url)

convert torrent bytes into magnet:
/api/torrentfile (POST body: torrent bytes)

修改配置:
/api/configure (POST body: json encode engine.Config)

根据磁力链下载:
/api/magnet (POST body: magnet url)

控制种子的开始和停止:
/api/torrent  (POST body: <start|stop>:<infohash>)

控制文件的开始和停止:
/api/file (POST body: <start|stop>:<infohash>:<filepath>)
*/
