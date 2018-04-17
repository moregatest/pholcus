延伸自
https://github.com/henrylee2cn/pholcus

變更紀錄
=======================================
## phantom的改善
官方版本的phantom模式js是固定無法隨抓取規則變更 我的版本可以透過SetTemp添加客製js
```
					req := &request.Request{
						Url:          initUrl,
						Rule:         "keywds",
						ConnTimeout:  -1,
						Reloadable:   true,
						DownloaderID: 1,
					}
					req.Prepare()
					req.SetTemp("js", cjs)
					ctx.AddQueue(req)
 ```
