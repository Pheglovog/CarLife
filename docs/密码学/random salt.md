![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202410011156488.png)
如果两个人的密码相同，那么他们地hash值就相同，那么黑客就可以通过尝试更加轻易地获取密码

破解密码的方式：
暴力攻击：尝试所有的密码
字典攻击：将字典中常见单词的hash值与密码的hash值进行对比，破解密码

通过向密码中加入random salt，我们可以让字典的映射出现随机性，更难破解
![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202410011212231.png)

![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202410011212891.png)

![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202410011214011.png)
