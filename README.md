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
### (2)&emsp;访问项目
浏览器访问 ip:9000


### (3)&emsp;示例



https://user-images.githubusercontent.com/45720137/187822669-6ec78734-cc96-4bab-80ac-009359e22fd0.mp4


### (3)&emsp;自定义配置
#### 拷贝配置文件
```
 docker cp -a go_file_server:/etc/goconf .
```
### 修改配置文件:  goconf/config.yaml (添加一个cc用户)
![image](https://user-images.githubusercontent.com/45720137/187824473-9efadbb1-ed91-4cac-8af3-a7926e51d8e8.png)

### 用当前配置覆盖docker内部配置(修改配置文件 docker-compose.yaml 删除红框中注释)
![image](https://user-images.githubusercontent.com/45720137/187824222-7aab9db8-a300-47f9-aaed-5bf25fb47452.png)

### 重启服务
```
docker-compose up -d
```


  
