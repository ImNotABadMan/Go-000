学习笔记

1. 使用errgroup控制全部goroutine的正常退出，errgroup底层是sync.WaitGroup   
2. 使用超时和waitgroup，使得goroutine得以控制，不会造成野生goroutine
