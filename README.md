# file_server_project
### 一个go+vue写的前后端分离文件服务器
 <br>
<br>


## (1)&emsp;部署项目
- ### clone项目
    ```
    git clone https://github.com/sy13123/file_server_project.git
    ```

- ### docker-compose方式开始部署
    ```
    cd file_server_project

    docker-compose up -d
    ```
- ### kubernetes方式部署
    ```
    cd file_server_project

    kubectl create configmap fileserver-conf --from-file=goconf/config.yaml --from-file=nginx_conf/default.conf
    
    kubectl apply -f fileserver-k8s-deploy.yaml
    ```


## (2)&emsp;访问项目
浏览器访问 ip:39000
<br>
默认用户名:aa  密码:1


## (3)&emsp;示例



https://user-images.githubusercontent.com/45720137/187822669-6ec78734-cc96-4bab-80ac-009359e22fd0.mp4


## (4)&emsp;自定义配置

### 修改配置文件:  goconf/config.yaml (添加一个bb用户)
![image](https://user-images.githubusercontent.com/45720137/192236776-d0805618-bd65-4d54-9a87-46e33b3da4b0.png)

### 重启服务
```
docker-compose restart
```


  
