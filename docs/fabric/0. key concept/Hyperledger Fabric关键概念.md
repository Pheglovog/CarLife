
# 什么是区块链

For now, it’s enough to think of a blockchain as a **shared**, **replicated** transaction system which is **updated via smart contracts** and kept consistently synchronized through a collaborative process called **consensus.**

# 什么是Hyperledger Fabric

 Rather than declaring a single blockchain standard, it encourages a collaborative approach to developing blockchain technologies via a community process, with intellectual property rights that encourage open development and the adoption of key standards over time.

特点：private， permissioned

the members of a Hyperledger Fabric network enroll through a trusted **Membership Service Provider (MSP)**

## shared ledger
Hyperledger Fabric has a ledger subsystem comprising two components: the **world state** and the **transaction log**.

- The world state component describes the state of the ledger at a given point in time. *It’s the database of the ledger.* 
- The transaction log component records all transactions which have resulted in the current value of the world state;
简单来书， world state中有ledger在每个时间段的状态，比如说交易总数，每个账户的余额，channel成员等等； transaction log就是包含了每笔交易的详细信息，比如说，A买了B一只鸡，给了B一块钱这样的信息

## smart Contracts

Hyperledger Fabric smart contracts are written in **chaincode** and are invoked by an application external to the blockchain when that application needs to interact with the ledger.

In most cases, chaincode interacts only with the database component of the ledger, the world state (querying it, for example), and not the transaction log.

客户端需要交易，就调用SC和 world state进行交互


## privacy
Hyperledger Fabric supports networks where privacy (using channels) is a key operational requirement as well as networks that are comparatively open.
封闭和半开放，全开放可由需求而定

## Consensus

Transactions must be written to the ledger in the order in which they occur, even though they might be between different sets of participants within the network. **For this to happen, the order of transactions must be established and a method for rejecting bad transactions that have been inserted into the ledger in error (or maliciously) must be put into place.**

就是大家同步账本的方法


# Fabric Model
## Assets
Assets can range from the tangible (real estate and hardware) to the intangible (contracts and intellectual property). Hyperledger Fabric provides the ability to modify assets using chaincode transactions.
资产就是一个人的财产，可以修改，一般是key-value形式保存

## Chaincode
Chaincode is software defining an asset or assets, and the transaction instructions for modifying the asset(s); in other words, it’s the business logic.

Chaincode execution results in a set of key-value writes (write set) that can be submitted to the network and applied to the ledger on all peers.
这不就是共识吗， 链码会产生一个write set，并提交到channel去修改所有peer的账本

## ledger Features
There is one ledger per channel. Each peer maintains a copy of the ledger for each channel of which they are a member.

