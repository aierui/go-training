# Timeout

如何实现超时控制，主要有三种:

- context.WithTimeout/context.WithDeadline + time.After
- context.WithTimeout/context.WithDeadline + time.NewTimer
- channel + time.After/time.NewTimer