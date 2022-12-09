# file_server_project
### 一个go+vue写的前后端分离文件服务器
 <br>
<br>


### (1)&emsp;部署项目
#### 要求
  - docker and docker-compose
#### 开始部署
```
git clone https://github.com/sy13123/file_server_project.git
```
```
cd file_server_project
```
```
docker-compose up -d
```
#### 如果构建失败,关闭buildx功能再试
```
DOCKER_BUILDKIT=0 docker-compose up -d
```

### (2)&emsp;访问项目
浏览器访问 ip:9000
<br>
默认用户名:aa  密码:1


### (3)&emsp;示例



https://user-images.githubusercontent.com/45720137/187822669-6ec78734-cc96-4bab-80ac-009359e22fd0.mp4


### (4)&emsp;自定义配置

### 修改配置文件:  goconf/config.yaml (添加一个bb用户)
![image](https://user-images.githubusercontent.com/45720137/192236776-d0805618-bd65-4d54-9a87-46e33b3da4b0.png)

### 重启服务
```
docker-compose restart
```


  
