index-cli
========

A cli tool that access rnd-dockerhub.huawei.com.

## 编译

	$ make

## 构建好的二进制文件

	wget -c http://containeros.huawei.com/tools/index-cli


## 运行

	$ ./bin/index-cli -h
	NAME:
	   ./bin/index-cli - The cli tool that access rnd-dockerhub.huawei.com.
	
	USAGE:
		index-cli [global options] command [command options] [arguments...]
   
	VERSION:
	 0.0.0
	
	AUTHOR(S):
		 w00291922 <long.wanglong@huawei.com> 
   
	COMMANDS:
	  login	Login the hub
	  logout	Logout the hub
	  status	Get the status of the rnd-dockerhub
	  userinfo	Get the user info of yourself
	  user, u	Control user
      repo, r	Control the Repository
      search	Search rnd-dockerhub for images

	GLOBAL OPTIONS:
	--debug, -d			Debug logging
	--help, -h			show help
	--generate-bash-completion	
	--version, -v		print the version

## 使用说明

> **注意1：**该工具用来访问`rnd-dockerhub.huawei.com`,所有的命令中不需要出现该地址。工具默认访问地址都是`rnd-dockerhub.huawei.com`.

> **注意2：**该工具使用的用户鉴权信息来自文件`~/.docker/config.json`,跟社区的`hub.docker.com`兼容。所以使用前应该使用
> `docker login` 或者`index-cli login`登录。


### 查看 `rnd-dockerhub.huawei.com`的状态

	$ ./bin/index-cli status
	Unicorn Index version: deploy
	BuildDate: 2016-04-14.16:07:11
	APIversion: v1.0
	Index Server is OK

### 创建用户 `test-cli`

	$ ./bin/index-cli user create
	Please input USERNAME you want to create: test-cli
	Please input PASSWORD: 
	Please re-input PASSWORD: 
	Please input EMAIL: long.wanglong@huawei.com
	Please input PHONE: 123456
	Account created successfully. A verification email has been sent out.

>  **注意：** 这里的邮件一定要写真实的地址，因为hub会给该邮箱发送激活邮件。

### 修改`test-cli`用户的密码

	$ ./bin/index-cli user update
	Username: test-cli
	Please input NEW PASSWORD: 
	Please re-input NEW PASSWORD: 
	Please input EMAIL: long.wanglong@huawei.com
	Please input PHONE: 12345678
	OK

### 登录hub，获得鉴权信息

	$ ./bin/index-cli login
	Username: test-cli
	Password: 
	Login Succeeded

### 查看用户的基本信息

> **注意：** 只能查看当前登录hub的用户的信息。

	$ ./bin/index-cli userinfo
	User Name  : test-cli
	Namespace  : test-cli
	Product    : europa
	Quote      : 68.72 GB
	Used Space : 0 B
	Number Of Image         : 0
	Number of Image(private): 0
	Number of Image(protect): 0
	Number of Image(public) : 0

### 查找指定的docker镜像

	$ ./bin/index-cli search ubuntu
	NAME								          DESCRIPTION	STARS	OFFICIAL
	rnd-dockerhub.huawei.com/official/ubuntu-upstart				0	true
	rnd-dockerhub.huawei.com/official/ubuntu-debootstrap			0	true
	rnd-dockerhub.huawei.com/official/ubuntu						0	true
	rnd-dockerhub.huawei.com/library/ubuntu							0	true
	rnd-dockerhub.huawei.com/library/rancher-vm-ubuntu				0	true
	rnd-dockerhub.huawei.com/library/ubuntu-arm64					0	true
	rnd-dockerhub.huawei.com/library/hukeping-dev-ubuntu-arm64		0	true
	rnd-dockerhub.huawei.com/library/ubuntu-14.04					0	true
	rnd-dockerhub.huawei.com/unicorn/os-ubuntuconsole				0	false
	rnd-dockerhub.huawei.com/unicorn/unicorn_ubuntu					0	false
	rnd-dockerhub.huawei.com/unicorn/ubuntu							0	false

### 获取`rnd-dockerhub.huawei.com/official/ubuntu`镜像的详细信息

> **注意：** 不需要输入域名`rnd-dockerhub.huawei.com`


	$ ./bin/index-cli repo get official/ubuntu
	Image Name         : official/ubuntu
	Image Size         : 345.9 MB
	Number of Images   : 6
	Access Level       : public
	Number of Download : 38
	NUM	NAME:TAG										SIZE
	1	rnd-dockerhub.huawei.com/official/ubuntu:latest	65.86 MB
	2	rnd-dockerhub.huawei.com/official/ubuntu:14.04	65.86 MB
	3	rnd-dockerhub.huawei.com/official/ubuntu:trusty	65.86 MB
	4	rnd-dockerhub.huawei.com/official/ubuntu:15.10	49.61 MB
	5	rnd-dockerhub.huawei.com/official/ubuntu:15.04	49.34 MB
	6	rnd-dockerhub.huawei.com/official/ubuntu:vivid	49.34 MB

### 设置镜像的访问权限为public

	$ ./bin/index-cli repo set --access public test-cli/buxybox 
	Set test-cli/buxybox Access Level to public: SUCCESS

### 设置镜像的描述信息

	$ ./bin/index-cli repo set --description "just for test"  test-cli/buxybox
	Set repo's Description success


> **注意**：也可以同时设置访问权限和描述信息

	$ ./bin/index-cli repo set --access private --description "just for test1"  test-cli/buxybox 
	Set test-cli/buxybox Access Level to private: SUCCESS
	Set repo's Description success
	$　./bin/index-cli repo get  test-cli/buxybox 
	Image Name         : test-cli/buxybox
	Image Size         : 2.028 MB
	Number of Images   : 3
	Descripton         : just for test1
	Access Level       : private
	Number of Download : 0
	NUM	NAME:TAG										SIZE
	1	rnd-dockerhub.huawei.com/test-cli/buxybox:test1	675.9 kB
	2	rnd-dockerhub.huawei.com/test-cli/buxybox:test2	675.9 kB
	3	rnd-dockerhub.huawei.com/test-cli/buxybox:test3	675.9 kB

### 删除指定tag的镜像`rnd-dockerhub.huawei.com/test-cli/buxybox:test1`

	$ ./bin/index-cli repo rmi test-cli/buxybox:test1
	Delete image(test-cli/buxybox) with tha tag(test1) success.

> **注意**:参数的格式必须为`namespace/reponame:tag`

### 删除整个仓库`rnd-dockerhub.huawei.com/test-cli/buxybox`

	$ ./bin/index-cli repo del  test-cli/buxybox
	Delete repo(test-cli/buxybox) success.

### 登出hub，取消鉴权信息

	$ ./bin/index-cli logout
    Remove login credentials for rnd-dockerhub.huawei.com
