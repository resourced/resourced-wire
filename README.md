[![GoDoc](https://godoc.org/github.com/resourced/resourced-wire?status.svg)](http://godoc.org/github.com/resourced/resourced-wire)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](LICENSE.md)

This library provides parsing logic on ResourceD TCP wire protocol.

Example:

```
type:base64|created:unix-timestamp|content:base64=
type:plain|created:unix-timestamp|content:plaintext
type:json|created:unix-timestamp|content:{"foo": "bar"}
topic:topic-name|type:json|created:unix-timestamp|content:{"foo": "bar"}  # inside master's MessageBus.
```

In one TCP connection, you can send multiple payloads by separating them using `\n`.

## Where does it get used?

* **Agent:** User can send a log line to the agent using this protocol.

* **Master:** Within Master daemon, it is used for passing data through the message bus.
