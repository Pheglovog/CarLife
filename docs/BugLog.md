
# 1.requires 1 of the 'Writers' sub-policies to be satisfied: permission denied
Error: got unexpected status: FORBIDDEN -- config update for existing channel did not pass final checks: implicit policy evaluation failed - 0 sub-policies were satisfied, but this policy requires 1 of the 'Writers' sub-policies to be satisfied: permission denied


[hyperledger fabric - Got "this policy requires 1 of the 'Writers' sub-policies to be satisfied: permission denied" error when creating a channel - Stack Overflow](https://stackoverflow.com/questions/59498185/got-this-policy-requires-1-of-the-writers-sub-policies-to-be-satisfied-permi)


# 2.Default client peer is down and no channel details available database
[Understanding Hyperledger Explorer Setup via Docker | by Davor Kljajic | Medium](https://medium.com/@dakljajic/understanding-hyperledger-explorer-setup-via-docker-6af845fcb82e)
## try1:
`docker network inspect network_name`
查看了net，explorer和fabric在同一个network中

最后发现是版本问题，使用3.0以下的镜像就好了

# 3.server认为我是来捣乱的

closing transport due to: connection error: desc = "error reading from server: EOF", received prior goaway: code: ENHANCE_YOUR_CALM, debug data: "too_many_pings"



找了半天，想起来了，我在gateway中建立连接的方式不对，我是希望建立一个grpc连接后，就可以一直使用，但是这样请求太频繁就会被认为是恶意的，所以我觉得应该使用一个请求建立一个grpc的方式，当然也可以有一个grpc pool可以让我们重复使用，但是这里为快速修复bug，就直接使用短链接的方法

可以看一下下面两个文章：
[rpc项目中的长连接与短连接的思考 - 思wu邪 - 博客园](https://www.cnblogs.com/swx123/p/17754469.html)
[gRPC 应用篇之客户端 Connection Pool - 熊喵君的博客 | PANDAYCHEN](https://pandaychen.github.io/2020/10/03/DO-WE-NEED-GRPC-CLIENT-POOL/)


# 4.无法背书
rpc error: code = Aborted desc = failed to endorse transaction, see attached details for more info
**猜想一：**
设定car tire时，背书不成功，但是我注册用户没问题，最后发现chaincode中有权限限制，
![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202411110013017.png)
这个显然不行，背书的时候大家都要执行，不然就达不成共识

fabric中private data可以指定如何指定背书节点，应该使用这样的机制，避免数据泄露

**猜想二：**
每次都是 10s后出错，我感觉是背书timeout设定有问题
![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202411110141952.png)
把orderer.yaml中的这几个都从10改到60
![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202411110146067.png)

还是不行

**猜想三**
![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202411110909815.png)
看了details，发现key为空，修改后，又出现了新的问题

![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202411110911711.png)
返现endorse的结果不一样，主要是因为：
`        Time:     time.Now(),`
所以要修改以下