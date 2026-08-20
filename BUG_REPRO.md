# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

人民路站点的灭火器整改派发偶尔报 500，可刷新详情后任务已经处于“已派发”，审计里却找不到对应记录，运营企业和监管端因此看到两套状态。请修复派发过程，让它要么整体成功、要么不留下任何状态变化。用于复现这次半成功的测试用例不得修改，也不能靠跳过检查或放宽断言通过。

## 含 Bug 版本

- 仓库：11DingKing/chargeguard-bug-01
- 仓库地址：https://github.com/11DingKing/chargeguard-bug-01.git
- parent SHA：ebce58e8169e8bef9a64d1083f159c5ae09c4616

## 复现步骤

```bash
git clone -- https://github.com/11DingKing/chargeguard-bug-01.git bug-repro
cd bug-repro
git checkout --detach ebce58e8169e8bef9a64d1083f159c5ae09c4616
go test ./internal/httpapi -run TestTaskBehavior -count=1
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/httpapi -run TestTaskBehavior -count=1
--- FAIL: TestTaskBehavior (0.00s)
    task_behavior_test.go:18: snapshot={Status:assigned Owner:operator-a Audits:0}
FAIL
FAIL	chargeguard/internal/httpapi	0.063s
FAIL

```

stderr：

```text
(empty)
```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/httpapi -run TestTaskBehavior -count=1
--- FAIL: TestTaskBehavior (0.00s)
    task_behavior_test.go:18: snapshot={Status:assigned Owner:operator-a Audits:0}
FAIL
FAIL	chargeguard/internal/httpapi	0.003s
FAIL

```

stderr：

```text
(empty)
```

## 通过条件

修复后，在题面描述的触发条件下应得到预期业务结果且不再出现原始症状；定向验证命令修复前必须失败、应用修复后必须通过，相关回归和仓库全量测试必须通过；不得新增、删除或修改测试文件，不得跳过测试、降低断言或绕过目标逻辑。
