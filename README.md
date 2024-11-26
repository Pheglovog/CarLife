# CarLife


通过对车辆原件供应，制造，销售，保险，维修整个流程的全程记录，让车辆信息全程透明，方便用户使用和维护，同时也可以方便转售

技术栈为：Hyperledger Fabric+golang+Gin+gRPC+Docker+Docker Compose+Vue+TypeScript


# 1.软件整体架构

![框架.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202411262231075.png)


# 2.网络配置

整个网络中共分为几个组织：
- 供应商组织
	- 轮胎
	- 电池
	- 车身
- 制造商组织
- 销售组织
- 保险组织
- 维修组织
- 用户，输入车辆编号就可以查看车辆的所有信息
![网络结构.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202411262257708.png)




## 1. 网络中的组织

| 组织   | name               | 作用          |
| ---- | ------------------ | ----------- |
| 原件供应 | component_supplier | 供应轮胎，车身，内饰等 |
| 制造   | manufacturer       | 使用原件制造车辆    |
| 销售   | store              | 售卖车辆        |
| 保险   | insurer            | 为车辆提供保险     |
| 维修   | maintenancer       | 车辆保养维修      |
| 用户   | consumer           | 车主          |

## 2. 每个组织的成员

### component_supplier

| node        | type    | g port | a/c port | o port | 作用   |
| ----------- | ------- | ------ | -------- | ------ | ---- |
| orderer1    | orderer | 7050   | 8050     | 9050   | 排序   |
| cartire     | peer    | 7051   | 8051     | 9051   | 提供轮胎 |
| carbody     | peer    | 7052   | 8052     | 9052   | 提供车身 |
| carinterior | peer    | 7053   | 8053     | 9053   | 提供内饰 |

### manufacturer

| node     | type    | g port | a/c port | o port | 作用   |
| -------- | ------- | ------ | -------- | ------ | ---- |
| orderer2 | orderer | 7054   | 8054     | 9054   | 排序   |
| factory1 | peer    | 7055   | 8055     | 9055   | 一号工厂 |
| factory2 | peer    | 7056   | 8056     | 9056   | 二号工厂 |
### store

| node    | type | g port | a/c port | o port | 作用    |
| ------- | ---- | ------ | -------- | ------ | ----- |
| sailer1 | peer | 7057   | 8057     | 9057   | 一号专卖店 |
| sailer2 | peer | 7058   | 8058     | 9058   | 二号专卖店 |
### insurer
| node     | type    | g port | a/c port | o port | 作用   |
| -------- | ------- | ------ | -------- | ------ | ---- |
| orderer3 | orderer | 7059   | 8059     | 9059   | 排序   |
| pcompany | peer    | 7060   | 8060     | 9060   | 平安保险 |
| rcompany | peer    | 7061   | 8061     | 9061   | 人寿保险 |

### maintenancer
| node   | type | g port | a/c port | o port | 作用    |
| ------ | ---- | ------ | -------- | ------ | ----- |
| fixer1 | peer | 7062   | 8062     | 9062   | 一号维修店 |
| fixer2 | peer | 7063   | 8063     | 9063   | 二号维修店 |
### consumer
| node  | type | g port | a/c port | o port | 作用  |
| ----- | ---- | ------ | -------- | ------ | --- |
| user1 | peer | 7064   | 8064     | 9064   | 用户1 |

## 3.CA机构

| name                  | org                | s port | o port |
| --------------------- | ------------------ | ------ | ------ |
| ca_tls                | tls                | 17050  | 18050  |
| ca_component_supplier | component_supplier | 17051  | 18051  |
| ca_manufacturer       | manufacturer       | 17052  | 18052  |
| ca_store              | store              | 17053  | 18053  |
| ca_insurer            | insurer            | 17054  | 18054  |
| ca_maintenancer       | maintenancer       | 17055  | 18055  |
| ca_consumer           | consumer           | 17056  | 18056  |

# 3.数据结构与流程

## 1.User

| name     | type     | usage    |
| -------- | -------- | -------- |
| UserID   | string   | 唯一标识符    |
| UserType | UserType | 类型       |
| Password | string   | 密码       |
| CarList  | []*Car   | 有关的Car列表 |
## 2.Car

| name     | type        | usage          |
| -------- | ----------- | -------------- |
| CarID    | string      | 唯一标识符，同时也作为溯源码 |
| Tires    | CarTires    | 四个轮胎的信息        |
| Body     | CarBody     | 车身信息           |
| Interior | CarInterior | 内饰信息           |
| Manu     | CarManu     | 制造相关信息         |
| Store    | CarStore    | 车辆销售信息         |
| Insure   | CarInsure   | 车辆保险信息         |
| Maint    | CarMaint    | 车辆维修信息         |
| Owner    | *User       | 拥有者            |
| Record   | CarRecord   | 交易记录           |
## 3.CarTires

| name     | type    | usage |
| -------- | ------- | ----- |
| Time     | time    | 制造时间  |
| Width    | float32 | 轮胎宽度  |
| Radius   | float32 | 轮胎半径  |
| Workshop | string  | 生产车间  |

##  4.CarBody

| name     | type    | usage |
| -------- | ------- | ----- |
| Material | string  | 车身材料  |
| Time     | time    | 制造时间  |
| Weight   | float32 | 车身重量  |
| Color    | string  | 车身颜色  |
| Workshop | string  | 生产车间  |
## 5. CarInterior

| name     | type   | usage |
| -------- | ------ | ----- |
| Time     | time   | 制造时间  |
| Color    | string | 内饰颜色  |
| Material | string | 内饰材料  |
| WorkShop | string | 生产车间  |

## 6.CarManu

| name     | type   | usage |
| -------- | ------ | ----- |
| Time     | time   | 拼装时间  |
| Workshop | string | 生产车间  |

## 7.CarStore

| name  | type    | usage |
| ----- | ------- | ----- |
| Time  | time    | 销售时间  |
| Store | string  | 门店    |
| Cost  | float32 | 费用    |
| User  | string  | 车主    |
## 8.CarInsure

| name    | type     | usage |
| ------- | -------- | ----- |
| Insures | []Insure | 车辆的保险 |
Insure

| name      | type    | usage |
| --------- | ------- | ----- |
| BeginTime | time    | 开始时间  |
| EndTime   | time    | 结束时间  |
| Name      | string  | 保险名   |
| Cost      | float32 | 花费    |
## 9.CarMaint

| name   | type    | usage |
| ------ | ------- | ----- |
| Part   | string  | 维修部分  |
| Extent | string  | 破损程度  |
| Cost   | float32 | 花费    |

## 10.CarRecord

| name    | type     | usage |
| ------- | -------- | ----- |
| Records | []Record | 转手记录  |
Record

| name    | type   | usage |
| ------- | ------ | ----- |
| OldUser | string | 老车主   |
| NewUser | string | 新车主   |
| Cost    | string | 转手费用  |
| Time    | time   | 转手时间  |