features：
- Transactions consist of the versions of keys/values that were read in chaincode (read set) and keys/values that were written in chaincode (write set)
- Peers validate transactions against endorsement policies and enforce the policies
- Prior to appending a block, a versioning check is performed to ensure that states for assets that were read have not changed since chaincode execution time
- A channel’s ledger contains a configuration block defining policies, access control lists, and other pertinent information（orderer，配置整个channel， 懂了吧）
- Channels contain [Membership Service Provider](https://hyperledger-fabric.readthedocs.io/en/release-2.5/glossary.html#msp) instances allowing for crypto materials to be derived from different certificate authorities（MSP的不同实例都可以提供认证）

## Privacy
Hyperledger Fabric employs an immutable ledger on a per-channel basis, as well as *chaincode that can manipulate and modify the current state of assets*
 In order to solve scenarios that want to bridge the gap between total transparency and privacy, chaincode can be installed only on peers that need to access the asset states to perform reads and writes。
 就是说peer访问账本是需要通过chaincode的，如果没有这个账本的chaincode，就无法访问和修改。这样信息就可以限制在特定的peer之间，同时对于这些peer所有信息都是透明的。

When a subset of organizations on that channel need to keep their transaction data confidential, a **private data collection** (collection) is used to segregate this data in a private database, logically separate from the channel ledger, accessible only to the authorized subset of organizations.
channel下面的更小团体可以使用private data collection，将账本分区，这块区域只能由这个小团体读写。private data collection可以使用chaincode来实现

- **channels** keep transactions private from the broader network 
- **private data collections** keep data private between subsets of organizations on the channel.

## Security & Membership Services
**Public Key** Infrastructure is used to generate cryptographic certificates which are tied to organizations, network components, and end users or client applications.

## Consensus
然而，共识不仅包括简单地就交易顺序达成一致，两者的区别在Hyperledger Fabric 中得到了凸显，从提案和背书到排序、验证和提交，共识在整个交易流程中都扮演着核心角色。

consensus is defined as the full-circle verification of the correctness of a set of transactions comprising a block.

当区块中交易的顺序和结果满足明确的策略标准检查时，最终会达成共识。

除了众多的背书、验证和版本检查外，交易流的各个方向上还会发生持续的身份验证。

总而言之，共识并不仅仅局限于对一批交易商定顺序；相反，它的首要特征是从交易提案到提交的过程中不断进行核查而附带实现的。


# 区块链网络

区块链网络是一个为应用程序提供账本及智能合约（chaincode）服务的技术基础设施。首先，智能合约被用来生成交易，接下来这些交易会被分发给网络中的每个节点，这些交易会被记录在他们的账本副本上并且是不可篡改的。

在大多数的情况下，多个[组织](https://hyperledger-fabric.readthedocs.io/zh-cn/latest/glossary.html#organization) 会聚集到一起作为一个[通道]，在通道上，交易在链码上被发起。而他们的权限是由一套在通道最初配置的时候联盟成员都同意的[规则](https://hyperledger-fabric.readthedocs.io/zh-cn/latest/glossary.html#policy)来决定的。并且，配置的规则可以在通道中的组织同意的情况下随时地被改变。
![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202405191406161.png)

## 创建网路
创建网络或通道的第一步是同意并定义其配置:
![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202405191407722.png)

通道配置CC1已由组织R1、R2和R0同意，并包含在称为“配置块”的块中，该块通常由`configtxgen`工具根据`configtx.yaml`文件创建。

此配置块包含可以连接组件并在通道上交互的组织的记录，以及定义如何做出决策和达到特定结果的结构的**策略**。

这些组织的定义及其管理员的身份，必须由与每个组织相关联的证书颁发机构(CA)创建。
- 首先，区块链网络的不同组件使用证书来相互识别自己来自特定组织。这就是为什么区块链网络通常有不止一个CA的原因——不同的组织通常使用不同的CA。证书到成员组织的映射是通过称为[成员服务提供者(MSP)](https://github.com/hyperledger/fabric-docs-i18n/blob/release-2.5/docs/locale/zh_CN/source/membership/membership.html)的结构实现，该结构通过创建与根CA证书绑定的MSP来定义组织，以识别由根CA创建的组件和身份。
- 其次，稍后我们将看到由ca颁发的证书是如何成为[交易](https://github.com/hyperledger/fabric-docs-i18n/blob/release-2.5/docs/locale/zh_CN/source/glossary.html#transaction)生成和验证过程的核心。具体来说，X.509证书用于客户端应用程序[交易提案](https://github.com/hyperledger/fabric-docs-i18n/blob/release-2.5/docs/locale/zh_CN/source/glossary.html#proposal)和智能合约[交易响应](https://github.com/hyperledger/fabric-docs-i18n/blob/release-2.5/docs/locale/zh_CN/source/glossary.html#response)对[交易](https://github.com/hyperledger/fabric-docs-i18n/blob/release-2.5/docs/locale/zh_CN/source/glossary.html#transaction)进行数字签名。随后，拥有账本副本的网络节点在接受交易之前验证交易签名是有效的。

## 节点加入通道

节点是网络的基本元素，因为它们拥有账本和链码(包含智能合约)，而且是在通道上进行交易的组织连接到另一个通道(另一个是应用程序)的物理节点之一。

另一方面，排序服务从应用程序中收集经过背书的交易，并将它们排序到交易块中，这些交易块随后分发到通道中的每个peer节点。在每个提交节点上，记录交易，并适当更新账本的本地副本。
排序服务对于特定通道是唯一的，服务于该通道的节点也称为“同意集”。

![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202405191446189.png)
R1的节点P1和R2的节点P2，以及R0的排序服务O，通过[创建通道](https://github.com/hyperledger/fabric-docs-i18n/blob/release-2.5/docs/locale/zh_CN/source/create_channel/create_channel_participation.html)教程中描述的过程加入通道。注意，虽然只有一个排序节点1加入到该通道，*但在生产场景中，排序服务应至少包含三个节点。*

在将排序服务加入到通道之后，可以向通道配置提出更新并提交更新，但除此之外就没什么了。想进行更多的操作，您必须在通道上安装、批准和提交链码。

## 安装，批准和提交链码

链码（这就是智能合约呗）被安装到peer节点，然后被定义和提交到通道：

![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202405191453261.png)

在Fabric中，定义节点组织如何与账本交互的业务逻辑(例如，更改资产所有权的交易)包含在**智能合约**中。包含智能合约的结构称为**链码(chaincode)**，安装在相关的peer节点上，由相关节点的组织批准，并在通道上提交。

安装、批准和提交链码的过程被称为链码的“生命周期”。

在链码定义中提供的最重要的信息是[背书策略](https://github.com/hyperledger/fabric-docs-i18n/blob/release-2.5/docs/locale/zh_CN/source/glossary.html#endorsement-policy)。它描述了哪些组织必须在交易被其他组织接受到其账本副本之前对交易进行背书。根据用例，可以将背书策略设置为通道中成员的任何组合。
链码具有在通道上的成员之间创建[私有数据交易](https://github.com/hyperledger/fabric-docs-i18n/blob/release-2.5/docs/locale/zh_CN/source/private_data_tutorial.html)的能力

## 在通道中使用应用程序
提交智能合约后，客户端应用程序可以通过Fabric Gateway服务(网关)调用链码上的事务。这就完成了我们在第一张图中展示的结构:
![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202405191502707.png)

就像peer节点和排序节点一样，客户端应用程序具有将其与组织关联起来的身份。在我们的示例中，客户端应用程序A1与组织R1关联，并连接到C1。

交易提案作为链码的输入，链码使用它来生成交易响应。

## 加入多个通道
### 创建新的通道
![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202405191508626.png)


和前面一样，既然已经创建了通道配置CC2，那么可以说该通道在**逻辑上**是存在的，即使没有节点加入。

### 将组件加入新通道
节点+链码+应用

![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202405191509431.png)

请注意，虽然C1和C2拥有相同的排序组织(R0)，但不同的排序节点为每个通道提供服务。这不是强制性配置，因为即使相同的排序节点连接到多个通道，每个通道也有一个单独的排序服务实例，并且在多个排序组织聚集在一起为排序服务提供节点的通道中更为常见。
注意，只有加入到特定通道的排序节点才拥有该通道的账本。

虽然R2也有可能部署一个新的peer节点来连接到通道C2，但在本例中，它们选择将P2部署到C2。


从逻辑上讲，这与C1的创建非常相似。两个peer节点组织与一个排序组织一起创建一个通道，并将组件和链码连接到该通道。


## 将组织添加到现有的通道
我们将描述如何将组织R3添加到通道C1的配置CC1的过程。

**注意，权利和权限是在通道级别定义的。** 仅仅因为一个组织是一个通道的管理员，并不意味着它将是另一个通道的管理员。
![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202405191518635.png)

向通道添加一个新组织在高层次上是一个三步的过程:
- 决定新组织的权限和角色
- 1. 更新通道，包括相关的链码，以反映这些决策。
- 1. 组织将其peer节点(以及潜在的排序节点)加入通道，并开始参与交易。

在本主题中，我们假设R3将以与R1和R2相同的权限和地位加入C1。**同样，R3也将作为S5链码的背书者加入，这意味着R1或R2必须重新定义S5(具体来说，是链码定义的背书策略部分)并在通道上批准它。**

更新通道配置将创建一个新的**配置块CC1.1**，它将作为通道配置，直到再次更新。注意，即使配置已经更改，通道仍然存在，P1和P2仍然连接到它。不需要重新向通道中添加组织或节点。

下图展示了一次性所有组件都加入的过程。（其实也可以一次加一点）
![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202405191523889.png)


# 身份

区块链网络中的不同参与者包括 Peer 节点、排序节点、客户端应用程序、管理员等。**每一个参与者（网络内部或外部能够使用服务的活动元素）都具有封装在 X.509 数字证书中的数字身份。** 这些身份确实很重要，因为它们**确定了对资源的确切权限以及对参与者在区块链网络中拥有的信息的访问权限。**

此外，数字身份还具有 Fabric 用于确定权限的一些其他属性，并且它为身份和关联属性的并集提供了特殊名称——**主体** 。当我们谈论主体时，它们是决定其权限的属性。

要使身份可以被**验证**，它必须来自**可信任的**权威机构。[成员服务提供者](https://github.com/hyperledger/fabric-docs-i18n/blob/release-2.5/docs/locale/zh_CN/source/membership/membership.html)（Membership Service Provider，MSP）是 Fabirc 中可以信任的权威机构。具体地说，**一个 MSP 是定义管理该组织有效身份规则的组件。**

Fabric 中默认的 MSP 实现使用 X.509 证书作为身份，采用传统的公钥基础结构（Public Key Infrastructure,PKI）分层模型（稍后将详细介绍PKI）。

## 类比场景

想象你去超市购买一些杂货。在结账时，你会看到一个标志，表明只接受 Visa，Mastercard 和 AMEX 卡。如果你尝试使用其他卡付款（我们称之为“想象卡”）无论该卡是否真实、或你的帐户中是否有足够的资金，都无关紧要。它不会被接受。
![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202405191706396.png)

PKI 证书和 MSP 提供了类似的功能组合。PKI 就像一个卡片提供商，它分配了许多不同类型的可验证身份。另一方面，MSP 类似于商店接受的卡提供商列表，确定哪些身份是商店支付网络的可信成员（参与者）。**MSP 将可验证的身份转变为区块链网络的成员** 。

## 什么是PKI

**公钥基础结构（PKI）是一组互联网技术，可在网络中提供安全通信。** 是 PKI 将 **S** 放在 **HTTPS** 中，如果你在网页浏览器上阅读这个文档，你可能正使用 PKI 来确保它来自一个验证过的来源。
![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202405191719819.png)

A PKI is comprised of Certificate Authorities who issue digital certificates to parties (e.g., users of a service, service provider), who then use them to authenticate themselves in the messages they exchange in their environment。
也就是说，CA证书发给主体（principal），这个主体此后标识自己身份就使用这个证书了。

CA 的证书撤销列表（CRL）构成不再有效的证书的参考。证书的撤销可能由于多种原因而发生。

PKI 有四个关键要素：
- **数字证书**
- **公钥和私钥**
- **证书授权中心**
- **证书撤销列表**

## 数字证书

数字证书是包含与证书持有者相关的属性的文档。最常见的证书类型是符合 [X.509标准](https://en.wikipedia.org/wiki/X.509)的证书，它允许在其结构中编码一些用于身份识别的信息。

可以将 X.509 证书视为无法改变的数字身份证。

## 授权，公钥和私钥

身份验证和消息完整性是安全通信中的重要概念。身份验证要求确保交换消息的各方创建特定消息的身份。对于具有“完整性”的消息意味着在其传输期间不能被修改。

数字签名机制要求每一方保存两个加密连接的密钥：广泛可用的公钥和充当授权锚的私钥，以及用于在消息上产生**数字签名**的私钥 。数字签名消息的接收者可以通过检查附加签名在预期发送者的公钥下是否有效来验证接收消息的来源和完整性。

密钥之间唯一的数学关系使得私钥在消息上的签名，只有对应公钥在相同的消息上才可以与之匹配。

![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202405191732414.png)

## 证书授权中心
如你所见，人员或节点能够通过由系统信任的机构为其发布的**数字身份**参与区块链网络。在最常见的情况下，数字身份（或简称**身份**）的形式为，符合 X.509 标准并由证书授权中心（CA）颁发的经加密验证的数字证书。
![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202405191739849.png)

_证书授权中心向不同的参与者颁发证书。这些证书由 CA 进行签名，并将参与者的公钥绑定在一起（并且可选是否具有全部属性列表）。As a result, if one trusts the CA (and knows its public key), it can trust that the specific actor is bound to the public key included in the certificate, and owns the included attributes, by validating the CA’s signature on the actor’s certificate._
就是相当于我的MSP认可了这个CA发的证书，那么我就信任所有获得这个CA证书的人

证书可以广泛传播，因为它们既不包括参与者也不包括 CA 的私钥。因此，它们可以用作信任的锚，用于验证来自不同参与者的消息。

**CA也有自己的证书**

**一个或多个 CA 从数字角度定义了组织的成员** 。

### 根CA，中间CA和信任链

CA 有两种形式：**根 CA**和**中间 CA** 。
这些中间 CA 具有由根 CA 或其他中间 CA 颁发的证书，允许为链中的任何 CA 颁发的任何证书建立“信任链”。

![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202405191812312.png)

_只要每个中间 CA 的证书的颁发 CA 是根 CA 本身或具有对根 CA 的信任链，这些中间CA和RCA就组成了一条信任链。_

因为 CA 非常重要，Fabric 提供了一个内置的 CA 组件，允许在你的区块链网络中创建 CA。此组件称为 **Fabric CA** ，是一个私有根 CA 提供者，能够管理具有 X.509 证书形式的 Fabric 参与者的数字身份。

## 证书撤销列表

证书撤销列表（Certificate Revocation List，CRL）很容易理解，它是 CA 知道由于某些原因而被撤销的证书的引用列表。如果你回想商店场景，CRL 就像被盗信用卡列表一样。

![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202405191823620.png)

现在你已经了解了 PKI 如何通过信任链提供可验证的身份，下一步是了解如何使用这些身份来代表区块链网络的可信成员。这就是 MSP 发挥作用的地方——**它确定了区块链网络特定组织的成员**。

# 成员服务提供者（MSP）
因为Fabric是一个认证性的网络, 所以区块链参与者需要一种向网络中的其他参与者证实自己身份的机制从而在网络中进行交易。

证书机构通过生成可以用来证实身份的由公钥和私钥形成的键值对来发放认证信息。因为一个私钥永远不会被公开，所以引入了一种可以证实身份的机制即MSP。
MSP是一个可让身份被信任和被网络中其他参与者公认的，而不需要暴露成员的私钥的机制。

![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202405192255293.png)
_Identities类似于你的信用卡，用来证明你可以支付。MSP类似于被商店接受的信用卡清单。_

MSPs are used to define the organizations that are trusted by the network members. MSPs are also the mechanism that provide members with a set of roles and permissions within the network.

Finally, consider if you want to join an _existing_ network, you need a way to turn your identity into something that is recognized by the network. The MSP is the mechanism that enables you to participate on a permissioned blockchain network.

成员在Fabric上交易需要：
- 1. Have an identity issued by a CA that is trusted by an organization. The organization MSP determines which CAs are trusted by the organization.
- 2. Check that the organization MSP is added to the channel. This means that the organization is recognized and approved by the network members.
- 3. Ensure the MSP is included in the [policy](https://hyperledger-fabric.readthedocs.io/en/release-2.5/policies/policies.html) definitions on the network.

## MSP是什么

Membership Service Provider does not actually provide anything. Rather, the implementation of the MSP requirement is a set of folders that are added to the configuration of the network and is used to define an organization both inwardly (organizations decide who its admins are) and outwardly (by allowing other organizations to validate that entities have the authority to do what they are attempting to do).
MSP只是一个定义net配置的文件夹集合，用于配置一个organization的内部和外部。

MSP决定哪个root CA和 中间CA去定义信任域的成员，或通过列出成员的ID，或授权CAs

- But the power of an MSP goes beyond simply listing who is a network participant or member of a channel
 - It is the MSP that turns an identity into a **role** by identifying specific privileges an actor has on a node or channel.
 - In addition, an MSP can allow for the identification of a list of identities that have been revoked

## MSP域

MSP 出现在两个位置：

- 在参与者节点本地（**本地 MSP**）
- 在通道配置中（**通道 MSP**）

本地MSPs和通道MSPs之间的关键区别不在于它们如何运作——它们都将身份转化为角色——而是它们的**范围**。
### 本地MSP
**本地MSP是为客户端和节点(peer节点和排序节点)定义的**。
本地MSPs定义节点的权限(例如，谁是可以操作节点的peer节点管理员)。客户端(以上银行场景的账户持有人)的本地MSP,允许用户作为一个通道成员或作为一个特定角色的所有者如组织管理者，在其交易(如链码交易)进行身份验证从而进入系统,例如,进行配置交易。

**每个节点都必须定义一个本地MSP**，因为它定义了在该级别上谁拥有管理权或参与权。请注意，一个组织可以拥有一个或多个节点。MSP定义了组织管理员。组织、组织的管理员、节点的管理员以及节点本身都应该具有相同的信任根。

排序节点的本地MSP也在节点的文件系统上定义，并且只应用于该节点。与peer节点一样，排序节点也由单个组织拥有，因此有一个MSP来列出它信任的参与者角色或节点。

### Channel MSPs

 **channel MSPs define administrative and participatory rights at the channel level**.

 This means that if an organization wishes to join the channel, `an MSP` incorporating the chain of trust for the organization’s members `would need to be included in the channel configuration.`
> 想要加入一个organization， 必须要加入MSP，这个MSP可以让organization’s member加入信任链。

 Whereas local MSPs are represented as a folder structure on the file system, channel MSPs are described in a channel configuration.
![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202405202255769.png)

**Channel MSPs identify who has authorities at a channel level**.The channel MSP defines the _relationship_ between the identities of channel members (which themselves are MSPs) and the enforcement of channel level policies.Channel MSPs contain the MSPs of the organizations of the channel members.
> channel MSPs确定谁有通道级别的权限， 定义了通道中不同身份的关系和通道的策略。 Channel MSPs 包含了通道成员的组织的MSP

**Every organization participating in a channel must have an MSP defined for it**.  The MSP defines which members are empowered to act on behalf of the organization.  This includes configuration of the MSP itself as well as approving administrative tasks that `the organization has role, such as adding new members to a channel.` 
> channel中的每一个参与者都有一个MSP。这个MSP定义了organization中的几个代表者，他们可以配置MSP，也可以批准授权的任务。

**The channel MSP includes the MSPs of all the organizations on a channel.** peer organization和 orderer organization都有。

**Local MSPs are only defined on the file system of the node or user** to which they apply.
**a channel MSP is also instantiated on the file system of every node in the channel and kept synchronized via consensus**. 但是Channel MSP逻辑上是在Channel级的

![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202405202318223.png)
_The MSPs related to ORG1, the local MSP of the node and the global MSP that formally represents ORG1 on the channel, have been created by RCA1, the CA for ORG1. The peer organization, ORG2, also has a local MSP for its peer and another global MSP that represents ORG2 on the channel._

## What role does an organization play in an MSP

An **organization** is a logical managed group of members.
 What’s most important about organizations (or **orgs**) is that they manage their members under a single MSP. The MSP allows an identity to be linked to an organization.

 Organization `ORG1` would likely have an MSP called something like `ORG1-MSP`。 In some cases an organization may require multiple membership groups，it makes sense to have multiple MSPs and name them accordingly, e.g., `ORG2-MSP-NATIONAL` and `ORG2-MSP-GOVERNMENT`.

### Organizational Units （OUs） and MSPs

An organization can also be divided into multiple **organizational units**, each of which has a certain set of responsibilities, also referred to as `affiliations`。For example, the `ORG1` organization might have both `ORG1.MANUFACTURING` and `ORG1.DISTRIBUTION` OUs to reflect these separate lines of business. When a CA issues X.509 certificates, the `OU` field in the certificate specifies the line of business to which the identity belongs.   A benefit of using OUs like this is that these values can then be used in policy definitions in order to restrict access or in smart contracts for attribute-based access control.
> 我认为，OU就是将一个组织分成多个组织，但是不是通过多个MSP划分，而是通过certificate中的OU字段来划分，这个OU字段可以在 policy和smart contracts中使用。

If OUs are not used, all of the identities that are part of an MSP — as identified by the Root CA and Intermediate CA folders — will be considered members of the organization.
> 如果OU字段没有启用，那么所有的OU的member都被认为是organization的member，等于说权限放宽了


### Node OU Roles and MSPs

Additionally, there is a special kind of OU, sometimes referred to as a `Node OU`, that can be used to confer a role onto an identity. These Node OU roles are defined in the `$FABRIC_CFG_PATH/msp/config.yaml` file and contain a list of organizational units whose members are considered to be part of the organization represented by this MSP.   This is particularly useful when you want to restrict the members of an organization to the ones holding an identity (signed by one of MSP designated CAs) with a specific Node OU role in it. For example, with node OU’s you can implement a more granular endorsement policy that requires Org1 peers to endorse a transaction, rather than any member of Org1.
> Node OU是一种特殊的OU，通过指定一个identity的role来划分不同的权限， 而OU是通过指定affiliation来实现权限划分， 类似但不太一样。其实可以将Node OU理解程更加细颗粒度的OU，通过指定Node的role，就可以在OU下面将权限划分的更加明白一点。
>  Node OU roles定义在config.yaml中，包含了Node OU的列表 ？？？
>  当你想指定一个organization中的部分成员的时候，通过identity对应的role，我们就可以指定仅identity属于role的时候，才能被选中。

```js
NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/ca.sampleorg-cert.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/ca.sampleorg-cert.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/ca.sampleorg-cert.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/ca.sampleorg-cert.pem
    OrganizationalUnitIdentifier: orderer
```

This convention allows you to distinguish MSP roles by the OU present in the CommonName attribute of the X509 certificate. The example above says that any certificate issued by cacerts/ca.sampleorg-cert.pem in which OU=client will identified as a client, OU=peer as a peer, etc。
>  如果说原来将一个peer变成一个admin在MSP文件夹下面将peer证书移动到admin所在的文件夹下面，那么现在你只需要把peer的node ou字段改成admin就可以了。

These Role and OU attributes are assigned to an identity when the Fabric CA or SDK is used to `register` a user with the CA. It is the subsequent `enroll` user command that generates the certificates in the users’ `/msp` folder.

![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202405210809727.png)


The resulting ROLE and OU attributes are visible inside the X.509 signing certificate located in the `/signcerts` folder. The `ROLE` attribute is identified as `hf.Type` and refers to an actor’s role within its organization,

![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202405210816130.png)


Finally, OUs could be used by different organizations to distinguish each other. But in such cases, the different organizations have to use the same Root CAs and Intermediate CAs for their chain of trust, and assign the OU field to identify members of each organization.
> 当然了，OU不但可以划分organization的子集，也可以用来划分不同的组织。


## MSP结构

local MSP
![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202405210820311.png)


Channel MSP
**Revoked Certificates:** If the identity of an actor has been revoked, identifying information about the identity — not the identity itself — is held in this folder.


# Policies

## What is a policy
At its most basic level, a policy is a set of rules that define the structure for how decisions are made and specific outcomes are reached. To that end, policies typically describe a **who** and a **what**, such as the access or rights that an individual has over an **asset**.
> policy就是一组规则，指定如何做决定和如何得到特定的结果。 policy通常描述`谁，什么`，表示谁对特定的资源有访问权限。

in Hyperledger Fabric, policies are the mechanism for infrastructure management. Fabric policies represent how members come to agreement on accepting or rejecting changes to the network, a channel, or a smart contract. Simply put, everything you want to do on a Fabric network is controlled by a policy.
> Policy 是用来管理网络结构的机制。Policy 表示了网络中的各个成员如何对发生的各种事情达成统一。所以，在网络中想做什么都需要被policy来控制。

## Why are policies needed

在Ethereum和Bitcoin这样的网络中，policy的更改只能通过改代码，过于死板。但是在Fabric 中，policy可以由user来决定，无论是在网络启动前还是运行时。

Policies allow members to decide which organizations can access or update a Fabric network, and provide the mechanism to enforce those decisions. Policies contain the lists of organizations that have access to a given resource, such as a user or system chaincode. They also specify how many organizations need to agree on a proposal to update a resource, such as a channel or smart contracts.
> Policy让member来决定哪个组织可以修改network，同时也提供了执行决定的机制。


## How are policies implemented

Policies are defined within the relevant administrative domain of a particular action defined by the policy. For example, the policy for adding a peer organization to a channel is defined within the administrative domain of the peer organizations. (known as the `Application` group). Similarly, adding ordering nodes in the consenter set of the channel is controlled by a policy inside the `Orderer` group. Actions that cross both the peer and orderer organizational domains are contained in the `Channel` group.
> 简而言之，Policy在哪里发挥作用，就在哪里被定义。











# Peers

## peer and Application

![d63fdbcc19245b7522efa8d911f9f61.jpg](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202405282311136.jpg)



# ledger

Ledger ：
- world State
- Blockchain


## Blocks

Block Header:
- Block number
- Current Block Hash
- Previous Block Header Hash
![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202405282314087.png)

Block Data:
- transactions

Block MetaData(不参与hash计算)
- certificate
- signature
- valid/invalid indicator
- a hash of 这个block以及他之前的所有状态更新

## Transaction


![image.png](https://mypictures-1308119878.cos.ap-shanghai.myqcloud.com/Obsidian_notebook/202405282322951.png)


